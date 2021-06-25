package Posting

import "C"
import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm/ledger"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
	"github.com/filecoin-project/lotus/chain/vm/ledger/stateTransition"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	logging "github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"time"
)


type LedgerPosting struct {

	Stinfo      *stateTransition.StateTrans
	CurrentTxId int
	SectorId     *int32
	MinerId 	 int32

	ts          *types.TipSet
	OriginMsg   *types.Message
	StateTree *state.StateTree
	cst       *cbor.BasicIpldStore

	Epoch abi.ChainEpoch

	//*miner_ledger.MinerLedger
	MarketLedger

	messages      []*models.TxMessage
	minerEntries  map[int]models.LedgerEntry
	powerEntries  map[int]*models.PowerEntry
	sectors       map[int]*models.Sector
	rewardEntries map[int]*models.RewardEntry


	PostingStack ledg_util.Stack
	adtStore     adt.Store
	//nextts       *types.TipSet
}
func (l *LedgerPosting) db(	) *gorm.DB{
	return ledg_util.GetPgDatabase()
}

func (l *LedgerPosting) dbInsert(v interface{},checkExistance bool) {
	if checkExistance {		l.db().FirstOrCreate(v)
	}else {
		_=l.db().Create(v)

		//val,ok:=v.(*models.LedgerEntry)
		//
		//if ok && val.TxId==130069000000{
		//	ledg_util.Llogf("insert entry %s",val.String())
		//}
	}
}
func (l *LedgerPosting) flush() {

	balance:=big.NewInt(0)

	for _,e:=range l.minerEntries {
		if l.CurrentTxId!=e.TxId {continue}
		balance=big.Add(abi.TokenAmount(e.Amount),balance)
		//ledg_util.Llogf("e.txId %d, id %d, amount %s, balance %s",e.TxId,e.Id,e.Amount, balance.String())
		l.dbInsert(&e,false)
	}

	ledg_util.Llogf("%d : balance inv %s",l.CurrentTxId,balance.String())
}

func (l *LedgerPosting) insert(v interface{},checkExistance bool) {

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

	//case *models.LedgerEntry:
	//	l.minerEntries[t.Id] = t
	//	l.dbInsert(t,checkExistance)
	case models.VestingEntry:
	case *models.VestingEntry:
		break
		//l.dbInsert(t,checkExistance)

	case models.LedgerEntry:
		l.minerEntries[int(t.Id)] = t
		//l.dbInsert(&t,checkExistance)
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



func (l *LedgerPosting) insertGasEntry(msg *types.Message, outp *ledg.GasOutputs) {

	callerAddressId,_:=ledg_util.GetOrCreateAccountFromAddress(msg.From,"",l.Epoch)

	//ledg_util.Llogf("AddressMongo from %s , %d",msg.From.String(),callerAddressId)
	gasHolderId:=int32(999)
	BurntId:=int32(99)
	RewardId:=int32(2)
	e:=l.minerEntryTemplate(ledg_util.ActorMethodParams{
		Msg:   msg,
		Ret:   nil,
		Depth: 0,
	},0	,false)
	e.SectorId=l.SectorId
	e.MinerId=l.MinerId

	gasSent:=ledg.FilAmount(big.Sum(outp.BaseFeeBurn,outp.OverEstimationBurn))

	e.Amount=gasSent.Neg()

	e.AddressId=callerAddressId
	e.OffsetId= gasHolderId
	e.MethodId=99
	e.MethodName="gasSent"
	e.EntryType="GasSent"


	l.insert(e,false)




	e.Id++
	e.Amount=ledg.FilAmount(outp.BaseFeeBurn)

	e.AddressId=BurntId
	e.OffsetId=gasHolderId
	e.MethodId=999
	e.MethodName="inferred"
	e.EntryType="BaseFeeBurn"
	l.insert(e,false)
	//

	//////////////////////////
	e.Id++
	e.Amount=ledg.FilAmount(outp.OverEstimationBurn)

	e.AddressId=BurntId
	e.OffsetId=gasHolderId
	e.MethodId=999
	e.MethodName="inferred"
	e.EntryType="OverEstimationBurn"
	l.insert(e,false)

/////////////////////////// Miner Tip

	e.Id++
	e.Amount=ledg.FilAmount(outp.MinerTip)

	e.AddressId=RewardId
	e.OffsetId=gasHolderId
	e.MethodId=999
	e.MethodName="inferred"
	e.EntryType="MinerTip"
	l.insert(e,false)

	e.Id++
	e.Amount=ledg.FilAmount(outp.MinerTip).Neg()

	e.AddressId=callerAddressId
	e.OffsetId=gasHolderId
	e.MethodId=999
	e.MethodName="inferred"
	e.EntryType="MinerTip"
	l.insert(e,false)
}


func (l *LedgerPosting) SetStateTree(ctx context.Context,st *state.StateTree) error{
	//if st==nil {
	//	return  xerrors.Errorf("Error: StateTree is nil for message from %s, to %s, meth %S", msg.From,msg.To,ledg_util.GetMethodName(msg))
	//}
	l.StateTree=st
	l.Stinfo.Stree=st

	blocks:=l.ts.Blocks()
	for _,block:=range blocks {  l.Stinfo.LoadActorState(ctx,block.Miner, ledg.Opening)	}
	return nil
}

func (l *LedgerPosting) StartMessage(ctx context.Context, st *state.StateTree, msg *types.Message, implicit bool) {
	if ctx.Value("replay") == nil {return}

	if  ctx.Value("epoch")==nil {		panic("GL.StartMessage: epoch must be not nil")	}

	if  l==nil  { panic ("GL.StartMessage: GL must be not nil") }
	if msg==nil {panic("GL.StartMessage: param Msg is nil")}

	//if err:=l.SetStateTree(st,msg);err!=nil {
	//	ledg_util.Llogf(err.Error())
	//}


	epoch:=ctx.Value("epoch").(abi.ChainEpoch)

	l.Epoch=epoch
	l.OriginMsg=msg

	l.InsertTxMessage(l.Epoch, len(l.messages),msg,implicit)

	ledg_util.Llogf("Start Explicit Message %d",l.CurrentTxId)

	//addr1,_:=address.NewIDAddress(1002)
	//minerAddr,_:=address.NewIDAddress(8557)
	//actor1002_1,_:=l.StateTree.GetActor(addr1)
	l.Stinfo.LoadActorState(ctx,msg.To, ledg.Opening)
	//l.Stinfo.LoadActorState(ctx,minerAddr,Opening)

}

func (l *LedgerPosting) FinalizeMessage(ctx context.Context, msg *types.Message,outp *ledg.GasOutputs){


	if ctx.Value("replay") == nil {return}
	if  l==nil  { panic ("GL.StartMessage: GL must be not nil") }

	//addr1,_:=address.NewIDAddress(1002)
	//l.Stinfo.LoadActorState(ctx,addr1,Closing)
	//minerAddr,_:=address.NewIDAddress(8557)
	//l.Stinfo.LoadActorState(ctx,minerAddr,Closing)

	l.Stinfo.LoadActorState(ctx,msg.To,ledg.Closing)

	//for i,v:=range l.Stinfo.ActorStatesOnEnd {
	//	//ledg_util.Llogf("ActorStateList %d StateTrans %s", i.String(),v.String())
	//}

	for p:=l.PostingStack.Pop(); p!=nil; p = l.PostingStack.Pop() {
		l.Posting(ctx,p)
	}

	if outp!=nil {
		l.insertGasEntry(msg,outp)
	}


	l.flush()

	ledg_util.Llogf("FinalizeMessage %d",l.CurrentTxId)

	l.OriginMsg=nil
	l.SectorId=nil
	l.MinerId=0
}

func (l *LedgerPosting) StartImplicitMessage(ctx context.Context,msg *types.Message) {
	if ctx.Value("replay") == nil {return}
	if  l==nil  { panic ("GL.StartImplicitMessage: GL must be not nil") }
	//llog.Info("StartImplicitMessage\n",msg)
	l.StartMessage(ctx, nil, msg, true)
}

func (l *LedgerPosting) FinalizeImplicitMessage(ctx context.Context,msg *types.Message){

	if ctx.Value("replay") == nil {return}
	if  l==nil  { panic ("GL.FinalizeImplicitMessage: GL must be not nil") }

	l.FinalizeMessage(ctx,msg,nil)

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


func CreateGL(ctx context.Context, adtStore adt.Store,  ts *types.TipSet,epoch abi.ChainEpoch ) *LedgerPosting {

	l:= &LedgerPosting{

		ts:        ts,
		adtStore:  adtStore,
		Epoch: epoch,
		minerEntries: make(map[int]models.LedgerEntry),
		Stinfo: &stateTransition.StateTrans{
			Stree:              nil,
			Ts:                 ts,
			Store:              adtStore,
			ActorStates: make(map[abi.ActorID]stateTransition.ActorStateTrans),
		},
	}

	l.startTipset(ts,epoch)
	return l
}

func (l *LedgerPosting) Send(ctx context.Context, msg *types.Message, methReturn []byte,gasUsed int64,epoch abi.ChainEpoch,callDepth uint64) error{
	return nil
}

func (l *LedgerPosting) InsertTxMessage(epoch abi.ChainEpoch,i int, msg *types.Message,implicit bool) int {

	txId :=int(epoch)*1000000+i

	l.CurrentTxId=txId
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

	methName:=ledg_util.GetMethodName(p.Msg)
	params:= ledg_util.ActorMethodParams{
		Msg:   p.Msg,
		Ret:   p.Ret,
		Depth: p.Depth,
	}

	var minerMethods ledger.MinerMethods = l
	var powerMethods ledger.PowerMethods = l
	var marketMethods ledger.MarketMethods = l
	var rewardMethods ledger.RewardMethods = l

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
			case "Constructor":rewardMethods.RewardActorConstructor(params)
			case "AwardBlockReward":rewardMethods.AwardBlockReward(params)
			case "ThisEpochReward":rewardMethods.ThisEpochReward(params)
			case "UpdateNetworkKPI":rewardMethods.UpdateNetworkKPI(params)
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
			case "ApplyRewards":minerMethods.ApplyRewards(params)
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
	implemented := map[string]bool{

		//miner
		"PreCommitSector": true,
		"Constructor":   true,
		"ApplyRewards":   true,

		//reward
		"ThisEpochReward":true,
		"AwardBlockReward":true,
		//power
		"CurrentTotalPower":true,
		"CreateMiner":   true,
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
	e.CallDepth=int16(params.Depth)
	e.TxId=l.CurrentTxId
	e.MinerId=l.MinerId
	e.SectorId=l.SectorId
	l.insert(e,false)

	return nil
}

func (l *LedgerPosting) FinalizeGL(ctx context.Context) {

	//for _,block:=range l.ts.Blocks(){
	//	l.Stinfo.LoadActorState(ctx,block.Miner,Closing)
	//}
	bal:=l.calcBalance()
	llog.Infof("tipset finalize: tx count %d, miner entries count %d, balance %s",len(l.messages), len(l.minerEntries),bal.String())

}

