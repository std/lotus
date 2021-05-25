package ledg_types

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	cbg "github.com/whyrusleeping/cbor-gen"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/xerrors"
	"io"
)





type SectorPreCommitInfo miner.SectorPreCommitInfo

type SectorPreCommitOnChainInfo miner.SectorPreCommitOnChainInfo

type SectorOnChainInfo miner.SectorOnChainInfo

func (t *SectorPreCommitInfo) UnmarshalCBOR(r io.Reader) error {
	*t = SectorPreCommitInfo{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 10 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.SealProof (abi.RegisteredSealProof) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SealProof = abi.RegisteredSealProof(extraI)
	}
	// t.SectorNumber (abi.SectorNumber) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.SectorNumber = abi.SectorNumber(extra)

	}
	// t.SealedCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SealedCID: %w", err)
		}

		t.SealedCID = c

	}
	// t.SealRandEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SealRandEpoch = abi.ChainEpoch(extraI)
	}
	// t.DealIDs ([]abi.DealID) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.DealIDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.DealIDs = make([]abi.DealID, extra)
	}

	for i := 0; i < int(extra); i++ {

		maj, val, err := cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return xerrors.Errorf("failed to read uint64 for t.DealIDs slice: %w", err)
		}

		if maj != cbg.MajUnsignedInt {
			return xerrors.Errorf("value read for array t.DealIDs was not a uint, instead got %d", maj)
		}

		t.DealIDs[i] = abi.DealID(val)
	}

	// t.Expiration (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Expiration = abi.ChainEpoch(extraI)
	}
	// t.ReplaceCapacity (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.ReplaceCapacity = false
	case 21:
		t.ReplaceCapacity = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.ReplaceSectorDeadline (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ReplaceSectorDeadline = uint64(extra)

	}
	// t.ReplaceSectorPartition (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ReplaceSectorPartition = uint64(extra)

	}
	// t.ReplaceSectorNumber (abi.SectorNumber) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ReplaceSectorNumber = abi.SectorNumber(extra)

	}
	return nil
}


//func (s Sector) MarshalBSON() ([]byte, error) {
//
//	sm:=bson.M{
//		//"_id"			:			  s.id,
//		"SectorNum"	:               s.SectorNum,
//		"Miner"		:				s.Miner.String(),
//
//		"PreCommitEpoch": s.PreCommitEpoch,
//		"PreCommitDeposit": s.PreCommitDeposit,
//
//		"SectorPreCommitInfo":s.SectorPreCommitInfo,
//
//	}
//	return bson.Marshal(sm)
//}


func (s SectorPreCommitInfo) MarshalBSON() ([]byte, error) {
	sm:=bson.M{
		//"_id"			:			  s.id,
		"SectorNumber":               s.SectorNumber,
		"SealProof":                  s.SealProof,
		"SealedCID":                  s.SealedCID.String(),
		"DealIDs":                    s.DealIDs,
		"Expiration":                 s.Expiration,

		"ReplaceCapacity":        s.ReplaceCapacity,
		"ReplaceSectorDeadline":  s.ReplaceSectorDeadline,
		"ReplaceSectorPartition": s.ReplaceSectorPartition,
		"ReplaceSectorNumber":    s.ReplaceSectorNumber,

	}


return bson.Marshal(sm)
}

func (s SectorPreCommitOnChainInfo) MarshalBSON() ([]byte, error) {

	sm:=bson.M{
		//"_id"			:			  s.id,
		"Info":s.Info,
		"PreCommitDeposit":   s.PreCommitDeposit,
		"PreCommitEpoch":     s.PreCommitEpoch,
		"DealWeight":         s.DealWeight,
		"VerifiedDealWeight": s.VerifiedDealWeight,
	}

	return bson.Marshal(sm)
}

func (s SectorOnChainInfo) MarshalBSON() ([]byte, error) {

	sm:=bson.M{
		//"_id"			:			  s.id,
		"SectorNumber":          s.SectorNumber,
		"SealProof":             s.SealProof,
		"SealedCID":             s.SealedCID,
		"DealIDs":               s.DealIDs,
		"Activation":            s.Activation,
		"Expiration":            s.Expiration,
		"DealWeight":            s.DealWeight,
		"VerifiedDealWeight":    s.VerifiedDealWeight,
		"InitialPledge":         s.InitialPledge,
		"ExpectedDayReward":     s.ExpectedDayReward,
		"ExpectedStoragePledge": s.ExpectedStoragePledge,
	}

	return bson.Marshal(sm)
}