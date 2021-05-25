package Posting

import (
	"bytes"
	"github.com/filecoin-project/go-state-types/big"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	m "github.com/filecoin-project/lotus/chain/vm/ledger/models"
	"github.com/filecoin-project/specs-actors/actors/builtin/power"
)

//func ( l * LedgerPosting) minerEntrySendTemplate(p ledg_util.ActorMethodParams,idx int32) m.LedgerEntry{
//	e:=l.minerEntryTemplate(p,idx)
//	e.MethodName=e.MethodName+"#Send"
//	amount:=ledg.FilAmount(big.NewFromGo(p.Msg.Value.Int).Neg())
//
//	addressId,_:=address.IDFromAddress(p.Msg.From)
//	offsetId,_:=address.IDFromAddress(p.Msg.To)
//	e.AddressId=int32(addressId)
//	e.OffsetId=int32(offsetId)
//
//	e.Amount =amount
//	return e
//}

func ( l *LedgerPosting) minerEntryTemplate(p ledg_util.ActorMethodParams,idx int32, initialEntry bool) m.LedgerEntry{

	if idx==0 {		idx=int32(len(l.minerEntries))	}

	var addressId, offsetId int32
	var methodName string
	var amount ledg.FilAmount
	if initialEntry {
		methodName="=> "+ledg_util.GetMethodName(p.Msg)
		addressId,_=ledg_util.GetOrCreateAccountFromAddress(p.Msg.From,"",l.Epoch)
		offsetId,_=ledg_util.GetOrCreateAccountFromAddress(p.Msg.To,"",l.Epoch)
		amount=ledg.FilAmount(big.NewFromGo(p.Msg.Value.Int).Neg())

	} else {
		methodName=ledg_util.GetMethodName(p.Msg)
		addressId, _ = ledg_util.GetOrCreateAccountFromAddress(p.Msg.To, "", l.Epoch)
		offsetId, _ = ledg_util.GetOrCreateAccountFromAddress(p.Msg.From, "", l.Epoch)
		amount=ledg.FilAmount(big.NewFromGo(p.Msg.Value.Int))
	}

	e:=m.LedgerEntry{
		Epoch: int32(l.Epoch),
		Id:         idx,
		//Version:    0,
		AddressId:  int32(addressId),

		//Address:    ,
		OffsetId:   int32(offsetId),
		//Nonce:      msg.Nonce,
		Amount:     amount,
		//GasLimit:   0,
		//GasFeeCap:  ledg.FilAmount{},
		//GasPremium: ledg.FilAmount{},
		Method:     p.Msg.Method,
		//Params:     nil,
		//GasFee:     ledg.FilAmount{},
		//MinerTip:   ledg.FilAmount{},
		//GasUsed:    0,
		CallDepth:  p.Depth,
		SectorId:   nil,
		Sector:     nil,
		Miner:      m.Account{},
		EntryType:  "",
		MethodName: methodName,
		//Note:       "",
		Implicit:   false,
		TxId: l.CurrentTxId,
	}
	return e

}
func(l *LedgerPosting) MinerActorConstructor(p ledg_util.ActorMethodParams) {

	msg:=p.Msg
	methReturn:=p.Ret

	llog.Info("LedgerPosting.MinerActorConstructor")
	cm		:=	power.CreateMinerParams{}
	cm_ret	:=	power.CreateMinerReturn{}

	cm.UnmarshalCBOR(bytes.NewReader(msg.Params))
	cm_ret.UnmarshalCBOR(bytes.NewReader(methReturn))

	//addrId,_:=address.IDFromAddress(cm_ret.IDAddress)
	//minerActor:=m.Account{
	//	ID:            int32(addrId),
	//	Address:       cm_ret.RobustAddress.String(),
	//
	//	Name: "Miner: "+cm_ret.IDAddress.String(),
	//
	//	CreationEpoch: l.Epoch,
	//	//Balance:       ledg.DimBalance{
	//	//	ledg.Available:ledg.FilAmountFromInt(11),
	//	//	ledg.InitialPledge: ledg.FilAmountFromInt(22),
	//	//},
	//	//SectorCounts: ledg.SectorCounts{Active: 11},
	//	//PowerBalance:ledg.PowerBalance{VerifiedStoragePower: ledg.StoragePowerFromInt(123)},
	//	//ActorTypeConst: models.Miner,
	//	//TotalReward: ledg.FilAmountFromInt(123),
	//	//MessagesCount: 999,
	//
	//	//Stats: ledg.MiningStats{
	//	//	PowerGrowth:           ledg.StoragePowerFromInt(32*1024*1024),
	//	//	BlocksMined:           8,
	//	//	MiningEfficiencyPerTB: ledg.FilAmountFromInt(100025),
	//	//	WinCount:              10,
	//	//	MinerEquivalent:       15.89,
	//	//},
	//	//Properties: ledg.MinerProperties{
	//	//	PeerId:  "12D3KooWRudzcMVAgZapWJXPDKDgUMZbbsQFeDHcAFQwxAWrtQTV",
	//	//	Owner:   ledg.NewAddressFromString("f3wkx7jksblo4kehbklknlivm6pniartluv3nqz3mpwjj5dyfu55pctvyxxejkjgki7qp3r3thxt3wk73hwsua"),
	//	//	Worker:  ledg.NewAddressFromString("f3rzvyvt6lnamvx7dc4ulrukenudq7ywzrdjelnpouwzsurqnep5vcick7x72w4tslmyvqbx2mkemkqalbtswq"),
	//	//	Region:  "europe",
	//	//	Country: "lv",
	//	//	Ip:      "8.8.8.8",
	//	//},
	//}
	//
	//l.insert(&minerActor,true)

	e:=l.minerEntryTemplate(p,0,false)
	e.MethodName="MinerActorCtor"
	e.AddressId=0
	e.TxId=l.CurrentTxId
	e.CallDepth=p.Depth
	l.insert(&e,true)
}


func(l *LedgerPosting) ProveCommitSector( p ledg_util.ActorMethodParams) {

	//l.update sector (statusProved)
	//l.postCommit() //bulk from cron call
	//l.UnlockPrecommitDeposit
	//l.addSectorPledge

	//for deal :=deals2activate{
	//	l.ActivateDeals()
	//	l.postDealActive()
	//}
	//l.unlovlSectorDeposit

}

func(l *LedgerPosting)ControlAddresses(p ledg_util.ActorMethodParams)         {} //2
func(l *LedgerPosting)ChangeWorkerAddress(p ledg_util.ActorMethodParams)      {} //3
func(l *LedgerPosting)ChangePeerID(p ledg_util.ActorMethodParams)             {} //4
func(l *LedgerPosting)SubmitWindowedPoSt(p ledg_util.ActorMethodParams)       {} //5
func(l *LedgerPosting)ExtendSectorExpiration(p ledg_util.ActorMethodParams)   {} //8
func(l *LedgerPosting)TerminateSectors(p ledg_util.ActorMethodParams)         {} //9
func(l *LedgerPosting)DeclareFaults(p ledg_util.ActorMethodParams)            {} //10
func(l *LedgerPosting)DeclareFaultsRecovered(p ledg_util.ActorMethodParams)   {} //11
func(l *LedgerPosting)OnDeferredCronEvent(p ledg_util.ActorMethodParams)      {} //12
func(l *LedgerPosting)CheckSectorProven(p ledg_util.ActorMethodParams)        {} //13
func(l *LedgerPosting)ApplyRewards(p ledg_util.ActorMethodParams)             {} //14
func(l *LedgerPosting)ReportConsensusFault(p ledg_util.ActorMethodParams)     {} //15
func(l *LedgerPosting)WithdrawBalance(p ledg_util.ActorMethodParams)          {} //{/16
func(l *LedgerPosting)ConfirmSectorProofsValid(p ledg_util.ActorMethodParams) {} //{/17
func(l *LedgerPosting)ChangeMultiaddrs(p ledg_util.ActorMethodParams)         {} //18
func(l *LedgerPosting)CompactPartitions(p ledg_util.ActorMethodParams)        {} //19
func(l *LedgerPosting)CompactSectorNumbers(p ledg_util.ActorMethodParams)     {} //20
func(l *LedgerPosting)ConfirmUpdateWorkerKey(p ledg_util.ActorMethodParams)   {} //21
func(l *LedgerPosting)RepayDebt(p ledg_util.ActorMethodParams)                {} //22
func(l *LedgerPosting)ChangeOwnerAddress(p ledg_util.ActorMethodParams)       {} //23
func(l *LedgerPosting)DisputeWindowedPoSt(p ledg_util.ActorMethodParams)      {} //{/24