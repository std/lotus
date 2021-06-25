package Posting

import (
	"github.com/filecoin-project/go-state-types/big"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	m "github.com/filecoin-project/lotus/chain/vm/ledger/models"
)


func  (l *LedgerPosting) NewSectorOnChain_del(s *m.Sector, origMsgCid, sectorPreCommitInfo ledg.Cid) error {

	ledg_util.GetOrCreateAccountFromId(s.MinerId,"",l.Epoch)

	l.insert(s,false)

	se:=&m.PowerEntry{
		MinerId:         s.MinerId,
		SectorId: s.ID,
		//SectorNumber:  s.SectorNum,
		//MsgCid:        origMsgCid,
		EntryType:     ledg.SectorEntryType.PreCommit,
		LockedBalance: s.PreCommitDeposit,
		//DealCount: 	   len(s.DealIDs),
		RawBytePower:   ledg.StoragePower(big.NewInt(112233)),
		QualityAdjPower: ledg.StoragePower(big.NewInt(1222)),
		//StateCid:      sectorPreCommitInfo,
	}

	l.insert(se,false)
	return nil

}

func (l *LedgerPosting) PowerConstructor(p ledg_util.ActorMethodParams) {

	
}

func (l *LedgerPosting) UpdateClaimedPower(p ledg_util.ActorMethodParams) {
	
}

func (l *LedgerPosting) EnrollCronEvent(p ledg_util.ActorMethodParams) {
	
}

func (l *LedgerPosting) OnEpochTickEnd(p ledg_util.ActorMethodParams) {
	
}

func (l *LedgerPosting) UpdatePledgeTotal(p ledg_util.ActorMethodParams) {
	
}

func (l *LedgerPosting) Deprecated1(p ledg_util.ActorMethodParams) {

}

func (l *LedgerPosting) SubmitPoRepForBulkVerify(p ledg_util.ActorMethodParams) {
	
}

func (l *LedgerPosting) CurrentTotalPower(p ledg_util.ActorMethodParams) {
	//readonly
}


