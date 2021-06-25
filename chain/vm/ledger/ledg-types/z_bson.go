package ledg_types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/ipfs/go-cid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func (m FilAmount) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(m.String())
}
func (m *FilAmount) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	//rv := bson.RawValue{Type: t, Amount: data}
	s:=""
	err:=bson.RawValue{Type: t, Value: data}.Unmarshal(&s)
	if err!=nil {return err}
	val,err:=big.FromString(s)
	if err!=nil {return err}


	m.Int=val.Int
	return nil
}

func (c Cid) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.String())
}

func (m *Cid) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	s:=""
	err:=bson.RawValue{Type: t, Value: data}.Unmarshal(&s)
	if err!=nil {return err}
	ret,_:=cid.Parse(s)
	m.Cid=ret
	return nil
}


func (m DealWeight) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(m.String())
}
func (m *DealWeight) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	//rv := bson.RawValue{Type: t, Amount: data}
	s:=""
	err:=bson.RawValue{Type: t, Value: data}.Unmarshal(&s)
	if err!=nil {return err}
	val,err:=big.FromString(s)
	if err!=nil {return err}

	m.Int=val.Int
	return nil
}

func (c AddressMongo) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(c.String())
}

func (m *AddressMongo) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	s:=""
	err:=bson.RawValue{Type: t, Value: data}.Unmarshal(&s)
	if err!=nil {return err}

	a,err:=address.NewFromString(s)
	m.Address=a
	if err!=nil {return err}

	return nil
}
func (m StoragePower) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(m.String())
}
func (m *StoragePower) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	//rv := bson.RawValue{Type: t, Amount: data}
	s:=""
	err:=bson.RawValue{Type: t, Value: data}.Unmarshal(&s)
	if err!=nil {return err}
	val,err:=big.FromString(s)
	if err!=nil {return err}


	m.Int=val.Int
	return nil
}


func BsonM2TokenAmount(v interface{}) (FilAmount){
	//fmt.Println(str)
	str:=v.(string)
	if str=="<nil>" {return FilAmount{}}
	amo, err := big.FromString(str)
	if err!=nil {return FilAmount(big.NewInt(0))
	}
	return FilAmount(amo)
}