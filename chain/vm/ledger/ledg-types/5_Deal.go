package ledg_types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type DealID abi.DealID
type DealWeight abi.DealWeight


//from DealProposal with DealId
//type DealStatus int
type MarketEntryType int
//from mongo
const (
	DealOffered = iota//?
	DealDataTransferred//?
	DealPublished
	DealSectorPreCommited
	DealSectorCommited
	DealSectorTerminated
	DealSectorExpired
)
//from mongo
type Deal struct {
	abi.DealID
	//DealId abi.DealID
	SectorId abi.SectorID
	Status   DealStatus

	PieceCID             cid.Cid
	PieceSize            abi.PaddedPieceSize
	VerifiedDeal         bool
	Client               address.Address
	Provider             address.Address
	Label                string
	StartEpoch           abi.ChainEpoch
	EndEpoch             abi.ChainEpoch
	StoragePricePerEpoch abi.TokenAmount
	ProviderCollateral   abi.TokenAmount
	ClientCollateral     abi.TokenAmount

	//market.DealState
	//type DealState struct {
	SectorStartEpoch abi.ChainEpoch // -1 if not yet included in proven sector
	LastUpdatedEpoch abi.ChainEpoch // -1 if deal state never updated
	SlashEpoch       abi.ChainEpoch // -1 if deal never slashed
}