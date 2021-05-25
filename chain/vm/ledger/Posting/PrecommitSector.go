package Posting

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	m "github.com/filecoin-project/lotus/chain/vm/ledger/models"
)

func(l *LedgerPosting) PreCommitSector(p ledg_util.ActorMethodParams) {

	if l.CurrentTxId ==0 { panic("PrecomitSector pre=>CurrentTxId is 0")}

	msg:=p.Msg
	params,_:=ledg.UnmarshalSectorPreCommitInfo(msg.Params)
	minerAddress:=msg.To
	minerId,_:=address.IDFromAddress(minerAddress)

	//_sectorPreCommitInfoCid,_:=abi.CidBuilder.Sum(Msg.Params)
	//sectorPreCommitInfoCid:=ledg.Cid{_sectorPreCommitInfoCid}
	//origingMsgCid:=ledg.Cid{Msg.Cid()}


	deals:=make([]ledg.DealID,len(params.DealIDs))
	//if len(params.DealIDs)>0 {
	//	ledg_util.Log(bson.M{"EpochWithDeals":l.Epoch})
	//}

	for i,d:=range params.DealIDs { deals[i]=ledg.DealID(d)	}

	s:=&m.Sector{
		ID:                  int32(params.SectorNumber),
		//SectorNum:           ledg.SectorNumber(params.SectorNumber),
		MinerId:              int32(minerId),
		SealProof:           params.SealProof, //registeredSealProof
		//Deals:               deals,

		Status: 			0,

		PreCommitEpoch:      l.Epoch,
		CommitEpoch:         0,
		ActivationEpoch:     0,
		ExpirationEpoch:     params.Expiration,

		PreCommitDeposit:    ledg.FilAmount(big.NewInt(1)),
		InitialPledge:       ledg.FilAmount(big.NewInt(0)),

		RawBytePower:        ledg.StoragePower(big.NewInt(32)),
		QualityAdjPower:     ledg.StoragePower(big.NewInt(32)),

		PreCommitFee: ledg.FilAmount(big.NewInt(0)),
		CommitFee: ledg.FilAmount(big.NewInt(0)),
		//SectorPreCommitInfo: params,
	}

	l.insert(s,true)

	se:=&m.PowerEntry{
		MinerId:         s.MinerId,
		SectorId: s.ID,
		TxId: l.CurrentTxId,
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

	e:=l.minerEntryTemplate(p,int32(len(l.minerEntries)),false)

	e.SectorId=&s.ID
	e.MinerId=s.MinerId
	e.TxId=l.CurrentTxId
	l.insert(&e,true)

	e2:=l.minerEntryTemplate(p,int32(len(l.minerEntries)),true)

	e2.SectorId=&s.ID
	e2.MinerId=s.MinerId
	e2.TxId=l.CurrentTxId
	l.insert(&e2,true)
}
