package ledg_types

import (
	"database/sql/driver"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)
type FilAmount abi.TokenAmount

func FilAmountFromInt(i int64) FilAmount{
	return FilAmount{big.NewInt(0).SetInt64(i)}
}

func (v FilAmount) Value()(value driver.Value,err error) {
	//llog.Infof("FilAmount GormValue")
	ret:=v.String()
	if (ret=="<nil>"){return "",nil}

	return v.String(),nil
}
func (v *FilAmount) Scan(value interface{}) error {

	result := value.(string)
	if result=="<nil>" || result=="" {*v=FilAmount{};return nil}

	result1,err:=big.FromString(result)
	*v = FilAmount(result1)
	return err
}



