package miner_ledger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"go.mongodb.org/mongo-driver/bson"
)


func (l *MinerLedger) ActorBalance(addr address.Address) (types.BigInt, aerrors.ActorError) {

	act, err := l.st.GetActor(addr)
	if err != nil {
		return types.EmptyInt, aerrors.Absorb(err, 1, "failed to find actor")
	}

	return act.Balance, nil
}


func(l *MinerLedger) AddPrecommitDeposit(sector ledg.SectorID){

}

func(l *MinerLedger) UnlockPrecommitDeposit(sector ledg.SectorID){

}

func (l *MinerLedger) getBalances() ActorBalances {

	burntBalance,_	:=l.ActorBalance(builtin3.BurntFundsActorAddr)
	rewardBalance,_	:=l.ActorBalance(builtin3.RewardActorAddr)
	powerBalance,_	:=l.ActorBalance(builtin3.StoragePowerActorAddr)
	marketBalance,_	:=l.ActorBalance(builtin3.StorageMarketActorAddr)
	minerBal,_ 		:=l.GetMinerBalanceByAddressDim(l.MinerAddress)
	//minerBal:=ZeroDimBalance()
	//fromBal,_		:=l.ActorBalance(l.currentMsg.From)

	//if &l.currentMsg.To==nil {panic("l.currentMsg target is nil")}

	return ActorBalances{
		builtin3.BurntFundsActorAddr:    ledg.DimBalance{ledg.Available: ledg.FilAmount(burntBalance)},
		builtin3.RewardActorAddr:        ledg.DimBalance{ledg.Available: ledg.FilAmount(rewardBalance)},
		builtin3.StoragePowerActorAddr:  ledg.DimBalance{ledg.Available: ledg.FilAmount(powerBalance)},
		builtin3.StorageMarketActorAddr: ledg.DimBalance{ledg.Available: ledg.FilAmount(marketBalance)},
		l.OriginMsg.To:                 minerBal,
		//l.originMsg.From:               ledg.DimBalance{ledg.Available: ledg.FilAmount(fromBal)},
	}
}


func (l *MinerLedger) InsertActorHead(typ string, msg *types.Message, c *cid.Cid, addr address.Address) {
	//minerAddress:=l.originMsg.To
	minerAddress:=addr
	//if minerAddress.String()!="t01098" {return}

	act,_:=l.st.GetActor(minerAddress)
	var head cid.Cid
	if c==nil { head=act.Head} else {head=*c}
	mgo:=ledg_util.GetOrCreateMongoConnection()
	mgo.InsertLogEntry(bson.M{
		"typ":typ,
		"originMsg":msg.Cid().String(),
		"minerAddress": minerAddress.String(),
		"head":head.String(),
	})
}

func (l *MinerLedger) GetMinerBalanceByAddressDim( address2 address.Address) (ledg.DimBalance,error) {
	minerState := State{}

	act, err := l.st.GetActor(address2)
	totalBalance:=ledg.FilAmount(big.NewInt(0))
	if act!=nil {
		blk, _ := l.cst.Blocks.Get(act.Head)

		oif, _ := l.dump(act, blk.RawData())
		//data, _ := json.MarshalIndent(oif, "", "  ")
		data, _ := json.Marshal(oif)
		json.Unmarshal(data, &minerState)
		totalBalance = ledg.FilAmount(act.Balance)
	}

	if minerState.PreCommitDeposits.Int == nil {	minerState.PreCommitDeposits=big.NewInt(0)}
	if minerState.LockedFunds.Int == nil {	minerState.LockedFunds=big.NewInt(0)}
	if minerState.InitialPledge.Int == nil {	minerState.InitialPledge=big.NewInt(0)}

	//if minerState.Vesting.Int == nil {	minerState.Vesting=big.NewInt(0)}

	return minerState.ToMinerBalance(totalBalance,ledg.FilAmountFromInt(0)), err
}


func Create(ctx context.Context, st *state.StateTree,cst *cbor.BasicIpldStore,
	dump DumpActorStateFunc,epoch abi.ChainEpoch) *MinerLedger {

	l:= &MinerLedger{

		st:           st,
		cst:          cst,
		//Entries:      make(map[string]*ledg.LedgerEntryMongo,0),
		Opening:      ZeroActorBalance(),
		Closing:      ZeroActorBalance(),
		BalancesDiff: ZeroActorBalance(),
		dump:dump,
		epoch: epoch,
		MinerAddress: address.Undef,
	}

	return l
}

//func (l *MinerLedger) fillOpeningBalance() { l.Opening=l.getBalances() }
func (l *MinerLedger) fillClosingBalance( ) {
	l.Closing=l.getBalances()
	for k,v:=range l.Closing{

		l.BalancesDiff[k]=v.Diff(l.Opening[k])
	}
}

func (l *MinerLedger) FillClosingBalance_( ) {

	//isRoot:=parent==nil
	////trc := &rt.executionTrace
	//if !isRoot{
	//	//trc.CallDepth = rt.depth
	//
	//	closing, _ := vm.GetMinerBalanceByAddressDim(rt, currentMsg.To)
	//	//trc.ActorBalanceDiff = closing.Diff(trc.ActorBalanceOpening)
	//	l.appendMsgEntry(currentMsg, trc.ActorBalanceDiff, big.NewInt(0),closing, trc.MsgRct.GasUsed, rt.depth, false)
	//}
}



//func (l *MinerLedger) FinalizeImplicitMsg(ctx context.Context,) {
//	l.Flush(ctx)
//}




//func (l *MinerLedger) FinalizeMsg(ctx context.Context,outputs ledg.GasOutputs,epoch abi.ChainEpoch) {
//
//	//l.fillClosingBalance()
//	//addr,_:=address.NewFromString("t01098")
//	//l.InsertActorHead("Finalize",nil,addr)
//
//	l.AppendRootEntry(ctx,outputs,ledg.ChainEpoch(epoch),0)
//
//	e:=l.Entries[l.currentMsg.Cid().String()]
//	if e!=nil {
//		miner:=l.originMsg.To
//
//		//e.MethodName=l.originMsg.Cid().String()//z_del
//		e.Opening=l.Opening[miner]
//		e.Balance =l.Closing[miner]
//		e.Amount=l.BalancesDiff[miner]
//		//l.Entries[l.currentMsg.Cid().String()]=e
//	}else{
//		panic("CurrentMsg not found in LedgerEntries collection")
//
//		ledg_util.Log(bson.M{
//			"originCid":l.originMsg.Cid().String(),
//			"currMsgToSearch":l.currentMsg.Cid().String(),
//			"from":l.originMsg.From.String(),
//			"to":l.originMsg.To.String(),
//			"method":ledg_util.GetMethodName(l.originMsg),
//			"node":"currentMsg not found",
//			"count": len(l.Entries),
//			"all":"",//s,
//		})
//	}
//	l.Flush(ctx)
//}

//from f00 (sys actor)
//1 msg for each miner in epoch
//1 msg to cron
//f00=>f03=>f04=>miner-confirmSectorsProofsValid

//func (l *MinerLedger) SetOriginMsg(msg *types.Message)  {l.originMsg=msg}
//func (l *MinerLedger) SetCurrentMsg(msg *types.Message) {l.currentMsg=msg}


func (l *MinerLedger) AddVestedFunds(ctx context.Context, msg *types.Message,sendRet []byte,gasUsed int64,callDepth uint64){

}

func (l *MinerLedger) AddEntry(ctx context.Context, msg *types.Message,sendRet []byte,callDepth uint64,note string){

	if ctx.Value("replay")==nil { return }
	if l==nil {panic("MinerLedger is nil")}
	if msg==nil {panic("MinerLedger.AddEntry: msg is nil")}
	if l.OriginMsg==nil {panic("MinerLedger.AddEntry originMsg is nil")}



	amount:= ledg.DimBalance{ledg.Available: ledg.FilAmount(msg.Value)}
	sectorNumber := ledg.SectorNumber(0)
	minerAddress:=ledg.Address{}
	if ledg_util.IsMiner(l.OriginMsg.To) {
		s:= l.getSectorFromParams(l.OriginMsg)
		sectorNumber=ledg.SectorNumber(s.Number)
		minerAddress=ledg.Address{l.OriginMsg.To}
	}

	//actorBalances:=l.getBalances()
	//minerBalance:=actorBalances[msg.To]

	//totalAmount:= big.Add(big.Add(msg.Amount,gasFee),minerTip).Neg()
	if note=="" {note="GO:AddEntry"}

	e := l.fillMsgFields(msg) //, address2, offset, amount, minerBalance, callDepth, epoch, sectorId)
	e.Id			= 	msg.Cid().String()
	e.EntryCid		=   ledg.Cid{msg.Cid()}
	e.MsgCid		=   ledg.Cid{l.OriginMsg.Cid()}

	e.TotalAmount	=	ZeroDimBalance()
	e.Opening		=	ZeroDimBalance()
	e.Amount		=	amount
	e.Balance 		=	ZeroDimBalance()
	e.GasFee		=   ledg.FilAmount{}
	e.MinerTip		=	ledg.FilAmount{}
	e.GasUsed		=	0
	e.CallDepth		=	callDepth

	e.Invariant =	ledg.FilAmount{}
	e.Epoch			=	l.epoch
	e.EntryType		=	""
	e.MethodName	=   ledg_util.GetMethodName(msg)

	e.Miner			=	minerAddress
	e.SectorNumber	=	sectorNumber

	//eId := e.Id + strconv.Itoa(len(l.Entries))
	e.Note			=	note
	//l.Entries[e.Id] = &e
}

func (l *MinerLedger) fillMsgFields(msg *types.Message) ledg.LedgerEntryMongo {
	address2:=ledg.Address{msg.To};
	offset:=ledg.Address{msg.From}

	e := ledg.LedgerEntryMongo{
		//Id:      msg.Cid().String(),

		Version :msg.Version,
		Address: address2,//msg.to
		Offset:  offset,//msg.From

		Nonce : msg.Nonce,

		Value:           ledg.FilAmount(msg.Value),//msg.value

		GasLimit   :msg.GasLimit,
		GasFeeCap  :ledg.FilAmount(msg.GasFeeCap),
		GasPremium :ledg.FilAmount(msg.GasPremium),

		Method : msg.Method,
		Params :msg.Params,

		//Epoch:      epoch,
		//TargetActorType: 0,
		//CallDepth:  callDepth,
		MsgCid:      ledg.Cid{},
		EntryCid:    ledg.Cid{},
		TotalAmount: nil,
		Opening:     nil,
		Amount:      nil,
		Balance:     nil,
		//Amount:     amount,
		//Balance:    minerBalance,
		GasFee:     ledg.FilAmount(big.NewInt(0)),
		MinerTip:   ledg.FilAmount{},
		GasUsed:    0,
		CallDepth:  0,
		Implicit: l.IsImplicit,
	}
	return e
}
func (l *MinerLedger) LockPrecommitDeposit( msg *types.Message,sendRet []byte,callDepth uint64){

	//
	//if ctx.Amount("replay")==nil { return }
	//if l==nil {panic("MinerLedger is nil")}
	//if msg==nil {panic("MinerLedger.LockPrecommitDeposit currentMsg is nil")}
	//if l.OriginMsg==nil {panic("MinerLedger.LockPrecommitDeposit originMsg is nil")}
	//
	//if l==nil { return }


	sectorNumber := ledg.SectorNumber(0)
	minerAddress:=ledg.Address{}
	if ledg_util.IsMiner(l.OriginMsg.To) {
		s:= l.getSectorFromParams(l.OriginMsg)
		sectorNumber=ledg.SectorNumber(s.Number)
		minerAddress=ledg.Address{l.OriginMsg.To}
	}

	amount:= ledg.DimBalance{ledg.Available: ledg.FilAmount(msg.Value)}

	//actorBalances:=l.getBalances()
	//minerBalance:=actorBalances[msg.To]
	minerBalance:=ZeroDimBalance()

	e:=l.fillMsgFields(msg)

	e.Id			= 	msg.Cid().String()
	e.EntryCid		=   ledg.Cid{msg.Cid()}
	e.MsgCid		=   ledg.Cid{l.OriginMsg.Cid()}

	e.TotalAmount	=	ZeroDimBalance()
	e.Opening		=	ZeroDimBalance()
	e.Amount		=	amount
	e.Balance 		=	minerBalance
	e.GasFee		=   ledg.FilAmountFromInt(0)
	e.MinerTip		=	ledg.FilAmountFromInt(0)
	e.GasUsed		=	0
	e.CallDepth		=	callDepth

	e.Invariant 	=	ledg.FilAmountFromInt(-1)
	e.Epoch			=	l.epoch
	e.EntryType		=	ledg.SectorEntryType.PreCommit
	e.MethodName	=   ledg_util.GetMethodName(msg)

	e.Miner			=	minerAddress
	e.SectorNumber	=	sectorNumber

	e.EntryType=ledg.SectorEntryType.PreCommit

	//eId := e.Id //+ strconv.Itoa(len(l.Entries))
	e.Note			=	"GO:LockPrecommitDeposit"

	//l.Entries[e.Id] = &e
}

func (l *MinerLedger) Flush(ctx context.Context){
	if ctx.Value("replay")==nil {return}
	if (l==nil) {return}

	//for i,_:=range l.Entries{
	//	InsertEntry(l.Entries[i])
	//}
	//l.Entries=make(map[string]*ledg.LedgerEntryMongo,0)
}


//func (l *MinerLedger) Export() []types.LedgerEntryExport{
//	ret:=make([]types.LedgerEntryExport,len(l.Entries))
//	//for i,e:= range l.Entries{  ret[i]=e.Export() }
//	return  ret
//}

func ZeroDimBalance() ledg.DimBalance {
	ret:=make (map[int]ledg.FilAmount)

	//for i:=range []int{Available,PreCommitDeposits,InitialPledge,LockedFunds,Vesting}{ ret[i]=ZeroAmount()	}
	ret[ledg.Available]= ZeroAmount()
	return ret
}

func ZeroActorBalance() ActorBalances {
	return ActorBalances{
		builtin.RewardActorAddr :       ZeroDimBalance(), //f02
		builtin.StoragePowerActorAddr:  ZeroDimBalance(), //f04
		builtin.StorageMarketActorAddr: ZeroDimBalance(), //f05
		builtin.BurntFundsActorAddr:    ZeroDimBalance(), //f099
	}
}

func ZeroAmount() ledg.FilAmount {
	return ledg.FilAmount(big.NewInt(0))
}
func (state State) ToMinerBalance(totalBalance ledg.FilAmount,vestingFunds ledg.FilAmount) ledg.DimBalance {

	deposit:=ledg.FilAmount(state.PreCommitDeposits)
	pledge:=ledg.FilAmount(state.InitialPledge)
	locked:=ledg.FilAmount(state.LockedFunds)
	debt:=ledg.FilAmount(state.FeeDebt)
	available:=ledg.FilAmount{
		totalBalance.
			Neg(deposit.Int).
			Neg(pledge.Int).
			Neg(locked.Int).
			Neg(vestingFunds.Int).
			Neg(debt.Int),
	}
	ret:=ledg.DimBalance{
		ledg.Available:         available,
		ledg.PreCommitDeposits: deposit,
		ledg.InitialPledge:     pledge,
		ledg.LockedFunds:       locked,
		ledg.Vesting:           vestingFunds,
		ledg.FeeDebt: 		   debt,
	}

	return  ret
}
func (l *MinerLedger) MsgInvariant() ledg.FilAmount{
	//inv:=ledg.FilAmountFromInt(0)
	//
	//for _,e:=range l.Entries {
	//	_=e.TotalAmount[0]
	//}
	return ledg.FilAmountFromInt(0)
}

func (l *MinerLedger)  AppendRootEntry(ctx context.Context, msg * types.Message,gasOutput *ledg.GasOutputs,callDepth uint64){

	if msg==nil {panic("MinerLedger.AppendRootEntry msg is nil")}
	if l.OriginMsg==nil {panic("MinerLedger.AppendRootEntry originMsg is nil")}
	if &gasOutput==nil {panic("GasOutput is nil")}
	if gasOutput.BaseFeeBurn.Int==nil{panic("gasFee is nil")}
	if gasOutput.MinerTip.Int==nil{panic("minerTip is nil")}
	if msg.Value.Int==nil{panic("msg.Amount is nil")}

	gasFee:=gasOutput.BaseFeeBurn
	minerTip:=gasOutput.MinerTip


	totalAmount:= big.Add(big.Add(msg.Value,gasFee),minerTip).Neg()
	methodName:=ledg_util.GetMethodName(l.OriginMsg)

	e:= ledg.LedgerEntryMongo{
		Id:          msg.Cid().String()+"-root",
		EntryCid:    ledg.Cid{msg.Cid()},
		MsgCid:      ledg.Cid{l.OriginMsg.Cid()},
		Address:     ledg.Address{msg.From},
		Offset:      ledg.Address{msg.To},
		TotalAmount: ledg.DimBalance{ledg.Available: ledg.FilAmount(totalAmount)},
		Amount:      ZeroDimBalance(),

		Opening: ZeroDimBalance(),
		Balance: ZeroDimBalance(),

		GasFee:     ledg.FilAmount(gasFee.Neg()),
		GasUsed:    gasOutput.GasUsed,
		MinerTip:   ledg.FilAmount(minerTip.Neg()),
		CallDepth:  callDepth,
		Epoch:      l.epoch,
		Method:     msg.Method,
		Value:      ledg.FilAmount(msg.Value.Neg()),
		MethodName: methodName,
		Note:"root for "+methodName,
		EntryType: ledg_util.GetEntryType(msg),
	}
	fmt.Println(e)
	//l.Entries[e.Id]=&e
}



