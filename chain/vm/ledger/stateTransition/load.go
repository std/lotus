package stateTransition

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	miner0 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	miner2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	miner3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	miner4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)


func (stinfo *StateTrans) CheckActorState(ctx context.Context , actorAddr address.Address) {
	_actorId,_:=address.IDFromAddress(actorAddr)
	actorId:=abi.ActorID(_actorId)

	actorState:=stinfo.ActorStates[actorId]

	_,ok:=actorState.Actor[ledg.Opening]
	if !ok {ledg_util.Llogf("Check ActorState %s Actor is nil",stateKindStr(ledg.Opening))}
	_,ok=actorState.Actor[ledg.Closing]
	if !ok {ledg_util.Llogf("Check ActorState %s Actor is nil",stateKindStr(ledg.Closing))}

	_,ok=actorState.minerState[ledg.Opening]
	if !ok {ledg_util.Llogf("Check ActorState %s MinerState is nil",stateKindStr(ledg.Opening))}
	_,ok=actorState.minerState[ledg.Closing]
	if !ok {ledg_util.Llogf("Check ActorState %s MinerState is nil",stateKindStr(ledg.Closing))}

	vCid:=actorState.VestingCid[ledg.Opening]
	if !vCid.Defined() {ledg_util.Llogf("Check ActorState %s VestingCid is Empty",stateKindStr(ledg.Closing))}
	vCid=actorState.VestingCid[ledg.Closing]
	if !vCid.Defined() {ledg_util.Llogf("Check ActorState %s VestingCid is Empty",stateKindStr(ledg.Closing))}

	_,ok=actorState.VestingFunds[ledg.Opening]
	if !ok {ledg_util.Llogf("Check ActorState %s VestingFunds is Empty",stateKindStr(ledg.Closing))}
	_,ok=actorState.VestingFunds[ledg.Closing]
	if !vCid.Defined() {ledg_util.Llogf("Check ActorState %s VestingFunds is Empty",stateKindStr(ledg.Closing))}

}
func (stinfo *StateTrans) LoadActorState(ctx context.Context , actorAddr address.Address, stateKind ledg.OpeningOrClosing) {


	if stinfo.Stree==nil {panic ("LoadActorState: StateTree is nil")}

	//ledg_util.Llogf("Begin load ActorState for %s",actorAddr)
	ipldStore:=stinfo.Stree.Store
	wrappedStore:=adt.WrapStore(ctx,ipldStore)

	stinfo.Store=wrappedStore

	_actorId,_:=address.IDFromAddress(actorAddr)
	if _actorId==4 {
		return
	}

	if stinfo.ActorStates==nil {panic("Panic in LoadActorState: ActorStates outer map is nil")}

	actorId:=abi.ActorID(_actorId)
	actor,_:=stinfo.Stree.GetActor(actorAddr)

	var s ActorStateTrans
	s,ok:=stinfo.ActorStates[actorId]
	if !ok {
		s=ActorStateTrans{
			Actor:          make(map[ledg.OpeningOrClosing]types.Actor),
			minerState:     make(map[ledg.OpeningOrClosing]miner.State),
			VestingCid:     make(map[ledg.OpeningOrClosing]cid.Cid),
			VestingFunds:   make(map[ledg.OpeningOrClosing]VestingFunds),
			actorId:        actorId,
			VestingEntries: make(map[abi.ChainEpoch]models.VestingEntry),
		}
	}

	s.Actor[stateKind]=*actor
	s.actorId=actorId

	if mstate, err:= miner.Load(wrappedStore, actor);err!=nil{
		ledg_util.Llogf("%s",err.Error())
	}else {
		s.minerState[stateKind]=mstate

	}

	vestingCid,_:=stinfo.getCidsFromStateVer(wrappedStore,actor.Code,actor.Head,actorId, stateKind)
	s.VestingCid[stateKind]=vestingCid

	if VestingFunds,err:=stinfo.getVestingFunds(actorId, stateKind,vestingCid);err!=nil {
		ledg_util.Llogf("LoadVestingFunds Err: %s",err.Error())
	}else {
		s.VestingFunds[stateKind]=*VestingFunds
	}
	stinfo.ActorStates[actorId]=s
}

func stateKindStr(stateKind ledg.OpeningOrClosing)string{
	openOrClose:="Closing"
	if stateKind {openOrClose="Opening"}
	return openOrClose
}

func (trans *StateTrans) getCidsFromStateVer(store adt.Store, actCode, actHead cid.Cid, actorId abi.ActorID, stateKind ledg.OpeningOrClosing) (cid.Cid,error){
	switch actCode {

	case builtin0.StorageMinerActorCodeID:
		if state,err:=load0(store,actHead); err!=nil {return cid.Undef,err} else {return state.VestingFunds,nil}
	case builtin2.StorageMinerActorCodeID:
		if state,err:=load2(store,actHead); err!=nil {return cid.Undef,err} else {return state.VestingFunds,nil}
	case builtin3.StorageMinerActorCodeID:
		if state,err:=load3(store,actHead); err!=nil {return cid.Undef,err} else {return state.VestingFunds,nil}
	case builtin4.StorageMinerActorCodeID:
		if state,err:=load4(store,actHead); err!=nil {return cid.Undef,err} else {return state.VestingFunds,nil}
	default:
		return cid.Undef,nil
	}
}

func (st *StateTrans) getVestingFunds(actorId abi.ActorID, stateKind ledg.OpeningOrClosing,vestingCid cid.Cid) ( *VestingFunds,error) {

	//ledg_util.Llogf("Vesting sch  %d",actorId)
	var funds VestingFunds


	if err := st.Store.Get(st.Store.Context(),vestingCid, &funds); err != nil {
		return  nil,xerrors.Errorf("failed to load vesting funds (%s): %w", vestingCid, err)
	}
	return  &funds,nil
}

func load0(store adt.Store, root cid.Cid) (*miner0.State, error) {
	//out := state3{store: store}
	var out miner0.State
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
func load2(store adt.Store, root cid.Cid) (*miner2.State, error) {
	//out := state3{store: store}
	var out miner2.State
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
func load3(store adt.Store, root cid.Cid) (*miner3.State, error) {
	//out := state3{store: store}
	var out miner3.State
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func load4(store adt.Store, root cid.Cid) (*miner4.State, error) {
	//out := state3{store: store}
	var out miner4.State
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}


