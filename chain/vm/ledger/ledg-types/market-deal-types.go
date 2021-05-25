package ledg_types

import (
	"fmt"
	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	acrypto "github.com/filecoin-project/go-state-types/crypto"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/xerrors"
	"io"
)


type PublishStorageDealsReturn struct {
	IDs []abi.DealID
}

type PublishStorageDealsParams struct {
	Deals []ClientDealProposal
}

type ClientDealProposal struct {
	Proposal        DealProposal
	ClientSignature acrypto.Signature
}

type DealState struct {
	SectorStartEpoch abi.ChainEpoch // -1 if not yet included in proven sector
	LastUpdatedEpoch abi.ChainEpoch // -1 if deal state never updated
	SlashEpoch       abi.ChainEpoch // -1 if deal never slashed
}

type DealProposal struct {
	Id DealID `bson:"_id"`

	PieceCID     cid.Cid `checked:"true"` // Checked in validateDeal, CommP
	PieceSize    abi.PaddedPieceSize
	VerifiedDeal bool
	Client       addr.Address
	Provider     addr.Address

	// Label is an arbitrary client chosen label to apply to the deal
	Label string

	// Nominal start epoch. Deal payment is linear between StartEpoch and EndEpoch,
	// with total amount StoragePricePerEpoch * (EndEpoch - StartEpoch).
	// Storage deal must appear in a sealed (proven) sector no later than StartEpoch,
	// otherwise it is invalid.
	StartEpoch           abi.ChainEpoch
	EndEpoch             abi.ChainEpoch
	StoragePricePerEpoch abi.TokenAmount

	ProviderCollateral abi.TokenAmount
	ClientCollateral   abi.TokenAmount

}

type DealStatus int


func (t *DealProposal) UnmarshalCBOR(r io.Reader) error {
	*t = DealProposal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 11 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	// t.PieceSize (abi.PaddedPieceSize) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PieceSize = abi.PaddedPieceSize(extra)

	}
	// t.VerifiedDeal (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.VerifiedDeal = false
	case 21:
		t.VerifiedDeal = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.Client (address._Address) (struct)

	{

		if err := t.Client.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Client: %w", err)
		}

	}
	// t.Provider (address._Address) (struct)

	{

		if err := t.Provider.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Provider: %w", err)
		}

	}
	// t.Label (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Label = string(sval)
	}
	// t.StartEpoch (abi.ChainEpoch) (int64)
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

		t.StartEpoch = abi.ChainEpoch(extraI)
	}
	// t.EndEpoch (abi.ChainEpoch) (int64)
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

		t.EndEpoch = abi.ChainEpoch(extraI)
	}
	// t.StoragePricePerEpoch (big.Int) (struct)

	{

		if err := t.StoragePricePerEpoch.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.StoragePricePerEpoch: %w", err)
		}

	}
	// t.ProviderCollateral (big.Int) (struct)

	{

		if err := t.ProviderCollateral.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ProviderCollateral: %w", err)
		}

	}
	// t.ClientCollateral (big.Int) (struct)

	{

		if err := t.ClientCollateral.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientCollateral: %w", err)
		}

	}
	return nil
}

func (t *ClientDealProposal) UnmarshalCBOR(r io.Reader) error {
	*t = ClientDealProposal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Proposal (market.DealProposal) (struct)

	{

		if err := t.Proposal.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Proposal: %w", err)
		}

	}
	// t.ClientSignature (crypto.Signature) (struct)

	{

		if err := t.ClientSignature.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientSignature: %w", err)
		}

	}
	return nil
}
func (t *PublishStorageDealsReturn) UnmarshalCBOR(r io.Reader) error {
	*t = PublishStorageDealsReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.IDs ([]abi.DealID) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.IDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.IDs = make([]abi.DealID, extra)
	}

	for i := 0; i < int(extra); i++ {

		maj, val, err := cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return xerrors.Errorf("failed to read uint64 for t.IDs slice: %w", err)
		}

		if maj != cbg.MajUnsignedInt {
			return xerrors.Errorf("value read for array t.IDs was not a uint, instead got %d", maj)
		}

		t.IDs[i] = abi.DealID(val)
	}

	return nil
}
func (t *PublishStorageDealsParams) UnmarshalCBOR(r io.Reader) error {
	*t = PublishStorageDealsParams{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Deals ([]market.ClientDealProposal) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Deals: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Deals = make([]ClientDealProposal, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v ClientDealProposal
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Deals[i] = v
	}

	return nil
}


func (s StorageDeal) MarshalBSON() ([]byte, error) {
	sm:=bson.M{
		"_id"			:			  s.Id,
		"PieceCID":             	s.PieceCID.String(),
		"SectorId":             s.SectorId.Number,
		"Status":               s.Status,
		"PieceSize":            s.PieceSize,
		"VerifiedDeal":         s.VerifiedDeal,
		"Client":               s.Client.String(),
		"Provider":             s.Provider.String(),
		"Label":                s.Label,
		"StartEpoch":           s.StartEpoch,
		"EndEpoch":             s.EndEpoch,
		"StoragePricePerEpoch": s.StoragePricePerEpoch.String(),
		"ProviderCollateral":   s.ProviderCollateral.String(),
		"ClientCollateral":     s.ClientCollateral.String(),
		"DealProposal":         s.DealProposal,

	}
	return bson.Marshal(sm)
}

func (s DealProposal) MarshalBSON() ([]byte, error) {
	sm:=bson.M{
		"_id"			:			  s.Id,
		"PieceCID":             	s.PieceCID.String(),
		"PieceSize":            s.PieceSize,
		"VerifiedDeal":         s.VerifiedDeal,
		"Client":               s.Client.String(),
		"Provider":             s.Provider.String(),
		"Label":                s.Label,
		"StartEpoch":           s.StartEpoch,
		"EndEpoch":             s.EndEpoch,
		"StoragePricePerEpoch": s.StoragePricePerEpoch.String(),
		"ProviderCollateral":   s.ProviderCollateral.String(),
		"ClientCollateral":     s.ClientCollateral.String(),

	}
	return bson.Marshal(sm)
}