package ledg_types

import (
	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	//"github.com/filecoin-project/lotus/chain/vm/ledger/models"

	"github.com/ipfs/go-cid"
)

type StorageDeal struct {
	Id abi.DealID `bson:"_id"`

	PieceCID     cid.Cid `bson:"PieceCID" checked:"true"` // Checked in validateDeal, CommP

	SectorId abi.SectorID `bson:"SectorId"`
	Status   DealStatus
	////////////////////

	PieceSize    abi.PaddedPieceSize `bson:"PieceSize"`
	VerifiedDeal bool `bson:"VerifiedDeal"`
	Client       addr.Address `bson:"Client"`
	Provider     addr.Address `bson:"Provider"`

	// Label is an arbitrary client chosen label to apply to the deal
	Label string `bson:"Label"`

	// Nominal start epoch. Deal payment is linear between StartEpoch and EndEpoch,
	// with total amount StoragePricePerEpoch * (EndEpoch - StartEpoch).
	// Storage deal must appear in a sealed (proven) sector no later than StartEpoch,
	// otherwise it is invalid.
	StartEpoch           abi.ChainEpoch `bson:"StartEpoch"`
	EndEpoch             abi.ChainEpoch `bson:"EndEpoch"`
	StoragePricePerEpoch abi.TokenAmount `bson:"StoragePricePerEpoch"`

	ProviderCollateral abi.TokenAmount `bson:"ProviderCollateral"`
	ClientCollateral   abi.TokenAmount `bson:"ClientCollateral"`

	//type DealState struct {
	//	SectorStartEpoch abi.ChainEpoch // -1 if not yet included in proven sector
	//	LastUpdatedEpoch abi.ChainEpoch // -1 if deal state never updated
	//	SlashEpoch       abi.ChainEpoch // -1 if deal never slashed
	//}

	DealProposal DealProposal  `bson:"DealProposal"`
}

type SectorMongo struct {
	Id        string `bson:"_id"`
	SectorNum SectorNumber `bson:"SectorNum"`

	Miner AddressMongo `bson:"Miner"`

	PreCommitEpoch 	abi.ChainEpoch  `bson:"PreCommitEpoch"`
	CommitEpoch		abi.ChainEpoch  `bson:"CommitEpoch"`
	ActivationEpoch		abi.ChainEpoch  `bson:"ActivationEpoch"`
	ExpirationEpoch abi.ChainEpoch  `bson:"ExpirationEpoch"`

	Deals  []DealID `bson:"Deals"`
	PreCommitDeposit FilAmount`bson:"PreCommitDeposit"`//???
	InitialPledge FilAmount `bson:"InitialPledge"`//???

	SealProof abi.RegisteredSealProof `bson:"SealProof"`//???

	Status uint `bson:"Status"`//???
	// Sum of raw byte power for a miner's sectors.
	RawBytePower StoragePower `bson:"StoragePower"`
	// Sum of quality adjusted power for a miner's sectors.
	QualityAdjPower StoragePower `bson:"QualityAdjPower"`

	PreCommitFee FilAmount `bson:"PreCommitFee"`
	CommitFee FilAmount `bson:"CommitFee"`
	//	Winning []int//??

	SectorPreCommitInfo SectorPreCommitInfo  `bson:"_SectorPreCommitInfo"`
}




type MiningStats struct{
	PowerGrowth StoragePower `bson:"PowerGrowth"`
	BlocksMined uint `bson:"BlocksMined"`
	MiningEfficiencyPerTB FilAmount `bson:"MiningEfficiencyPerTB"`
	WinCount uint `bson:"WinCount"`
	MinerEquivalent float32 `bson:"MinerEquivalent"`
}

type MinerProperties struct {
	PeerId  string       `bson:"PeerId"`
	Owner   AddressMongo `bson:"Owner"`
	Worker  AddressMongo `bson:"Worker"`
	Region  string       `bson:"Region"`
	Country string       `bson:"Country"`
	Ip      string       `bson:"Ip"`
}



type ActorMongo struct {
	Id string `bson:"_id"`
	Address string `bson:"AddressMongo"`
	Name string `bson:"Name"`
	CompanyId string `bson:"CompanyId"`
	Signature string `bson:"Signature"`

	CreationEpoch abi.ChainEpoch `bson:"CreationEpoch"`
	//Balance DimBalance `bson:"DimBalance"`
	//PowerBalance `bson:"PowerBalance"`
	//SectorCounts `bson:"SectorCounts"`
	//
	//Protocol models.Protocol `bson:"Protocol"`
	MessagesCount uint64 `bson:"MessageCount"`
	//TotalReward FilAmount `bson:"TotalReward"`
	//Stats MiningStats `bson:"MiningStats"`
	//Properties MinerProperties

}