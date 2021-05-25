package ledg_types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

//func (e LedgerEntryMongo) Export() types.LedgerEntryExport{
//	return types.LedgerEntryExport{
//		TxCid:       e.EntryCid.String(),
//		Address:     e.Address.String(),
//		Offset:      e.Offset.String(),
//		MethodNum:   e.Method,
//		Value:       e.Amount.String(),
//		TotalAmount: e.TotalAmount.Export(),
//		Amount:      e.Amount.Export(),
//
//		Opening: 	 e.Opening.Export(),
//
//		Balance:     e.Balance.Export(),
//		GasFee:      e.GasFee.String(),
//		MinerTip:    e.MinerTip.String(),
//		GasUsed:     e.GasUsed,
//		CallDepth:   e.CallDepth,
//		Sector:      strconv.FormatUint(uint64(e.SectorNumber),10),
//		//Deals:        e.Deals,
//	}
//}
func (e *LedgerEntryMongo) String() string{
	return "String:"+e.EntryCid.String()
}



func (e LedgerEntryMongo) MarshalBSON1() ([]byte, error) {

	cid:=e.EntryCid.String()
	d:=bson.M{
		//"_id":          e.Id,
		"EntryCid":       cid,
		"Address":     e.Address.String(),
		"Offset":      e.Offset.String(),
		"Method":   e.Method.String(),
		"Value":       e.Value.String(),
		"TotalAmount": e.TotalAmount,
		"Amount":      e.Amount,
		"Available":     e.Balance,
		"GasFee":      e.GasFee.String(),
		"GasUsed":     strconv.FormatInt(e.GasUsed,10),
		"CallDepth":   strconv.FormatUint(e.CallDepth,10),
		"Sector":      bson.M{"miner":e.Miner.String(),"number":strconv.FormatUint(uint64(e.SectorNumber),10)},
		//"Deals":        e.Deals,
		"Epoch":		   strconv.FormatUint(uint64(e.Epoch),10),
		"Nonce":		   strconv.FormatUint(uint64(e.Nonce),10),
		"MethodName": e.MethodName,

	}


	//return bson.Marshal(Addr{"Marshalled"})
	return bson.Marshal(d)
}

func (e *LedgerEntryMongo) UnmarshalBSON1(data []byte) error {
	var m bson.M
	err := bson.Unmarshal(data, &m);	if err != nil {return err}

	//e.Id=m["_id"].(string)
	//c,_:=ledg.CidFromString(e.Id)
	//e.EntryCid=c
	method,_:=strconv.ParseUint(m["Method"].(string),10,64)
	//fmt.Println(m)
	addr,err:=address.NewFromString(m["Address"].(string));		if err != nil {return err}
	e.Address =Address{addr}
	offset,err:=address.NewFromString(m["Offset"].(string));		if err != nil {return err}
	e.Offset=Address{offset}
	e.Method = abi.MethodNum(method)
	e.Value = BsonM2TokenAmount(m["Amount"])

	e.TotalAmount,_ = DimBalanceFromBsonM(m["TotalAmount"])
	e.Amount,_= DimBalanceFromBsonM(m["Amount"])
	e.Balance,_= DimBalanceFromBsonM(m["Available"])
	e.GasFee= BsonM2TokenAmount(m["GasFee"])
	e.GasUsed,_=strconv.ParseInt(m["GasUsed"].(string),10,64)
	e.CallDepth,_=strconv.ParseUint(m["CallDepth"].(string),10,64)

	//e.Sector,_= SectorFromBsonM(m["Sector"])

	//todo ledger deals bson unmarshal
	//deal,err:=strconv.ParseUint(m["Deals"].(string),10,64);if err != nil {return err}
	//e.Deals =abi.DealID(deal)

	return nil
}




func SectorFromBsonM(data interface{}) (SectorID,error){
	m,ok:=data.(bson.M)
	if ok {
		miner,err:=strconv.ParseUint(m["miner"].(string),10,64);if err!=nil {return SectorID{},err}
		secN,err:=strconv.ParseUint(m["number"].(string),10,64);if err!=nil {return SectorID{},err}
		return SectorID{
			Miner:  abi.ActorID(miner),
			Number: abi.SectorNumber(secN),
		},nil
	} else {return SectorID{},nil}
}