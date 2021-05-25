package ledg_types

import (
	"github.com/filecoin-project/go-state-types/abi"
	"strconv"
)

type SectorNumber abi.SectorNumber

type SectorID abi.SectorID

func (id SectorID) String() string {
	return id.Miner.String()+"-"+strconv.FormatUint(uint64(id.Number),10)
}





