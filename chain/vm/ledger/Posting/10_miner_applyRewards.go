package Posting

import (
	"bytes"
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	ledg_types "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

//github.com/filecoin-project/specs-actors/v3 v3.1.0 h1:s4qiPw8pgypqBGAy853u/zdZJ7K9cTZdM1rTiSonHrg=
//github.com/filecoin-project/specs-actors/v3 v3.1.0/go.mod h1:mpynccOLlIRy0QnR008BwYBwT9fen+sPR13MA1VmMww=

func(l *LedgerPosting)ApplyRewards(p ledg_util.ActorMethodParams)             {

	_actorId,_:=address.IDFromAddress(p.Msg.To)
	actorId:=abi.ActorID(_actorId)

	//rewardActorId:=abi.ActorID(2)

	l.Stinfo.LoadActorState(context.TODO(),p.Msg.To,ledg_types.Closing)
	l.Stinfo.CheckActorState(context.TODO(),p.Msg.To)
	var applyRewardParams builtin.ApplyRewardParams

	applyRewardParams.UnmarshalCBOR(bytes.NewReader(p.Msg.Params))
	reward:=applyRewardParams.Reward
	penalty:=applyRewardParams.Penalty

	e:=l.minerEntryTemplate(p,int32(len(l.minerEntries)),true)
	e.MethodName="reward transfer"
	l.insert(e,false)


	//////

	st:=l.Stinfo.ActorStates[actorId]

	//actorBalance0:=st.Actor[ledg_types.Opening].Balance
	//actorBalance:=st.Actor[ledg_types.Closing].Balance


	closingState,_:=st.State(ledg_types.Closing)
	lockedBalance,_:=closingState.LockedFunds()
	vesting:=lockedBalance.VestingFunds

	closingState0,_:=st.State(ledg_types.Opening)
	lockedBalance0,_:=closingState0.LockedFunds()
	vesting0:=lockedBalance0.VestingFunds



	//e_r:=l.minerEntryTemplate(p,int32(len(l.minerEntries)),false)
	//e_r.Amount=ledg_types.FilAmount(reward)
	//e_r.MethodName="reward entry"
	//e_r.Balance0=ledg_types.FilAmount(actorBalance0)
	//e_r.Balance=ledg_types.FilAmount(actorBalance)
	//l.insert(e_r,false)


	//lock reward entry
	lockedReward:=big.Sub(vesting,vesting0)
	availReward:=big.Sub(reward,lockedReward)
	e_lock:=l.minerEntryTemplate(p,int32(len(l.minerEntries)),false)
	e_lock.Amount=ledg_types.FilAmount(lockedReward)
	e_lock.Balance=ledg_types.FilAmount(vesting)
	e_lock.Balance0=ledg_types.FilAmount(vesting0)
	e_lock.DimensionId=ledg_types.LockedFunds

	e_lock.MethodName="lock reward"
	l.insert(e_lock,false)

	e_lock.Id++
	e_lock.Amount=ledg_types.FilAmount(availReward)
	e_lock.DimensionId=0
	e_lock.MethodName="avail reward"
	l.insert(e_lock,false)



	ledg_util.Llogf("Apply Reward %s penalty %s",reward,penalty)

	addrId,_:=address.IDFromAddress(p.Msg.To)





	//st1:=l.Stinfo.ActorStates[abi.ActorID(addrId)]
	//st2:=l.Stinfo.ActorStates[abi.ActorID(addrId)]
	ledg_util.Llogf("ApplyRewards for actor %s",p.Msg.To.String())


	//istate,_:=st.State()
	//ledg_util.Llogf("state %s",istate)
	//ledg_util.Llogf("ApplyRewards:VestingCid onStart %s",st1.VestingCid[ledg_types.Opening].String())
	//ledg_util.Llogf("ApplyRewards:VestingCid OnEnd %s",st2.VestingCid[ledg_types.Closing].String())
	//ledg_util.Llogf("ApplyRewards:VestingFunds2 len %s", len(st1.VestingFunds[ledg_types.Opening].Funds))
	//ledg_util.Llogf("ApplyRewards:VestingFunds2 len %s", len(st2.VestingFunds[ledg_types.Closing].Funds))

	//for i,v:=range st1.VestingFunds[ledg_types.Opening].Funds{
	//	ledg_util.Llogf("ApplyRewards:VestingFund %d # %s",i,v.String())
	//}
	//
	//for i,v:=range st2.VestingFunds[ledg_types.Closing].Funds{
	//	ledg_util.Llogf("ApplyRewards:VestingFund %d # %s",i,v.String())
	//}
	l.Stinfo.VestingFundsDiff(p.Msg.To,l.Epoch)
	for _,e:=range l.Stinfo.ActorStates[abi.ActorID(addrId)].VestingEntries{
		l.insert(&e,false)
	}

} //14
