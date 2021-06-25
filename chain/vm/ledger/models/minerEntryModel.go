package models

import (
	"fmt"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
)

type LedgerEntry struct {
	Epoch int32		`gorm:"type:int;primaryKey;autoIncrement:false"`
	Id int32       	`gorm:"type:int;primaryKey;autoIncrement:false"`
	///////////////////////////// types.Message fields {

    AddressId int32  `gorm:"type:int"`
	Address Account `gorm:"foreignKey:AddressId"`
	OffsetId int32  `gorm:"type:int"`
	Offset  Account `gorm:"foreignKey:OffsetId"`

	//Nonce uint64 //Msg.Nonce

	Amount     ledg.FilAmount `gorm:"type:numeric(26)"`
	Balance0    ledg.FilAmount `gorm:"type:numeric(26)"`
	Balance    ledg.FilAmount `gorm:"type:numeric(26)"`
	//GasLimit   int64
	//GasFeeCap  ledg.FilAmount
	//GasPremium ledg.FilAmount
	//TargetActorType models.ActorTypeConst
	MethodId int16 `gorm:"type:int"`

	DimensionId int16
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
	CallDepth 	int16

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

func (e *LedgerEntry) String() string{
	return fmt.Sprintf("Epoch %d Id:%d TxId: %d addrId:%d offsetId:%d MethId:%d MinerId:%d SecId:%d Amount:%s",
		e.Epoch,e.Id,e.TxId,e.AddressId,e.OffsetId,e.MethodId,e.MinerId,e.SectorId,e.Amount)
}



type StorageDealEntry struct{
	ID uint `gorm:"type:bigint;primaryKey;autoIncrement:false"`
}

type RewardEntry struct {
	Id int32 `gorm:"type:int;primaryKey;autoIncrement:false"`
}

