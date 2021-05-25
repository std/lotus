package ledg_util

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"reflect"
)


type MongoConnection struct {
	client *mongo.Client
}


func (con *MongoConnection) ping (c *mongo.Client){

	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
}

func (con *MongoConnection) RecreateCollection(ctx context.Context,colName string){
	con.client.Database("ledger").Collection(colName).Drop(ctx)
	con.client.Database("ledger").CreateCollection(ctx,"entry")

}

func (con *MongoConnection) GetCollection(colName string) *mongo.Collection{
	return con.client.Database("ledger").Collection(colName)
}
func (con *MongoConnection) FindOne(id string,result interface{}) {
	//var result ObjHero
	//fmt.Println(reflect.TypeOf(result).String())
	if reflect.TypeOf(result).String()=="*power_ledger.Sector" {
		filter := bson.M{"_id": id}
		col:=con.GetCollection("sectors")
		err := col.FindOne(context.TODO(), filter).Decode(result)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(result)
			fmt.Println("Found a single document: ", result)
		}
	}
}




func (con *MongoConnection) InsertLogEntry(m bson.M) (error){
	_,err:=con.GetCollection("log").InsertOne(context.TODO(), m)
	//if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
	return err
}

func (con *MongoConnection) removeEntry(cid string) {
	filter := bson.M{"_id": cid}
	_, err:=con.GetCollection("entry").DeleteOne(context.TODO(),filter)
	if err!=nil {		log.Fatal(err)	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
}
var mgo *MongoConnection



func GetOrCreateMongoConnection() *MongoConnection {

	if mgo ==nil {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		//ping(c)
		mgo =&MongoConnection{client}

		mgo.GetCollection("log").Drop(context.Background())
		mgo.GetCollection("entry").Drop(context.Background())
		mgo.GetCollection("sectors").Drop(context.Background())
		mgo.GetCollection("Accounts").Drop(context.Background())

	}
	return mgo
}