package stmgr

import (
	"context"
	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/chain/vm/ledger/Posting"
	logging "github.com/ipfs/go-log/v2"

	"github.com/ipfs/go-cid"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"
)


var llog			= logging.Logger("gledger")

func (sm *StateManager) computeTipSetStateEx(ctx context.Context, ts *types.TipSet, cb ExecCallback,epoch abi.ChainEpoch) (cid.Cid, cid.Cid, error) {
	ctx, span := trace.StartSpan(ctx, "computeTipSetState")
	defer span.End()

	//st,_:=sm.ParentState(ts)

	adt:=sm.cs.ActorStore(ctx)

	gl:= Posting.CreateGL(ctx,adt,ts,epoch)
	ctx=context.WithValue(ctx,"gl",gl)
	//
	blks := ts.Blocks()
	//
	//minerAddr,_:=addr.NewFromString("t08557")
	//actor,_:=st.GetActor(minerAddr)
	//llog.Infof("actor pre %s",actor.Head.String())

	for i := 0; i < len(blks); i++ {

		for j := i + 1; j < len(blks); j++ {
			if blks[i].Miner == blks[j].Miner {
				return cid.Undef, cid.Undef,
					xerrors.Errorf("duplicate miner in a tipset (%s %s)",
						blks[i].Miner, blks[j].Miner)
			}
		}
	}

	var parentEpoch abi.ChainEpoch
	pstate := blks[0].ParentStateRoot
	if blks[0].Height > 0 {
		parent, err := sm.cs.GetBlock(blks[0].Parents[0])
		if err != nil {
			return cid.Undef, cid.Undef, xerrors.Errorf("getting parent block: %w", err)
		}

		parentEpoch = parent.Height
	}

	r := store.NewChainRand(sm.cs, ts.Cids())

	blkmsgs, err := sm.cs.BlockMsgsForTipset(ts)


	if err != nil {
		return cid.Undef, cid.Undef, xerrors.Errorf("getting block messages for tipset: %w", err)
	}

	baseFee := blks[0].ParentBaseFee


	ret1,ret2,err:=sm.ApplyBlocks(ctx, parentEpoch, pstate, blkmsgs, blks[0].Height, r, cb, baseFee, ts)


	gl.FinalizeGL(ctx)

	return ret1,ret2,err
}

func (sm * StateManager) getMinerState(ctx context.Context, maddr addr.Address,ts *types.TipSet) (miner.State,error){

		act, err := sm.LoadActor(ctx, maddr, ts)
		//if err != nil {
		//	return nil, xerrors.Errorf("(get sset) failed to load miner actor: %w", err)
		//}

		mas,err := miner.Load(sm.cs.ActorStore(ctx), act)
		//if err != nil {
		//	return nil, xerrors.Errorf("(get sset) failed to load miner actor state: %w", err)
		//}


		return mas,err

}




func ComputeStateReplay(ctx context.Context, sm *StateManager, height abi.ChainEpoch, msgs []*types.Message, ts *types.TipSet) (cid.Cid, []*api.InvocResult, error) {
	if ts == nil {
		ts = sm.cs.GetHeaviestTipSet()
	}
	var trace []*api.InvocResult
	st, _, err := sm.computeTipSetStateEx(ctx, ts, traceFunc(&trace),height)
	if err != nil {
		return cid.Undef, nil, err
	}

	return st, trace, nil
}
func ComputeStateReplay_del(ctx context.Context, sm *StateManager, height abi.ChainEpoch, msgs []*types.Message, ts *types.TipSet) (cid.Cid, []*api.InvocResult, error) {
	if ts == nil {
		ts = sm.cs.GetHeaviestTipSet()
	}


	base, trace, err := sm.ExecutionTrace(ctx, ts)

	return cid.Cid{}, nil, err



	if err != nil {
		return cid.Undef, nil, err
	}

	for i := ts.Height(); i < height; i++ {
		// handle state forks
		base, err = sm.handleStateForks(ctx, base, i, traceFunc(&trace), ts)
		if err != nil {
			return cid.Undef, nil, xerrors.Errorf("error handling state forks: %w", err)
		}

		// TODO: should we also run cron here?
	}

	r := store.NewChainRand(sm.cs, ts.Cids())
	vmopt := &vm.VMOpts{
		StateBase:      base,
		Epoch:          height,
		Rand:           r,
		Bstore:         sm.cs.StateBlockstore(),
		Syscalls:       sm.cs.VMSys(),
		CircSupplyCalc: sm.GetVMCirculatingSupply,
		NtwkVersion:    sm.GetNtwkVersion,
		BaseFee:        ts.Blocks()[0].ParentBaseFee,
		LookbackState:  LookbackStateGetterForTipset(sm, ts),
	}
	vmi, err := sm.newVM(ctx, vmopt)
	if err != nil {
		return cid.Undef, nil, err
	}

	for i, msg := range msgs {
		// TODO: Use the signed message length for secp messages
		ret, err := vmi.ApplyMessage(ctx, msg)
		if err != nil {
			return cid.Undef, nil, xerrors.Errorf("applying message %s: %w", msg.Cid(), err)
		}
		if ret.ExitCode != 0 {
			log.Infof("compute state apply message %d failed (exit: %d): %s", i, ret.ExitCode, ret.ActorErr)
		}
	}

	root, err := vmi.Flush(ctx)
	if err != nil {
		return cid.Undef, nil, err
	}

	return root, trace, nil
}
