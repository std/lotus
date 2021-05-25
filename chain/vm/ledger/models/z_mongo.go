package models

import (
	"github.com/filecoin-project/go-state-types/abi"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
)

type SectorMongo struct {
	ID        string `bson:"_id"`
	SectorNum ledg.SectorNumber `bson:"SectorNum"`


	MinerId uint32
	Miner   Account //`gorm:"foreignKey:MId"`

	PreCommitEpoch 	abi.ChainEpoch  `bson:"PreCommitEpoch"`
	CommitEpoch		abi.ChainEpoch  `bson:"CommitEpoch"`
	ActivationEpoch		abi.ChainEpoch  `bson:"ActivationEpoch"`
	ExpirationEpoch abi.ChainEpoch  `bson:"ExpirationEpoch"`

	//Deals  []ledg.DealID `bson:"Deals"`
	PreCommitDeposit ledg.FilAmount `sql:"type:varchar"`// `bson:"PreCommitDeposit"`//???
	InitialPledge ledg.FilAmount `sql:"type:varchar"`//`bson:"InitialPledge"`//???

	SealProof abi.RegisteredSealProof `bson:"SealProof"`//???

	Status uint `bson:"Status"`//???
	// Sum of raw byte power for a miner's sectors.
	RawBytePower ledg.StoragePower `sql:"type:varchar"`//`bson:"StoragePower"`
	// Sum of quality adjusted power for a miner's sectors.
	QualityAdjPower ledg.StoragePower `sql:"type:varchar"`//`bson:"QualityAdjPower"`

	PreCommitFee ledg.FilAmount `sql:"type:varchar"`//`bson:"PreCommitFee"`
	CommitFee ledg.FilAmount `sql:"type:varchar"`//`bson:"CommitFee"`
	//	Winning []int//??

	//SectorPreCommitInfo ledg.SectorPreCommitInfo  `sql:"type:varchar"`//`bson:"_SectorPreCommitInfo"`
}

//func  (l *PowerLedger) insertMinerMongo(m models.Account) error {
//	con := ledg_util.GetOrCreateMongoConnection()
//	_, err := con.GetCollection("Accounts").InsertOne(context.TODO(), m)
//	if err != nil {
//		ledg_util.Log(bson.M{
//			"source":    "PowerLedger.insertMiner",
//			"Miner":     m.ID,
//			"err":       err.Error(),
//		})
//	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
//	return err
//}