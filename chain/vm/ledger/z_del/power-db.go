package z_del
//package power_ledger
//
//import (
//	"context"
//	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
//	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
//	"go.mongodb.org/mongo-driver/bson"
//)
//
//func  FindEntry( cid string  ) models.PowerEntry {
//	con:=ledg_util.GetOrCreateMongoConnection()
//	l:=&models.PowerEntry{}
//	con.FindOne(cid,l)
//	return *l
//}
//
////func  (l *PowerLedger)InsertEntry(e PowerEntry){
////	con:=ledg_util.GetOrCreateMongoConnection()
////	_,err:=con.GetCollection("entry").InsertOne(context.TODO(), e)
////	if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
////}
//
//
///* SECTORS */
//func  (l *PowerLedger) GetSector( id string  ) models.Sector {
//	con:=ledg_util.GetOrCreateMongoConnection()
//	s:=&models.Sector{}
//	con.FindOne(id,s)
//	return *s
//}
//
//
//func  (l *PowerLedger) insertSector1(s *models.Sector) error{
//
//	db:=ledg_util.GetPgDatabase()
//
//	var found models.Sector
//	result:=db.First(&found,s.ID)
//	if result.RowsAffected>0{ return nil}
//
//	db.Create(s)
//	return nil
//}
//
//
//
//
//func  (l *PowerLedger) updateSector(s models.Sector,se models.PowerEntry) error{
//	con:=ledg_util.GetOrCreateMongoConnection()
//
//	filter := bson.M{"_id": s.ID}
//	update:=	bson.M{
//		"$set": bson.M{
//			"InitialPledge":"updated!!!",
//		},
//		//"$push": bson.M{"Entries":bson.M{"entry":"create sector","sector":s.SectorNum}},
//		"$push": bson.M{"Entries":se},
//	}
//
//	_,err:=con.GetCollection("sectors").UpdateOne(context.TODO(), filter,update)
//	//if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
//	return err
//}
//
//
///* ACTOR and Miners */
//
//
//func  (l *PowerLedger) insertMinerMongo(m *models.Account) error {
//	db:=ledg_util.GetPgDatabase()
//
//	var found models.Account
//	result:=db.First(&found,m.ID)
//	if result.RowsAffected>0{ return nil }
//
//	db.Create(m)
//	//if db.Create(m).Error!=nil{return db.Save(m).Error}
//	return nil
//
//}
//
//func  (l *PowerLedger) insertSectorMongo(s models.Sector) error{
//
//	con:=ledg_util.GetOrCreateMongoConnection()
//	_,err:=con.GetCollection("sectors").InsertOne(context.TODO(), s)
//	if err!=nil {
//		ledg_util.Log(bson.M{
//			"source":"PowerLedger.insertSector",
//			//"Miner":s.Miner.String(),
//			//"SectorNum":s.SectorNum,
//			"Epoch":s.PreCommitEpoch,
//			"msgCid":" not implemented ",
//			"err":err.Error(),
//
//		})
//	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
//	return err
//}
//
//func  (l *PowerLedger)insertSectorEntryMongo(s *models.Sector,se *models.PowerEntry) error{
//	con:=ledg_util.GetOrCreateMongoConnection()
//
//	filter := bson.M{"_id": s.ID}
//	update:=	bson.M{
//		"$push": bson.M{"entries":se},
//	}
//	_,err:=con.GetCollection("sectors").UpdateOne(context.TODO(), filter,update)
//	//if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
//	return err
//}
//
//
//
