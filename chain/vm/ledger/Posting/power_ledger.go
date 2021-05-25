package Posting

import (
	"github.com/filecoin-project/go-state-types/big"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	m "github.com/filecoin-project/lotus/chain/vm/ledger/models"
)

//type PowerLedger struct {
//	MinerId uint32
//	Miner models.Account
//
//	SectorId uint32
//	Sector models.Sector
//
//	RawBytePower ledg.StoragePower
//	QualityAdjPower ledg.StoragePower
//
//	epoch abi.ChainEpoch
//	//db *gorm.DB
//}
//func  CreateGeneralLedger(ctx context.Context, epoch abi.ChainEpoch) *PowerLedger {
//
//	return &PowerLedger{
//
//		epoch: epoch,
//		db:ledg_util.GetPgDatabase(),
//	}
//}

//type PowerTotal struct {
//	RawBytePower ledg.StoragePower
//	// Sum of quality adjusted power for a miner's sectors.
//	QualityAdjPower ledg.StoragePower
//}
//
//type LedgerPosting struct {
//	MinerId uint32
//	Miner models.Account
//
//	SectorId uint32
//	Sector models.Sector
//
//	RawBytePower ledg.StoragePower
//	QualityAdjPower ledg.StoragePower
//
//	epoch abi.ChainEpoch
//	db *gorm.DB
//}

//type DelGeneralLedger struct {
//	details map[ledg.Address]string
//	epoch abi.ChainEpoch
//	db *gorm.DB
//}


//type GeneralLedgerDetails struct {
//	details map[ledg.Address]map[ledg.SectorNumber] string
//}




func  (l *LedgerPosting) NewSectorOnChain_del(s *m.Sector, origMsgCid, sectorPreCommitInfo ledg.Cid) error {

	ledg_util.GetOrCreateAccountFromId(s.MinerId,"",l.Epoch)

	//if ledg_util.Exists(&models.Sector{},int(s.ID)) { return nil}

	l.insert(s,true)

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


	l.insert(se,true)
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
	
}


