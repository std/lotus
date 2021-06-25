package z_miner_ledger

import (
	"bytes"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
)

func (l *MinerLedger) getSectorFromParams(msg *types.Message) ledg.SectorID{
	minerId ,_:=address.IDFromAddress(msg.To)
	var snum abi.SectorNumber=0
	snum = 0

	if msg.To.String()=="f04"{

	} else if msg.To.String()=="f05"{

	} else	if ledg_util.IsMiner(msg.To) {
		if msg.Method==6 {
			sectorPrecommitInfo,_:= ledg.UnmarshalSectorPreCommitInfo(msg.Params)
			snum=sectorPrecommitInfo.SectorNumber
		} else if msg.Method==7{
			proveCommitInfo,_:= l.getProveCommitParams(msg.Params)
			snum=proveCommitInfo.SectorNumber
		}
	}
	return ledg.SectorID{
		Miner:  abi.ActorID(minerId),
		Number: snum,
	}
}


func (l *MinerLedger) getProveCommitParams(b []byte) (ProveCommitSectorParams,error) {
	ret:=&ProveCommitSectorParams{}
	err:=ret.UnmarshalCBOR(bytes.NewReader(b))
	if err!=nil {return ProveCommitSectorParams{
		SectorNumber: 9999,
		Proof:        nil,
	},err}
	return *ret,err
}



