package models

import (
	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	"time"
)

type ActorTypeConst uint



const (
	System   ActorTypeConst =iota //f00
	Init                          //f01
	Reward                        //f02
	Cron                          //f03
	Power                         //f04
	Market                        //f05
	Registry                      //f06
	Burn     = 99
	Miner    = 1000

	//Miner //Reward
	//Send
)

type TablePart struct {
	Id int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	TablePartName string
}
type AccountType struct {
	Id int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	Descr string
}

type Tipset struct {
	Id int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	Time time.Time
	Block0 ledg.Cid
	Block1 ledg.Cid
	Block2 ledg.Cid
	Block3 ledg.Cid
	Block4 ledg.Cid
	Block5 ledg.Cid
	Block6 ledg.Cid
	Block7 ledg.Cid
}
type Block struct {
	Id       int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	Time     time.Time
	TipsetId int32
	Tipset   Tipset
	MinerId  int32
	Miner    Account
	Cid      ledg.Cid
}
type TxMessage struct {

	Id         int `gorm:"primaryKey;autoIncrement:false"`
	FromId     int32
	ToId       int32
	From       Account
	To         Account
	Value      ledg.FilAmount
	GasFeeCap  ledg.FilAmount
	GasPremium ledg.FilAmount
	GasLimit   int64
	Method     abi.MethodNum
	Params     []byte
	MethodName string
	Implicit bool

}

type Account struct {
	//gorm.Model
	ID int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	Address string
	RobustAddress string
	Name string
	CompanyId string
	Signature string

	CreationEpoch abi.ChainEpoch
	//Balance DimBalance `bson:"DimBalance"`
	//PowerBalance `bson:"PowerBalance"`
	//SectorCounts `bson:"SectorCounts"`
	//
	Protocol addr.Protocol
	//MessagesCount uint64
	//TotalReward FilAmount `bson:"TotalReward"`
	//Stats MiningStats `bson:"MiningStats"`
	//Properties MinerProperties

}



type Sector struct {
	MinerId int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	ID int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
	//SectorNum ledg.SectorNumber `bson:"SectorNum"`

	Miner Account

	PreCommitEpoch 	abi.ChainEpoch  `bson:"PreCommitEpoch"`
	CommitEpoch		abi.ChainEpoch  `bson:"CommitEpoch"`
	ActivationEpoch		abi.ChainEpoch  `bson:"ActivationEpoch"`
	ExpirationEpoch abi.ChainEpoch  `bson:"ExpirationEpoch"`

	//Deals  []ledg.DealID `bson:"Deals"`
	PreCommitDeposit ledg.FilAmount `sql:"type:varchar"`// `bson:"PreCommitDeposit"`//???
	InitialPledge ledg.FilAmount `sql:"type:varchar"`//`bson:"InitialPledge"`//???

	SealProof abi.RegisteredSealProof `bson:"SealProof"`//???

	Status int8 `bson:"Status"`//???
	// Sum of raw byte power for a miner's sectors.
	RawBytePower ledg.StoragePower `sql:"type:varchar"`//`bson:"StoragePower"`
	// Sum of quality adjusted power for a miner's sectors.
	QualityAdjPower ledg.StoragePower `sql:"type:varchar"`//`bson:"QualityAdjPower"`

	PreCommitFee ledg.FilAmount `sql:"type:varchar"`//`bson:"PreCommitFee"`
	CommitFee ledg.FilAmount `sql:"type:varchar"`//`bson:"CommitFee"`
	//	Winning []int//??

	//SectorPreCommitInfo ledg.SectorPreCommitInfo  `sql:"type:varchar"`//`bson:"_SectorPreCommitInfo"`
}

type StorageDeal struct {
	Id abi.DealID `bson:"_id"`

	PieceCID     ledg.Cid `bson:"PieceCID" checked:"true"` // Checked in validateDeal, CommP

	SectorId abi.SectorID `bson:"SectorId"`
	Status   ledg.DealStatus
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

	//DealProposal DealProposal  `bson:"DealProposal"`
}
