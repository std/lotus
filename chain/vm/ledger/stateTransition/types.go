package stateTransition

import "C"
import (
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	ledg_types "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type ActorStatesKey struct {
	actorId abi.ActorID
	stateKind ledg_types.OpeningOrClosing
}

type StateTrans struct {

	Stree *state.StateTree
	Ts *types.TipSet
	//minerState,minerState2 miner.State

	Store adt.Store


	ActorStates map[abi.ActorID]ActorStateTrans
	//ActorStates map[ActorStatesKey]ActorStateTrans
}


type ActorStateTrans struct{

	Actor      map[ledg_types.OpeningOrClosing]types.Actor
	minerState map[ledg_types.OpeningOrClosing]miner.State


	VestingCid     map[ledg_types.OpeningOrClosing]cid.Cid
	VestingFunds   map[ledg_types.OpeningOrClosing]VestingFunds

	actorId        abi.ActorID
	VestingEntries map[abi.ChainEpoch]models.VestingEntry
}

type amountTuple struct {
	start,end abi.TokenAmount
}




func (st *ActorStateTrans) State(stateKind ledg_types.OpeningOrClosing) (miner.State,error){


	if st.minerState!=nil {return st.minerState[stateKind], nil}

	return nil, xerrors.Errorf("Miner %d,  minerState is nil",st.actorId)
}

func (st *ActorStateTrans)String()string{
	balance:=miner.LockedFunds{}

	balance, _ = st.minerState[ledg_types.Closing].LockedFunds()

	ret:=fmt.Sprintf("State Actor %s Balance %s vestingCid %s",st.actorId,balance,st.VestingCid[ledg_types.Closing])
	return ret
}

type VestingFunds struct {
	Funds []VestingFund
}
type VestingFund struct {
	Epoch  abi.ChainEpoch
	Amount abi.TokenAmount
}
func (f *VestingFund) String()string{
	return fmt.Sprintf("Epoch %d Amount %s",f.Epoch,f.Amount.String())
}


