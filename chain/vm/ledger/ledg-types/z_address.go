package ledg_types

import "github.com/filecoin-project/go-address"

type AddressMongo struct{ address.Address }

func NewAddressFromString(s string) AddressMongo {
	addr,err:=address.NewFromString(s)
	if err!=nil {return AddressMongo{}}else{
		return AddressMongo{addr}
	}

}


