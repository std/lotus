package models

import (
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
)



type PowerEntry struct {
	ID      int     `gorm:"primaryKey"`
	MinerId int32  `gorm:"type:int;foreignKey"`
	Miner   Account `bson:"Miner"`
	TxId int `gorm:"type:int;foreignKey"`
	Tx TxMessage


	SectorId int32  `gorm:"type:int"`
	Sector 	Sector `gorm:"foreignKey:MinerId,SectorId"`
	//MsgCid			ledg.Cid `bson:"MsgCid"`
	EntryType    	ledg.EntryTypeConst

	LockedBalance	ledg.FilAmount
	DealCount		int
	RawBytePower     ledg.StoragePower
	QualityAdjPower  ledg.StoragePower
	//StateCid        ledg.Cid `bson:"StateCid"`
}
