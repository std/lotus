package Posting

import "C"
import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm/ledger"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	miner_ledger "github.com/filecoin-project/lotus/chain/vm/ledger/miner-ledger"
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	logging "github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"time"
)




type LedgerPosting struct {
	//vm           *vm.VM
	CurrentTxId int
	OriginMsg  *types.Message
//	CurrentMsg *types.Message
	st         *state.StateTree
	cst        *cbor.BasicIpldStore

	Epoch   abi.ChainEpoch
	//gasOutputs miner_ledger.GasOutputs

	*miner_ledger.MinerLedger
	MarketLedger
	//*PowerLedger
	//RewardLedger

	messages []*models.TxMessage
	minerEntries map[int32]*models.LedgerEntry
	powerEntries map[int]*models.PowerEntry
	sectors  map[int]*models.Sector
	rewardEntries map[int]*models.RewardEntry


	SectorId int32
	PostingStack ledg_util.Stack
}
func (l *LedgerPosting) db(	) *gorm.DB{
	return ledg_util.GetPgDatabase()
}

func (l *LedgerPosting) dbInsert(v interface{},checkExistance bool) {
	if checkExistance {
		l.db().FirstOrCreate(v)
	}else {
		l.db().Create(v)	}
}

func (l *LedgerPosting) insert(v interface{},checkExistance bool) {
	//if ledg_util.Exists(v,int(id)) {return }

	switch t := v.(type) {

	case *models.Tipset:
		l.dbInsert(t, checkExistance)
	case *models.Block:
		l.dbInsert(t, checkExistance)
	case *models.Account:
		l.dbInsert(t, checkExistance)
	case *models.Sector:
		l.dbInsert(t, checkExistance)


	case *models.TxMessage:
		l.dbInsert(t, checkExistance)
		l.messages=append(l.messages,t)

	case *models.PowerEntry:
		l.dbInsert(t,checkExistance)
	case *models.LedgerEntry:
		l.minerEntries[t.Id] = t
		l.dbInsert(t,checkExistance)


		//l.minerEntries[e.Id]=e
	case *models.StorageDealEntry:
		l.dbInsert(t,checkExistance)
	case *models.RewardEntry:
		l.dbInsert(t,checkExistance)
	default:
		llog.Info("unexpected type %T", t, " of ", v)

	}
}


var llog			= logging.Logger("gledger")



//func (l *LedgerPosting) OnStartMessage(ctx context.Context,Msg *types.Message){
//
//
//
//	addr:=Msg.To
//	switch ledg_util.GetActorType(addr) {
//
//	case "account":
//
//	case"init"://f01
//
//	case "reward"://f02
//
//	case "cron"://f03
//
//	case "power"://f04
//
//	case "market"://f05
//
//	case "registry"://f06
//
//	case "burnt":
//	case "paych":
//
//	case "miner":
//		l.MinerLedger=&miner_ledger.MinerLedger{
//			MinerAddress: address.Address{},
//			//Entries:      nil,
//			Opening:      nil,
//			Closing:      nil,
//			BalancesDiff: nil,
//		}
//
//	}
//}

func (l *LedgerPosting) insertMinerEntrySend_del(p ledg_util.ActorMethodParams){
	//e:=l.minerEntrySendTemplate(p,0)
	e:=l.minerEntryTemplate(p,0,false)
	e.MethodName=e.MethodName+"#Send"
	amount:=ledg.FilAmount(big.NewFromGo(p.Msg.Value.Int).Neg())

	addressId,_:=ledg_util.GetOrCreateAccountFromAddress(p.Msg.From,"",l.Epoch)
	offsetId,_:=ledg_util.GetOrCreateAccountFromAddress(p.Msg.To,"",l.Epoch)
	e.AddressId=int32(addressId)
	e.OffsetId=int32(offsetId)

	e.Amount =amount
	l.insert(&e,true)
}

func (l *LedgerPosting) insertGasEntry(msg *types.Message, outp *ledg.GasOutputs) {

	e:=l.minerEntryTemplate(ledg_util.ActorMethodParams{
		Msg:   msg,
		Ret:   nil,
		Depth: 0,
	},0	,false)

	e.Amount=ledg.FilAmount{big.NewInt(0).SetInt64(outp.GasBurned)}
	//addressId,_:=address.IDFromAddress(msg.From)
	offsetId,_:=address.IDFromAddress(msg.From)
	e.AddressId=int32(99)
	e.OffsetId=int32(offsetId)
	e.Method=99
	e.MethodName="GasBurnt"
	e.EntryType="GasBurnt"
	l.insert(&e,true)

	}
func (l *LedgerPosting) StartMessage(ctx context.Context,msg *types.Message,implicit bool) {
	if ctx.Value("replay") == nil {return}

	if  ctx.Value("epoch")==nil {		panic("GL.StartMessage: epoch must be not nil")	}

	if  l==nil  { panic ("GL.StartMessage: GL must be not nil") }
	if msg==nil {panic("GL.StartMessage: param Msg is nil")}


	epoch:=ctx.Value("epoch").(abi.ChainEpoch)


	l.Epoch=epoch
	l.OriginMsg=msg

	l.InsertTxMessage(l.Epoch, len(l.messages),msg,implicit)



	llog.Info("Start Explicit Message ",l.CurrentTxId)


	//l.MinerLedger.OriginMsg=Msg
	//l.MinerLedger.IsImplicit=implicit

	//l.Opening=l.getBalances()
}

func (l *LedgerPosting) FinalizeMessage(ctx context.Context, msg *types.Message,outp *ledg.GasOutputs){


	if ctx.Value("replay") == nil {return}
	if  l==nil  { panic ("GL.StartMessage: GL must be not nil") }


	//var p *ledg_util.ActorMethodParams

	for p:=l.PostingStack.Pop(); p!=nil; p = l.PostingStack.Pop() {
		//llog.Info("Posting: "+p.Msg.Cid().String()+" "+strconv.Itoa(l.PostingStack.GetCount()))
		l.Posting(ctx,p)
	}

	if outp!=nil {	l.insertGasEntry(msg,outp) }



	l.OriginMsg=nil
	return

	//l.FinalizeMsg(ctx, outp,l.Epoch)
	//l.fillClosingBalance()
	//addr,_:=address.NewFromString("t01098")
	//l.InsertActorHead("Finalize",nil,addr)

	l.AppendRootEntry(ctx,l.OriginMsg,outp,0)

	//e:=l.Entries[Msg.Cid().String()]


	//if e!=nil {
	//	miner:=l.OriginMsg.To

		//e.MethodName=l.originMsg.Cid().String()//z_del
		//e.Opening=l.Opening[miner]
		//e.Balance =l.Closing[miner]
		//e.Amount=l.BalancesDiff[miner]
		//l.Entries[l.currentMsg.Cid().String()]=e
	//}
	//l.Flush(ctx)


}

func (l *LedgerPosting) StartImplicitMessage(ctx context.Context,msg *types.Message) {
	if ctx.Value("replay") == nil {return}
	if  l==nil  { panic ("GL.StartImplicitMessage: GL must be not nil") }
	llog.Info("StartImplicitMessage\n",msg)
	l.StartMessage(ctx,msg,true)
}

func (l *LedgerPosting) FinalizeImplicitMessage(ctx context.Context,msg *types.Message){

	if ctx.Value("replay") == nil {return}
	if  l==nil  { panic ("GL.FinalizeImplicitMessage: GL must be not nil") }

	l.FinalizeMessage(ctx,msg,nil)

	llog.Info("FinalizeImplicitMessage\n",msg)

	l.MinerLedger.Flush(ctx)

	l.OriginMsg=nil
}

func (l*LedgerPosting) insertBlockHeader(header *types.BlockHeader, epoch abi.ChainEpoch,blockIdx int) {
	blockId:=int(epoch)*10+blockIdx


	minerId,_:=address.IDFromAddress(header.Miner)
	ledg_util.GetOrCreateAccountFromId(int32(minerId),"",epoch)
	b:=models.Block{
		Id:       int32(blockId),
		Time:     time.Time{},
		TipsetId: int32(epoch),
		//Tipset:   models.Tipset{},
		MinerId:  int32(minerId),
		//Miner:    models.Account{},
		Cid:      ledg.Cid{header.Cid()},
	}
	l.insert(&b,true)
	//db.FirstOrCreate(&b)
}

func (l*LedgerPosting)startTipset(ts *types.TipSet,epoch abi.ChainEpoch) {
	l.insertTipset(ts,epoch)
	blks := ts.Blocks()
	for i := 0; i < len(blks); i++ {
		l.insertBlockHeader(ts.Blocks()[i], epoch, i)
	}
}
func (l*LedgerPosting)insertTipset(ts *types.TipSet,epoch abi.ChainEpoch) {

	ts_id:=int32(epoch)


	//var found models.Tipset
	//result:=db.First(&found,ts_id)
	//if result.RowsAffected>0{ return }

	var arr [8]cid.Cid
	copy(arr[:],ts.Cids())

	tipset:=models.Tipset{
		Id:     ts_id,
		Time:   time.Time{},
		Block0: ledg.Cid{arr[0]},
		Block1: ledg.Cid{arr[1]},
		Block2: ledg.Cid{arr[2]},
		Block3: ledg.Cid{arr[3]},
		Block4: ledg.Cid{arr[4]},
		Block5: ledg.Cid{arr[5]},
		Block6: ledg.Cid{arr[6]},
		Block7: ledg.Cid{arr[7]},

	}
	l.insert(&tipset,true)

}

//func (l * LedgerPosting) Log(m bson.M){
//	mgo:=ledg_util.GetOrCreateMongoConnection()
//	mgo.InsertLogEntry(m)
//}

func CreateGL(ctx context.Context, st *state.StateTree,  ts *types.TipSet,epoch abi.ChainEpoch ) *LedgerPosting {



	l:= &LedgerPosting{

		st:           st,
		//cst:          cst,
		Epoch: epoch,
		minerEntries: make(map[int32]*models.LedgerEntry),

		//MarketLedger: market_ledger.MarketLedger{},
		//MinerLedger:  miner_ledger.CreateGL(ctx,st,cst,dump,epoch),
		//PowerLedger:  CreatePowerLedger(ctx,epoch),
	}

	llog.Info("Ledger.CreateGL")

	l.startTipset(ts,epoch)



	return l
}



func (l *LedgerPosting) Send(ctx context.Context, msg *types.Message, methReturn []byte,gasUsed int64,epoch abi.ChainEpoch,callDepth uint64) error{
	return nil
}

func (l *LedgerPosting) InsertTxMessage(epoch abi.ChainEpoch,i int, msg *types.Message,implicit bool) int {

	txId :=int(epoch)*1000000+i

	//var found models.TxMessage
	l.CurrentTxId=txId
	//result:=db.First(&found, txId)
	//if result.RowsAffected>0{ return 0 }

	fromId,_:=address.IDFromAddress(msg.From)
	ToId,_:=address.IDFromAddress(msg.To)


	ledg_util.GetOrCreateAccountFromId(int32(fromId),"",epoch)
	ledg_util.GetOrCreateAccountFromId(int32(ToId),"",epoch)

	tx:=models.TxMessage{
		Id:         txId,
		FromId:     int32(fromId),
		ToId:       int32(ToId),
		From:       models.Account{},
		To:         models.Account{},
		Value:      ledg.FilAmount(msg.Value),
		GasFeeCap:  ledg.FilAmount(msg.GasFeeCap),
		GasPremium: ledg.FilAmount(msg.GasPremium),
		GasLimit:   msg.GasLimit,
		Method:     msg.Method,
		MethodName:  ledg_util.GetMethodName(msg),
		Params:     msg.Params,
		Implicit: implicit,
	}
	l.insert(&tx,true)

	//if !implicit {
	//	l.insertMinerEntrySend(ledg_util.ActorMethodParams{
	//			Msg:   msg,
	//			Ret:   nil,
	//			Depth: 0,
	//		})
		//e:=l.minerEntrySendTemplate(ledg_util.ActorMethodParams{
		//	Msg:   msg,
		//	Ret:   nil,
		//	Depth: 0,
		//
		//},0)
		//e.TxId=txId
		//l.insert(&e,true)

	//}

	return txId
}

func (l *LedgerPosting) EnqueuePosting(ctx context.Context, msg *types.Message, methReturn []byte,callDepth uint64) error{
	if ctx.Value("replay") == nil {return nil}
	if  l==nil  { panic ("GL.Posting: GL must be not nil") }

	if l.OriginMsg==nil {return nil}

	l.PostingStack.Push (&ledg_util.ActorMethodParams{
		Msg:        msg,
		Ret: methReturn,
		Depth:  callDepth,
	})
	return nil
}
func (l *LedgerPosting) calcBalance() ledg.FilAmount{

	bal:=ledg.FilAmountFromInt(0)
	for _,e:=range l.minerEntries {
		bal=ledg.FilAmount{Int:bal.Add(bal.Int,e.Amount.Int)}
	}

	return bal
}


func (l *LedgerPosting) Posting(ctx context.Context, p *ledg_util.ActorMethodParams) error{
	if ctx.Value("replay") == nil {return nil}
	if  l==nil  { panic ("GL.Posting: GL must be not nil") }

	if l.OriginMsg==nil {return nil}
	//l.SetCurrentMsg(curMsg)


	methName:=ledg_util.GetMethodName(p.Msg)
	params:= ledg_util.ActorMethodParams{
		Msg:   p.Msg,
		Ret:   p.Ret,
		Depth: p.Depth,
	}

	var minerMethods ledger.MinerMethods = l
	var powerMethods ledger.PowerMethods = l
	var marketMethods ledger.MarketMethods = l

	addr:= p.Msg.To

	if l.CurrentTxId ==0 { panic("GL.Posting pre=>CurrentTxId is 0")}
	switch ledg_util.GetActorType(addr) {

	case "account":
		switch methName {
			case "send":
			case "PubkeyAddress":
		}
	case"init"://f01
		switch methName {
			case "Constructor":
			case "Exec":
		}
	case "reward"://f02
		switch methName {
			case "Constructor":
			case "AwardBlockReward":
			case "ThisEpochReward":
			case "UpdateNetworkKPI":
		}
	case "cron"://f03
		switch methName {
			case "Constructor":
			case "EpochTick":
		}
	case "power"://f04
		switch methName {
			case "Constructor": powerMethods.PowerConstructor(params)
			case "CreateMiner": powerMethods.CreateMiner(params)
			case "UpdateClaimedPower":powerMethods.UpdateClaimedPower(params)
			case "EnrollCronEvent":powerMethods.EnrollCronEvent(params)
			case "OnEpochTickEnd":powerMethods.OnEpochTickEnd(params)
			case "UpdatePledgeTotal":powerMethods.UpdatePledgeTotal(params)
			case "Deprecated1":powerMethods.Deprecated1(params)
			case "SubmitPoRepForBulkVerify":powerMethods.SubmitPoRepForBulkVerify(params)
			case "CurrentTotalPower":powerMethods.CurrentTotalPower(params)
		}
	case "market"://f05
		switch methName {
			case "Constructor":
			case "AddBalance":
			case "WithdrawBalance":
			case "PublishStorageDeals":marketMethods.PublishStorageDeals(params) //Ret="PublishStorageDeals"
			case "VerifyDealsForActivation":
			case "ActivateDeals":
			case "OnMinerSectorsTerminate":
			case "ComputeDataCommitment":
			case "CronTick":
		}
	case "registry"://f06
		switch methName {
			case "Constructor":
			case "AddVerifier":
			case "RemoveVerifier":
			case "AddVerifiedClient":
			case "UseBytes":
			case "RestoreBytes":
		}
	case "burnt":
	case "paych":
		switch methName {
			case "Constructor":
			case "EpochTick":
		}
	case "miner":
		switch methName {
			case "Constructor":minerMethods.MinerActorConstructor(params)

			case "ControlAddresses":minerMethods.ControlAddresses(params)
			case "ChangeWorkerAddress":minerMethods.ChangeWorkerAddress(params)
			case "ChangePeerID":minerMethods.ChangePeerID(params)
			case "SubmitWindowedPoSt":  minerMethods.SubmitWindowedPoSt(params)

			case "PreCommitSector":		minerMethods.PreCommitSector(params)
			case "ProveCommitSector":	minerMethods.ProveCommitSector(params)

			case "ExtendSectorExpiration":minerMethods.ExtendSectorExpiration(params)

			case "TerminateSectors":minerMethods.TerminateSectors(params)
			case "DeclareFaults":minerMethods.DeclareFaults(params)
			case "DeclareFaultsRecovered":minerMethods.DeclareFaultsRecovered(params)
			case "OnDeferredCronEvent":minerMethods.OnDeferredCronEvent(params)
			case "CheckSectorProven":
			case "ApplyRewards":
			case "ReportConsensusFault":
			case "WithdrawBalance":
			case "ConfirmSectorProofsValid":
			case "ChangeMultiaddrs":
			case "CompactPartitions":
			case "CompactSectorNumbers":
			case "ConfirmUpdateWorkerKey":
			case "RepayDebt":
			case "ChangeOwnerAddress":
			case "DisputeWindowedPoSt":
		}
	}

	//if l.CurrentTxId!=130069000001 {return nil}
	implemented := map[string]bool{
		"PreCommitSector": true,
		"CreateMiner":   true,
		"Constructor":   true,
	}


	if implemented[methName] {
		//llog.Info(methName," # ",l.CurrentTxId,", Depth: ",params.Depth," Implemented")
		return 	nil
	}else {
		//llog.Info(methName," # ",l.CurrentTxId,", Depth: ",params.Depth," Not Implemented")
	}


	idx:=int32(len(l.minerEntries))

	e:=l.minerEntryTemplate(params,idx,false)
	ledg_util.GetOrCreateAccountFromId(e.AddressId,"",l.Epoch)
	ledg_util.GetOrCreateAccountFromId(e.OffsetId,"",l.Epoch)
	e.CallDepth=params.Depth
	e.TxId=l.CurrentTxId
	l.insert(&e,false)

	//l.minerEntryTemplate(params,int(l.Epoch*1000000)+idx)

	return nil
}

func (l *LedgerPosting) FinalizeGL() {
	bal:=l.calcBalance()
	llog.Infof("tipset finalize: tx count %d, miner entries count %d, balance %s",len(l.messages), len(l.minerEntries),bal.String())
	for _,e:=range l.minerEntries{
		llog.Infof("Entry# %d method %s",e.Id,e.MethodName)
	}
}



//func (l *LedgerPosting) Flush1(ctx context.Context){
//	if ctx.Amount("export")==nil {return}
//	if (l==nil) {return}
//
//	//for i:=range l.Entries{
//	//	InsertEntry(l.Entries[i])
//	//}
//}


