package z_miner_ledger

import (
	"context"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"go.mongodb.org/mongo-driver/bson"
)

func RecreateCollection(ctx context.Context,colName string){
	con:=ledg_util.GetOrCreateMongoConnection()
	con.RecreateCollection(ctx,colName)

}

const COLLECTION ="entry"
func  FindEntry(cid string  ) ledg.LedgerEntryMongo {
	con:=ledg_util.GetOrCreateMongoConnection()
	l:=&ledg.LedgerEntryMongo{}
	con.FindOne(cid,l)
	return *l
}


func   updateEntry(e *ledg.LedgerEntryMongo) error{
	con:=ledg_util.GetOrCreateMongoConnection()
	filter := bson.M{"_id": e.Id}
	_,err:=con.GetCollection(COLLECTION).UpdateOne(context.TODO(),filter,e)
	//if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
	return err
}

func  InsertEntry(e *ledg.LedgerEntryMongo){
	con:=ledg_util.GetOrCreateMongoConnection()
	con.GetCollection(COLLECTION).InsertOne(context.TODO(), e)

}
