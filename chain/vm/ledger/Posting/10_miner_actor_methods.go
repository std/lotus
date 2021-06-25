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
	var methodName,entryType string
	var amount ledg.FilAmount
	if initialEntry {
		methodName="=> "+ledg_util.GetMethodName(p.Msg)
		entryType="transfer"
		addressId,_=ledg_util.GetOrCreateAccountFromAddress(p.Msg.From,"",l.Epoch)
		offsetId,_=ledg_util.GetOrCreateAccountFromAddress(p.Msg.To,"",l.Epoch)
		amount=ledg.FilAmount(big.NewFromGo(p.Msg.Value.Int).Neg())

	} else {
		entryType=""
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

		//AddressMongo:    ,
		OffsetId:   int32(offsetId),
		//Nonce:      msg.Nonce,
		Amount:     amount,
		Balance0: ledg.FilAmountFromInt(0),
		Balance: ledg.FilAmountFromInt(0),
		//GasLimit:   0,
		//GasFeeCap:  ledg.FilAmount{},
		//GasPremium: ledg.FilAmount{},
		MethodId:     int16(p.Msg.Method),
		//Params:     nil,
		//GasFee:     ledg.FilAmount{},
		//MinerTip:   ledg.FilAmount{},
		//GasUsed:    0,
		CallDepth:  int16(p.Depth),
		SectorId:   nil,
		Sector:     nil,
		Miner:      m.Account{},
		EntryType:  ledg.EntryTypeConst(entryType),
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

	cm		:=	power.CreateMinerParams{}
	cm_ret	:=	power.CreateMinerReturn{}

	cm.UnmarshalCBOR(bytes.NewReader(msg.Params))
	cm_ret.UnmarshalCBOR(bytes.NewReader(methReturn))

	e:=l.minerEntryTemplate(p,0,false)
	e.MethodName="MinerActorCtor"
	e.AddressId=0
	e.TxId=l.CurrentTxId
	e.CallDepth=int16(p.Depth)
	e.MinerId=l.MinerId
	e.SectorId=l.SectorId
	l.insert(e,false)
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