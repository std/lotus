package models

import (
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
)

type VestingEntry struct{
	MinerId int32  `gorm:"type:int;primaryKey;autoIncrement:false"`
	Epoch int32		`gorm:"type:int;primaryKey;autoIncrement:false"`
	EntryType int16
	Modified int32
	Amount ledg.FilAmount  `gorm:"type:numeric(26)"`
	Amount1 ledg.FilAmount  `gorm:"type:numeric(26)"`
	Amount2 ledg.FilAmount  `gorm:"type:numeric(26)"`
	AmountStr ledg.FilAmount  `gorm:"type:numeric(32,26)"`
	//Id int32       	`gorm:"type:int;primaryKey;autoI ncrement:false"`
	/////////////////////////////// types.Message fields {
	//
	//AddressId int32  `gorm:"type:int"`
	//Address Account `gorm:"foreignKey:AddressId"`
	//OffsetId int32  `gorm:"type:int"`
	//Offset  Account `gorm:"foreignKey:OffsetId"`
}

type VestingSchedule struct{
	MinerId int32  `gorm:"type:int;primaryKey;autoIncrement:false"`
	Epoch int32		`gorm:"type:int;primaryKey;autoIncrement:false"`
	EntryType int16
	Modified int32
	//Id int32       	`gorm:"type:int;primaryKey;autoIncrement:false"`
	/////////////////////////////// types.Message fields {
	//
	//AddressId int32  `gorm:"type:int"`
	//Address Account `gorm:"foreignKey:AddressId"`
	//OffsetId int32  `gorm:"type:int"`
	//Offset  Account `gorm:"foreignKey:OffsetId"`
}

func (VestingSchedule) TableName() string {
	return "vesting_stage"
}
