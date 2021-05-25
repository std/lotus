package ledg_types

import "github.com/filecoin-project/go-state-types/abi"

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount

	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount
	Refund       abi.TokenAmount

	GasRefund int64
	GasBurned int64
	GasUsed   int64
}


















