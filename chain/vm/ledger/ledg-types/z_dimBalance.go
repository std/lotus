package ledg_types

import (
	"github.com/filecoin-project/go-state-types/big"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type DimBalance map[int16]FilAmount

func  Sub(op1_, op2_ FilAmount) FilAmount{

	if op1_.Int== nil {op1_=FilAmount(big.NewInt(0))}
	if op2_.Int== nil {op2_=FilAmount(big.NewInt(0))}

	ret:=big.NewInt(0).Sub(op1_.Int, op2_.Int)

	return FilAmount{ret}
}
const  (
	Available int16 = iota
	PreCommitDeposits
	InitialPledge
	LockedFunds
	Vesting
	FeeDebt
)

func  Sub_orig(op1, op2 big.Int) big.Int{
	if op1.Nil() && op2.Nil() { return big.Int{} }
	if op1.Nil(){ op1 = big.Zero() }
	if op2.Nil(){ op2 = big.Zero() }
	return big.Sub(op1,op2)
}

func (d DimBalance) MarshalBSON() ([]byte, error) {
	//fmt.Printf("Marshal dim balance1 %v",d)
	m:=bson.M{}
	for i,v:=range d {m[strconv.Itoa(int(i))]=v.String()}
	return bson.Marshal(m)
}


func (closing DimBalance) Diff (opening DimBalance) DimBalance {

	return DimBalance{
		PreCommitDeposits: Sub(closing[PreCommitDeposits], opening[PreCommitDeposits]),
		LockedFunds:       Sub(closing[LockedFunds], opening[LockedFunds]),
		InitialPledge:     Sub(closing[InitialPledge], opening[InitialPledge]),
		Vesting:           Sub(closing[Vesting], opening[Vesting]),
		Available:         Sub(closing[Available], opening[Available]),
	}
}


//func nonNilDimBalance(d DimBalance) DimBalance{
//	if d==nil {return DimBalance{}}else {return d }
//}


//func (dimBalance DimBalance) Export() types.DimBalance{
//	ret:=make(types.DimBalance,0)
//
//	for k,v:=range dimBalance { ret[k]=v.String() }
//	return  ret
//}
func  DimBalanceFromBsonM(data interface{}) (DimBalance,error) {
	m,ok:=data.(bson.M)

	//fmt.Println(m)
	//err := bson.Unmarshal(data, &m);	if err != nil {return err}
	if ok {
		d := DimBalance{}
		//fmt.Println("Unmarshal DimBalance")
		for k, v := range m {
			i, err := strconv.Atoi(k);
			if err != nil {
				return nil, err
			}
			str, ok := v.(string)
			if ok {
				amo := BsonM2TokenAmount(str)
				if err != nil {
					return nil, err
				}
				d[int16(i)] = FilAmount(amo)
			} else {
				d[int16(i)] = FilAmount{}
			}
		}
		return d, nil
	}else {
		return DimBalance{}, nil
	}
}


