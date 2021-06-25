package stateTransition

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	ledg_types "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	logging "github.com/ipfs/go-log/v2"
)





var llog			= logging.Logger("gledger")


func (st *StateTrans)PrintStateTrans(){

	//ledg_util.Llogf("minerState %s ",st.ActorStates)
	//ledg_util.Llogf("minerState2 %s ",st.ActorStatesOnEnd)
}


func  (stinfo *StateTrans) GetSectorPreCommitOnChainInfo( ctx context.Context,minerAddr address.Address,sectorNum abi.SectorNumber) (*miner.SectorPreCommitOnChainInfo, error) {

	addrId,_:=address.IDFromAddress(minerAddr)

	ledg_util.Llogf("AddrId %d",addrId)

	astates:=stinfo.ActorStates[abi.ActorID(addrId)]
	ledg_util.Llogf("ActionStates %s",astates)
	mstate:=astates.minerState
	ledg_util.Llogf("MinerState %s",mstate)
	st:=stinfo.ActorStates[abi.ActorID(addrId)].minerState[ledg_types.Closing]
	ledg_util.Llogf("GetSectorPreCommitOnChainInfo. Actor %s state %s",minerAddr.String(),st)
	sector,err:=st.GetPrecommittedSector(sectorNum)

	return sector, err
}
func  (stinfo *StateTrans)GetSectorOnChainInfo( ctx context.Context,minerAddr address.Address,sectorNum abi.SectorNumber) (*miner.SectorOnChainInfo, error) {

	//stinfo.LoadActorState(ctx,minerAddr)
	addrId,_:=address.IDFromAddress(minerAddr)
	st:=stinfo.ActorStates[abi.ActorID(addrId)].minerState[ledg_types.Closing]
	sector,err:=st.GetSector(sectorNum)
	return sector, err
}






//func (stinfo *StateTrans) SetState(ctx context.Context) {
//
//	addr,_:=address.NewIDAddress(1002)
//
//	stinfo.LoadActorState(ctx,addr)
//	stinfo.LoadVestingFunds(0)
//
//	//ledg_util.Llogf("Vesting sch 1002 %s",stinfo.VestingFunds)
//
//
//}