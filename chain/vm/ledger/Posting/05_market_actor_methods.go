package Posting

import (
	"bytes"
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	ledg_util "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ipfs/go-cid"
)
func (l *MarketLedger) GetPublishStorageDealsReturn(b []byte) (ledg.PublishStorageDealsReturn,error) {
	ret:=ledg.PublishStorageDealsReturn{}
	err:=ret.UnmarshalCBOR(bytes.NewReader(b))
	return ret,err
}


type MarketState struct {
	Proposals cid.Cid // AMT[DealID]DealProposal
	States    cid.Cid // AMT[DealID]DealState

	// PendingProposals tracks dealProposals that have not yet reached their deal start date.
	// We track them here to ensure that miners can't publish the same deal proposal twice
	PendingProposals cid.Cid // HAMT[DealCid]DealProposal

	// Total amount held in escrow, indexed by actor address (including both locked and unlocked amounts).
	EscrowTable cid.Cid // BalanceTable

	// Amount locked, indexed by actor address.
	// Note: the amounts in this table do not affect the overall amount in escrow:
	// only the _portion_ of the total escrow amount that is locked.
	LockedTable cid.Cid // BalanceTable

	NextID abi.DealID

	// Metadata cached for efficient iteration over deals.
	DealOpsByEpoch cid.Cid // SetMultimap, HAMT[epoch]Set
	LastCron       abi.ChainEpoch

	// Total Client Collateral that is locked -> unlocked when deal is terminated
	TotalClientLockedCollateral abi.TokenAmount
	// Total Provider Collateral that is locked -> unlocked when deal is terminated
	TotalProviderLockedCollateral abi.TokenAmount
	// Total storage fee that is locked in escrow -> unlocked when payments are made
	TotalClientStorageFee abi.TokenAmount

	///////////////////////////////////////
}

type MarketBalanceEntry struct {
	ProviderCollateral abi.TokenAmount
	ClientCollateral abi.TokenAmount
	PiecSize abi.PaddedPieceSize
}


type MarketLedger struct{

	//Epoch   abi.ChainEpoch
	State   MarketState
	Balance map[address.Address]map[abi.DealID]ledg.StorageDeal
}

func(l *MarketLedger) InsertMarketLedgerEntry(msg ,originMsg *types.Message, sd ledg.StorageDeal, depth uint64) MarketLedgerEntry {
	//ml:=l.MarketLedger
	e:= MarketLedgerEntry{
		OriginMsgCid: originMsg.Cid(),
		MsgCid:       msg.Cid(),
		//id:        "",
		Miner:     msg.To,
		Sector:    abi.SectorID{},
		MethodNum: msg.Method,
		Value:     msg.Value,


		CallDepth: depth,
		DealId:		sd.Id,
		//Epoch:     Msg.,
		EntryType: 0,
		Method:    ledg_util.GetMethodName(msg),
	}
	//ledg_util.GetOrCreateMongoConnection().insertDeal(e)
	return e
}

func  (l *MarketLedger) insertDeal(s ledg.StorageDeal) error{

	con:=ledg_util.GetOrCreateMongoConnection()
	_,err:=con.GetCollection("deals").InsertOne(context.TODO(), s)


	if err!=nil {
		ledg_util.Log(bson.M{
			"source":"MarketLedger.insertDeal",
			"Miner":s.Provider,
			"SectorNum":s.SectorId.Number,
			"StartEpoch":s.StartEpoch,
			"EndEpoch":s.EndEpoch,
			"msgCid":" not implemented ",
			"err":err.Error(),

		})
	} //else {fmt.Println("Inserted a single document: ", res.InsertedID)}
	return err
}

func (ml *MarketLedger) CreateDealItem(p ledg.DealProposal,dealId abi.DealID) ledg.StorageDeal {

	p.Id=ledg.DealID(dealId)
	sd:= ledg.StorageDeal{

		Id:		              dealId,
		PieceCID:             p.PieceCID,
		SectorId:             abi.SectorID{},
		Status:               0,
		PieceSize:            p.PieceSize,
		VerifiedDeal:         p.VerifiedDeal,
		Client:               p.Client,
		Provider:             p.Provider,
		Label:                p.Label,
		StartEpoch:           p.StartEpoch,
		EndEpoch:             p.EndEpoch,
		StoragePricePerEpoch: p.StoragePricePerEpoch,
		ProviderCollateral:   p.ProviderCollateral,
		ClientCollateral:     p.ClientCollateral,

		DealProposal: p,
	}

	ml.insertDeal(sd)
	return sd
}

func(l *LedgerPosting) MarketActorConstructor(p ledg_util.ActorMethodParams)   {}
func(l *LedgerPosting) VerifyDealsForActivation(p ledg_util.ActorMethodParams) {}
func(l *LedgerPosting) ActivateDeals(p ledg_util.ActorMethodParams)            {}
func(l *LedgerPosting) OnMinerSectorsTerminate(p ledg_util.ActorMethodParams)  {}
func(l *LedgerPosting) ComputeDataCommitment(p ledg_util.ActorMethodParams)    {}

func(l *LedgerPosting) AddBalance(p ledg_util.ActorMethodParams)            {}
func(l *LedgerPosting) WithdrawMarketBalance(p ledg_util.ActorMethodParams) {}
func(l *LedgerPosting) CronTick(p ledg_util.ActorMethodParams)              {}


func(l *LedgerPosting) PublishStorageDeals( p ledg_util.ActorMethodParams) {

	ret:=p.Ret
	msg:=p.Msg
	depth:=p.Depth

	originMsg:=l.OriginMsg

	dealsRet,_:=l.GetPublishStorageDealsReturn(ret)
	dealIds:=dealsRet.IDs

	publish:= ledg.PublishStorageDealsParams{}
	_=publish.UnmarshalCBOR(bytes.NewReader(msg.Params))
	deals:=publish.Deals

	for i, d:=range deals {
		storageDeal:=l.MarketLedger.CreateDealItem(d.Proposal,dealIds[i])
		_=l.MarketLedger.InsertMarketLedgerEntry(msg,originMsg,storageDeal,depth)
	}
}


type MarketLedgerEntry struct {
	OriginMsgCid cid.Cid
	MsgCid       cid.Cid
	Id           string

	Miner 			address.Address
	Sector    		abi.SectorID

	//_Address         address._Address
	//Offset          address._Address
	//TargetActorType Protocol
	MethodNum       abi.MethodNum
	Value           abi.TokenAmount


	CallDepth uint64

	DealId     abi.DealID

	Epoch     abi.ChainEpoch
	EntryType ledg.MarketEntryType
//	Nonce     uint64
	Method    string

//	invariant abi.TokenAmount//should be zero
}




//'deal lifecyce

//func (l *MarketLedger) getDealsFromParams(Msg *types.Message) []abi.DealID {
//	var Ret []abi.DealID
//	if Msg.To.String()=="f05"{
//		method:=ledger.GetMethodName(Msg)
//		switch method {
//		case "AddBalance":
//		case "WithdrawBalance":
//		case "PublishStorageDeals":
//			publish:= PublishStorageDealsParams{}
//			_=publish.UnmarshalCBOR(bytes.NewReader(Msg.Params))
//			Ret=publish.Deals
//		case "ComputeDataCommitment":
//		case "VerifyDealsForActivation":
//		case "ActivateDeals":
//		case "OnMinerSectorsTerminate":
//		case "CronTick":
//		}
//
//	}
//	return nil
//}

