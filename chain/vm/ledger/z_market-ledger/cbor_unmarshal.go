package z_market_ledger

//func (t *DealProposal) UnmarshalCBOR(r io.Reader) error {
//	*t = DealProposal{}
//
//	br := cbg.GetPeeker(r)
//	scratch := make([]byte, 8)
//
//	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//	if maj != cbg.MajArray {
//		return fmt.Errorf("cbor input should be of type array")
//	}
//
//	if extra != 11 {
//		return fmt.Errorf("cbor input had wrong number of fields")
//	}
//
//	// t.PieceCID (cid.Cid) (struct)
//
//	{
//
//		c, err := cbg.ReadCid(br)
//		if err != nil {
//			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
//		}
//
//		t.PieceCID = c
//
//	}
//	// t.PieceSize (abi.PaddedPieceSize) (uint64)
//
//	{
//
//		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
//		if err != nil {
//			return err
//		}
//		if maj != cbg.MajUnsignedInt {
//			return fmt.Errorf("wrong type for uint64 field")
//		}
//		t.PieceSize = abi.PaddedPieceSize(extra)
//
//	}
//	// t.VerifiedDeal (bool) (bool)
//
//	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//	if maj != cbg.MajOther {
//		return fmt.Errorf("booleans must be major type 7")
//	}
//	switch extra {
//	case 20:
//		t.VerifiedDeal = false
//	case 21:
//		t.VerifiedDeal = true
//	default:
//		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
//	}
//	// t.Client (address._Address) (struct)
//
//	{
//
//		if err := t.Client.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.Client: %w", err)
//		}
//
//	}
//	// t.Provider (address._Address) (struct)
//
//	{
//
//		if err := t.Provider.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.Provider: %w", err)
//		}
//
//	}
//	// t.Label (string) (string)
//
//	{
//		sval, err := cbg.ReadStringBuf(br, scratch)
//		if err != nil {
//			return err
//		}
//
//		t.Label = string(sval)
//	}
//	// t.StartEpoch (abi.ChainEpoch) (int64)
//	{
//		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
//		var extraI int64
//		if err != nil {
//			return err
//		}
//		switch maj {
//		case cbg.MajUnsignedInt:
//			extraI = int64(extra)
//			if extraI < 0 {
//				return fmt.Errorf("int64 positive overflow")
//			}
//		case cbg.MajNegativeInt:
//			extraI = int64(extra)
//			if extraI < 0 {
//				return fmt.Errorf("int64 negative oveflow")
//			}
//			extraI = -1 - extraI
//		default:
//			return fmt.Errorf("wrong type for int64 field: %d", maj)
//		}
//
//		t.StartEpoch = abi.ChainEpoch(extraI)
//	}
//	// t.EndEpoch (abi.ChainEpoch) (int64)
//	{
//		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
//		var extraI int64
//		if err != nil {
//			return err
//		}
//		switch maj {
//		case cbg.MajUnsignedInt:
//			extraI = int64(extra)
//			if extraI < 0 {
//				return fmt.Errorf("int64 positive overflow")
//			}
//		case cbg.MajNegativeInt:
//			extraI = int64(extra)
//			if extraI < 0 {
//				return fmt.Errorf("int64 negative oveflow")
//			}
//			extraI = -1 - extraI
//		default:
//			return fmt.Errorf("wrong type for int64 field: %d", maj)
//		}
//
//		t.EndEpoch = abi.ChainEpoch(extraI)
//	}
//	// t.StoragePricePerEpoch (big.Int) (struct)
//
//	{
//
//		if err := t.StoragePricePerEpoch.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.StoragePricePerEpoch: %w", err)
//		}
//
//	}
//	// t.ProviderCollateral (big.Int) (struct)
//
//	{
//
//		if err := t.ProviderCollateral.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.ProviderCollateral: %w", err)
//		}
//
//	}
//	// t.ClientCollateral (big.Int) (struct)
//
//	{
//
//		if err := t.ClientCollateral.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.ClientCollateral: %w", err)
//		}
//
//	}
//	return nil
//}

//func (t *ClientDealProposal) UnmarshalCBOR(r io.Reader) error {
//	*t = ClientDealProposal{}
//
//	br := cbg.GetPeeker(r)
//	scratch := make([]byte, 8)
//
//	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//	if maj != cbg.MajArray {
//		return fmt.Errorf("cbor input should be of type array")
//	}
//
//	if extra != 2 {
//		return fmt.Errorf("cbor input had wrong number of fields")
//	}
//
//	// t.Proposal (market.DealProposal) (struct)
//
//	{
//
//		if err := t.Proposal.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.Proposal: %w", err)
//		}
//
//	}
//	// t.ClientSignature (crypto.Signature) (struct)
//
//	{
//
//		if err := t.ClientSignature.UnmarshalCBOR(br); err != nil {
//			return xerrors.Errorf("unmarshaling t.ClientSignature: %w", err)
//		}
//
//	}
//	return nil
//}
//func (t *PublishStorageDealsReturn) UnmarshalCBOR(r io.Reader) error {
//	*t = PublishStorageDealsReturn{}
//
//	br := cbg.GetPeeker(r)
//	scratch := make([]byte, 8)
//
//	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//	if maj != cbg.MajArray {
//		return fmt.Errorf("cbor input should be of type array")
//	}
//
//	if extra != 1 {
//		return fmt.Errorf("cbor input had wrong number of fields")
//	}
//
//	// t.IDs ([]abi.DealID) (slice)
//
//	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//
//	if extra > cbg.MaxLength {
//		return fmt.Errorf("t.IDs: array too large (%d)", extra)
//	}
//
//	if maj != cbg.MajArray {
//		return fmt.Errorf("expected cbor array")
//	}
//
//	if extra > 0 {
//		t.IDs = make([]abi.DealID, extra)
//	}
//
//	for i := 0; i < int(extra); i++ {
//
//		maj, val, err := cbg.CborReadHeaderBuf(br, scratch)
//		if err != nil {
//			return xerrors.Errorf("failed to read uint64 for t.IDs slice: %w", err)
//		}
//
//		if maj != cbg.MajUnsignedInt {
//			return xerrors.Errorf("value read for array t.IDs was not a uint, instead got %d", maj)
//		}
//
//		t.IDs[i] = abi.DealID(val)
//	}
//
//	return nil
//}
//func (t *PublishStorageDealsParams) UnmarshalCBOR(r io.Reader) error {
//	*t = PublishStorageDealsParams{}
//
//	br := cbg.GetPeeker(r)
//	scratch := make([]byte, 8)
//
//	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//	if maj != cbg.MajArray {
//		return fmt.Errorf("cbor input should be of type array")
//	}
//
//	if extra != 1 {
//		return fmt.Errorf("cbor input had wrong number of fields")
//	}
//
//	// t.Deals ([]market.ClientDealProposal) (slice)
//
//	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
//	if err != nil {
//		return err
//	}
//
//	if extra > cbg.MaxLength {
//		return fmt.Errorf("t.Deals: array too large (%d)", extra)
//	}
//
//	if maj != cbg.MajArray {
//		return fmt.Errorf("expected cbor array")
//	}
//
//	if extra > 0 {
//		t.Deals = make([]ClientDealProposal, extra)
//	}
//
//	for i := 0; i < int(extra); i++ {
//
//		var v ClientDealProposal
//		if err := v.UnmarshalCBOR(br); err != nil {
//			return err
//		}
//
//		t.Deals[i] = v
//	}
//
//	return nil
//}