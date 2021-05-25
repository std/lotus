package models

import (
	"github.com/filecoin-project/go-state-types/abi"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
)

type LedgerEntry struct {
	Epoch int32
	Id int32       `gorm:"type:int;primaryKey;autoIncrement:false"`
	///////////////////////////// types.Message fields {
    //Version int8 //Msg.Version

    AddressId int32  `gorm:"type:int"`
	Address Account `gorm:"foreignKey:AddressId"`
	OffsetId int32  `gorm:"type:int"`
	Offset  Account `gorm:"foreignKey:OffsetId"`

	//Nonce uint64 //Msg.Nonce

	Amount     ledg.FilAmount //msg.Amount
	//GasLimit   int64
	//GasFeeCap  ledg.FilAmount
	//GasPremium ledg.FilAmount
	//TargetActorType models.ActorTypeConst
	Method abi.MethodNum //msg.Method

	//Params []byte
	////////////////////////////// }

	//MsgCid   ledg.Cid
	//EntryCid ledg.Cid

	//TotalAmount DimBalance
	//Opening     DimBalance
	//Amount      DimBalance
	//Balance     DimBalance

	//GasFee	    ledg.FilAmount
	//MinerTip	ledg.FilAmount
	//GasUsed   	int64 `bson:"gasused,omitempty"`
	CallDepth 	uint64

	SectorId *int32 `gorm:"type:int"`
	Sector *Sector `gorm:"foreignKey:MinerId,SectorId"`
	MinerId int32 `gorm:"type:int"`
	Miner  Account `gorm:"foreignKey:MinerId"`

	TxId int
	Tx TxMessage `gorm:"foreignKey:TxId"`

	//Deals     []abi.DealID
	//Invariant FilAmount //should be zero
	//Epoch     abi.ChainEpoch
	EntryType ledg.EntryTypeConst

	MethodName string
	//Note string
	Implicit bool
}




type StorageDealEntry struct{
	ID uint `gorm:"type:bigint;primaryKey;autoIncrement:false"`
}

type RewardEntry struct {
	Id int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
}

