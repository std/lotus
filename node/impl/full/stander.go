package full

import (
	"context"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

//func (a *StateModule) StateExportTipset(ctx context.Context, epoch abi.ChainEpoch,tsk types.TipSetKey) (*api.ComputeStateOutput, error) {
//	ts, err := a.Chain.GetTipSetFromKey(tsk)
//	if err != nil {
//		return nil, xerrors.Errorf("loading tipset %s: %w", tsk, err)
//	}
//
//	if ctx.Amount("export")==nil {
//		panic("key export is nil")
//	}
//
//	ctxV:=context.WithValue(ctx,"replay",true)
//
//
//	//pstate:=ts.Blocks()[0].ParentStateRoot
//
//	st_cid, t, err := stmgr.ComputeStateReplay(ctxV, a.StateManager, epoch, make([]*types.Message,0), ts)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &api.ComputeStateOutput{
//		Root:  st_cid,
//		Trace: t,
//	}, nil
//}


func (a *StateAPI) StateListMessages(ctx context.Context, match *api.MessageMatch, tsk types.TipSetKey, epoch abi.ChainEpoch) ([]cid.Cid, error) {


	ts, err := a.Chain.GetTipSetFromKey(tsk)
	if err != nil {
		return nil, xerrors.Errorf("loading tipset %s: %w", tsk, err)
	}

	ctx1:=context.WithValue(ctx,"replay",true)
	ctxV:=context.WithValue(ctx1,"epoch",epoch)

	//pstate:=ts.Blocks()[0].ParentStateRoot

	st_cid, _, err := stmgr.ComputeStateReplay(ctxV, a.StateManager, epoch, make([]*types.Message,0), ts)

	if err != nil {
		return nil, err
	}

	return []cid.Cid{st_cid}	, nil
}