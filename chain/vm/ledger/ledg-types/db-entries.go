package ledg_types

import (
	"github.com/filecoin-project/go-state-types/abi"

)

type LedgerEntryMongo struct {
	Id       string `bson:"_id" json:"Id,omitempty"`
	///////////////////////////// types.Message fields {

	Version uint64 //Msg.Version

	Address Address `bson:"Address"` //Msg.To
	Offset  Address //Msg.From

	Nonce uint64 //Msg.Nonce

	Value  FilAmount //msg.Amount
	GasLimit   int64
	GasFeeCap  FilAmount
	GasPremium FilAmount
	//TargetActorType models.Protocol
	Method abi.MethodNum //msg.Method
	Params []byte
	////////////////////////////// }

	MsgCid   Cid
	EntryCid Cid

	TotalAmount DimBalance
	Opening     DimBalance
	Amount      DimBalance
	Balance     DimBalance

	GasFee	    FilAmount
	MinerTip	FilAmount
	GasUsed   	int64 `bson:"gasused,omitempty"`
	CallDepth 	uint64

	SectorNumber SectorNumber
	Miner		Address

	//Deals     []abi.DealID
	Invariant FilAmount //should be zero
	Epoch     abi.ChainEpoch
	EntryType EntryTypeConst

	MethodName string
	Note string
	Implicit bool
}
