package miner_ledger

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	cbor "github.com/ipfs/go-ipld-cbor"
)
//type DimBalance map[int]abi.TokenAmount
//type DimBalanceExport map[int]abi.TokenAmount
type ActorBalances map[address.Address]ledg.DimBalance

type DumpActorStateFunc func (act *types.Actor, b []byte) (interface{}, error)

type MinerLedger struct {

	MinerAddress address.Address
	//vm           *vm.VM
	OriginMsg  *types.Message
	IsImplicit bool
	//currentMsg *types.Message
	st         *state.StateTree
	cst        *cbor.BasicIpldStore
	//Entries    map[string]*models.LedgerEntry
	Opening    ActorBalances
	//Opening1      ActorBalances
	Closing      ActorBalances
	BalancesDiff ActorBalances
	dump         DumpActorStateFunc

	epoch abi.ChainEpoch
	//gasOutputs GasOutputs

	//market_ledger.MarketLedger
	//power_ledger.PowerLedger
	//reward_ledger.RewardLedger

}


type EntryTypeConst int
type LedgerDim int


type Message struct {
	Version uint64

	To   address.Address
	From address.Address

	Nonce uint64

	Value abi.TokenAmount

	GasLimit   int64
	GasFeeCap  abi.TokenAmount
	GasPremium abi.TokenAmount

	Method abi.MethodNum
	Params []byte
}










