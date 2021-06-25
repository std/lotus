package stateTransition

import (
	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
)
type VestingEntryType int16
const (
	LockFunds VestingEntryType= iota
	UnlockVestedFunds
	UnlockUnvestedFunds
)

func (st *StateTrans) VestingFundsDiff(actorAddress addr.Address,processingEpoch abi.ChainEpoch) {
	_actorId, _ := addr.IDFromAddress(actorAddress)
	actorId:=abi.ActorID(_actorId)
	fundsOnStart := st.ActorStates[actorId].VestingFunds[ledg.Opening].Funds
	fundsOnEnd := st.ActorStates[actorId].VestingFunds[ledg.Closing].Funds

	m := make(map[abi.ChainEpoch]amountTuple)

	for _, v := range fundsOnStart { m[v.Epoch]=amountTuple{start: v.Amount} }
	for _, v := range fundsOnEnd {
		entry:=m[v.Epoch]
		entry.end= v.Amount
		m[v.Epoch]=entry
	}

	for epoch,v := range m {
		if v.start.IsZero() {
			st.AddVestingEntry(actorId,v.end,v.start,v.end,epoch,LockFunds,processingEpoch)
		} else if v.end.IsZero() && epoch<=processingEpoch {
			st.AddVestingEntry(actorId,v.end,v.start,v.end,epoch,UnlockVestedFunds,processingEpoch)
		}else if !v.start.Equals(v.end){
			st.AddVestingEntry(actorId,big.Sub(v.end,v.start),v.start,v.end,epoch,LockFunds,processingEpoch)
		}else{
			//st.AddVestingEntry(actorId,v.end,epoch,UnlockUnvestedFunds,processingEpoch)
		}
	}
}


func (st *StateTrans) AddVestingEntry(actorId abi.ActorID,amount,amount1,amount2 abi.TokenAmount,epoch abi.ChainEpoch,
	entryType VestingEntryType,processingEpoch abi.ChainEpoch,

	){

	if entryType!=LockFunds {
		amount=amount.Abs().Neg()
	}

	v:=st.ActorStates[actorId]


	e:=models.VestingEntry{
		MinerId:   int32(v.actorId),
		Epoch:     int32(epoch),
		EntryType: int16(entryType),
		Modified:  int32(processingEpoch),
		Amount: ledg.FilAmount(amount),
		Amount1: ledg.FilAmount(amount1),
		Amount2: ledg.FilAmount(amount2),
		AmountStr:ledg.FilAmount(big.Div(amount,big.NewInt( 1_000_000_000_000_000_000))) ,
	}
	v.VestingEntries[epoch]=e

}

