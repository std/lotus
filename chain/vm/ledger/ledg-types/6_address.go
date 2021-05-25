package ledg_types

import "github.com/filecoin-project/go-address"

type Address struct{ address.Address }

func NewAddressFromString(s string) Address {
	addr,err:=address.NewFromString(s)
	if err!=nil {return Address{}}else{
		return Address{addr}
	}

}


