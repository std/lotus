package ledg_types

import (
	"database/sql/driver"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)
func StoragePowerFromInt(i int64) StoragePower{
	return StoragePower{big.NewInt(0).SetInt64(i)}
}

type StoragePower abi.StoragePower

func (v StoragePower) Value()(driver.Value,error){
	return v.String(),nil
}

func (v *StoragePower) Scan(value interface{}) error {

	result := value.(string)
	if result=="<nil>" {*v=StoragePower{};return nil}

	result1,err:=big.FromString(result)
	*v = StoragePower(result1)
	return err
}
