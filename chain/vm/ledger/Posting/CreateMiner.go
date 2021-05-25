package Posting

import (
	"bytes"
	addr "github.com/filecoin-project/go-address"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	m "github.com/filecoin-project/lotus/chain/vm/ledger/models"
	"github.com/filecoin-project/specs-actors/v4/actors/builtin/power"
)

func (ledger *LedgerPosting) CreateMiner(p ledg_util.ActorMethodParams) {

	//ledger.AddEntry(ctx,Msg,methReturn,callDepth,"CreateMiner")
	msg:=p.Msg
	methReturn:=p.Ret

	llog.Info("CreateMiner impl")

	cm		:=	power.CreateMinerParams{}
	cm_ret	:=	power.CreateMinerReturn{}

	cm.UnmarshalCBOR(bytes.NewReader(msg.Params))
	cm_ret.UnmarshalCBOR(bytes.NewReader(methReturn))

	addrId,_:=addr.IDFromAddress(cm_ret.IDAddress)

	minerActor:=m.Account{
		ID:            int32(addrId),
		Address: cm_ret.IDAddress.String(),
		RobustAddress:       cm_ret.RobustAddress.String(),

		Name: "Miner: "+cm_ret.IDAddress.String(),

		CreationEpoch: ledger.Epoch,
		//Balance:       ledg.DimBalance{
		//	ledg.Available:ledg.FilAmountFromInt(11),
		//	ledg.InitialPledge: ledg.FilAmountFromInt(22),
		//},
		//SectorCounts: ledg.SectorCounts{Active: 11},
		//PowerBalance:ledg.PowerBalance{VerifiedStoragePower: ledg.StoragePowerFromInt(123)},
		//ActorTypeConst: models.Miner,
		//TotalReward: ledg.FilAmountFromInt(123),
		//MessagesCount: 999,

		//Stats: ledg.MiningStats{
		//	PowerGrowth:           ledg.StoragePowerFromInt(32*1024*1024),
		//	BlocksMined:           8,
		//	MiningEfficiencyPerTB: ledg.FilAmountFromInt(100025),
		//	WinCount:              10,
		//	MinerEquivalent:       15.89,
		//},
		//Properties: ledg.MinerProperties{
		//	PeerId:  "12D3KooWRudzcMVAgZapWJXPDKDgUMZbbsQFeDHcAFQwxAWrtQTV",
		//	Owner:   ledg.NewAddressFromString("f3wkx7jksblo4kehbklknlivm6pniartluv3nqz3mpwjj5dyfu55pctvyxxejkjgki7qp3r3thxt3wk73hwsua"),
		//	Worker:  ledg.NewAddressFromString("f3rzvyvt6lnamvx7dc4ulrukenudq7ywzrdjelnpouwzsurqnep5vcick7x72w4tslmyvqbx2mkemkqalbtswq"),
		//	Region:  "europe",
		//	Country: "lv",
		//	Ip:      "8.8.8.8",
		//},
	}


	ledger.insert(&minerActor,true)

	e:=ledger.minerEntryTemplate(p,0,false)

	e.TxId=ledger.CurrentTxId
	e.CallDepth=p.Depth
	e.MinerId=int32(addrId)
	ledger.insert(&e,true)

	e2:=ledger.minerEntryTemplate(p,0,true)

	e2.TxId=ledger.CurrentTxId
	e2.CallDepth=p.Depth
	e2.MinerId=int32(addrId)
	ledger.insert(&e2,true)



}