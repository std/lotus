package market_ledger

import (
	"context"
	"github.com/filecoin-project/go-state-types/abi"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
)

const COLLECTION_NAME="deals"

func  insertDeal(deal ledg.StorageDeal) (error){
	con:=ledg_util.GetOrCreateMongoConnection()
	_,err:=con.GetCollection(COLLECTION_NAME).InsertOne(context.TODO(), deal)
	//if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
	return err
}
func  FindDeal(dealId abi.DealID) (*ledg.StorageDeal,error){
	//con:=ledg_util.GetOrCreateMongoConnection()
	//_,err:=con.getCollection("deals").FindOne()
	//if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
	return nil,nil
}


