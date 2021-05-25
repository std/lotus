package market_ledger

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/vm/ledger/Posting"
	ledg "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	"github.com/ipfs/go-cid"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)
type MongoResult struct {
	Result string
}

const (
	DealOffered = iota//?
	DealDataTransferred//?
	DealPublished
	DealSectorPreCommited
	DealSectorCommited
	DealSectorTerminated
	DealSectorExpired
)

type Deal struct {
	abi.DealID
	//DealId abi.DealID
	SectorId abi.SectorID
	Status   Posting.DealStatus

	PieceCID             cid.Cid
	PieceSize            abi.PaddedPieceSize
	VerifiedDeal         bool
	Client               address.Address
	Provider             address.Address
	Label                string
	StartEpoch           abi.ChainEpoch
	EndEpoch             abi.ChainEpoch
	StoragePricePerEpoch abi.TokenAmount
	ProviderCollateral   abi.TokenAmount
	ClientCollateral     abi.TokenAmount

	//market.DealState
	//type DealState struct {
	SectorStartEpoch abi.ChainEpoch // -1 if not yet included in proven sector
	LastUpdatedEpoch abi.ChainEpoch // -1 if deal state never updated
	SlashEpoch       abi.ChainEpoch // -1 if deal never slashed
}


func (d *Deal) FromMarketDeal(deal api.MarketDeal,dealId abi.DealID) error {


		d.DealID=               dealId
		d.PieceCID=             deal.Proposal.PieceCID
		d.PieceSize=            deal.Proposal.PieceSize
		d.VerifiedDeal=         deal.Proposal.VerifiedDeal
		d.Client=               deal.Proposal.Client
		d.Provider=             deal.Proposal.Provider
		d.Label=                deal.Proposal.Label
		d.StartEpoch=           deal.Proposal.StartEpoch
		d.EndEpoch=             deal.Proposal.EndEpoch
		d.StoragePricePerEpoch= deal.Proposal.StoragePricePerEpoch
		d.ProviderCollateral=   deal.Proposal.ProviderCollateral
		d.ClientCollateral=     deal.Proposal.ClientCollateral

		d.SectorStartEpoch= deal.State.SectorStartEpoch // -1 if not yet included in proven sector
		d.LastUpdatedEpoch= deal.State.LastUpdatedEpoch // -1 if deal state never updated
		d.SlashEpoch=       deal.State.SlashEpoch // -1 if deal never slashed
	//d=&Deals{
	//	DealID:               dealId,
	//	PieceCID:             deal.Proposal.PieceCID,
	//	PieceSize:            deal.Proposal.PieceSize,
	//	VerifiedDeal:         deal.Proposal.VerifiedDeal,
	//	Client:               deal.Proposal.Client,
	//	Provider:             deal.Proposal.Provider,
	//	Label:                deal.Proposal.Label,
	//	StartEpoch:           deal.Proposal.StartEpoch,
	//	EndEpoch:             deal.Proposal.EndEpoch,
	//	StoragePricePerEpoch: deal.Proposal.StoragePricePerEpoch,
	//	ProviderCollateral:   deal.Proposal.ProviderCollateral,
	//	ClientCollateral:     deal.Proposal.ClientCollateral,
	//
	//	SectorStartEpoch: deal.State.SectorStartEpoch, // -1 if not yet included in proven sector
	//	LastUpdatedEpoch: deal.State.LastUpdatedEpoch, // -1 if deal state never updated
	//	SlashEpoch:       deal.State.SlashEpoch, // -1 if deal never slashed
	//}

	return nil
}

func (deal Deal) MarshalBSON() ([]byte, error){
	ret:=bson.M{
		"_id":					strconv.FormatUint(uint64(deal.DealID),10),
		"PieceCID":             deal.PieceCID.String(),
		"PieceSize":            strconv.FormatUint(uint64(deal.PieceSize),10),
		"VerifiedDeal":         deal.VerifiedDeal,
		"Client":               deal.Client.String(),
		"Provider":             deal.Provider.String(),
		"Label":                deal.Label,
		"StartEpoch":           deal.StartEpoch.String(),
		"EndEpoch":             deal.EndEpoch.String(),
		"StoragePricePerEpoch": deal.StoragePricePerEpoch.String(),
		"ProviderCollateral":   deal.ProviderCollateral.String(),
		"ClientCollateral":     deal.ClientCollateral.String(),

		"SectorStartEpoch":  	deal.SectorStartEpoch,// -1 if not yet included in proven sector
		"LastUpdatedEpoch": 	deal.LastUpdatedEpoch, // -1 if deal state never updated
		"SlashEpoch":       	deal.SlashEpoch, // -1 if deal never slashed

	}
	return bson.Marshal(ret)
}

func parseAddress(v interface{}) address.Address {
	addr,err:=address.NewFromString(v.(string));		if err != nil {return address.Address{}}
	return addr
}

func parseUint64 ( v interface{}) uint64 {
	ret,_:= strconv.ParseUint(v.(string),10,64)
	return ret
}

func (e *Deal) UnmarshalBSON(data []byte) error {
	var m bson.M
	err := bson.Unmarshal(data, &m);	if err != nil {return err}

	dealId,_:=strconv.ParseUint(m["_id"].(string),10,64)
	e.DealID=abi.DealID(dealId)


		e.DealID		=				abi.DealID(dealId)
		e.PieceCID,_	=            	cid.Parse(m["PieceCID"].(string))
		e.PieceSize		=            	abi.PaddedPieceSize(parseUint64(m["PieceSize"]))

		e.VerifiedDeal=         m["VerifiedDeal"].(bool)
		e.Client=				parseAddress(m["Client"])
		e.Provider= 			parseAddress(m["Provider"])

		e.Label=                m["Label"].(string)
		e.StartEpoch=           abi.ChainEpoch(parseUint64(m["StartEpoch"]))
		e.EndEpoch =	        abi.ChainEpoch(parseUint64(m["EndEpoch"]))
		e.StoragePricePerEpoch= abi.TokenAmount(ledg.BsonM2TokenAmount(m["StoragePricePerEpoch"]))
		e.ProviderCollateral= abi.TokenAmount(ledg.BsonM2TokenAmount(m["ProviderCollateral"]))
		e.ClientCollateral= abi.TokenAmount(ledg.BsonM2TokenAmount(m["ClientCollateral"]))
		e.SectorStartEpoch=     abi.ChainEpoch(parseUint64(m["SectorStartEpoch"]))
		e.LastUpdatedEpoch=		abi.ChainEpoch(parseUint64(m["LastUpdatedEpoch"]))
		e.SlashEpoch=           abi.ChainEpoch(parseUint64(m["SlashEpoch"]))


	//e.EntryCid,_=cid.Parse(e.id)
	//method,_:=strconv.ParseUint(m["Method"].(string),10,64)
	////fmt.Println(m)
	//addr,err:=address.NewFromString(m["_Address"].(string));		if err != nil {return err}
	//e._Address=addr
	//offset,err:=address.NewFromString(m["Offset"].(string));		if err != nil {return err}
	//e.Offset=offset
	//e.Method = abi.Method(method)
	//e.Amount,_ = BsonM2TokenAmount(m["Amount"].(string))
	//
	//e.TotalAmount,_ =DimBalanceFromBsonM(m["TotalAmount"])
	//e.Amount,_=DimBalanceFromBsonM(m["Amount"])
	//e.balance,_=DimBalanceFromBsonM(m["balance"])
	//e.GasFee,_=BsonM2TokenAmount(m["GasFee"].(string))
	//e.GasUsed,_=strconv.ParseInt(m["GasUsed"].(string),10,64)
	//e.CallDepth,_=strconv.ParseUint(m["CallDepth"].(string),10,64)
	//
	//e.Sector,_=SectorFromBsonM(m["Sector"])
	//
	//deal,err:=strconv.ParseUint(m["Deals"].(string),10,64);if err != nil {return err}
	//
	//e.Deals=abi.DealID(deal)

	return nil
}

func InsertDeal(dealId abi.DealID,deal api.MarketDeal) error {

	//mgo:= mongo.GetOrCreateMongoConnection()
	//var d *Deals
	d:=&Deal{}
	d.FromMarketDeal(deal,dealId)
	//mgo.insertDeal(*d)
	return nil//mgo.insertDeal(*d)
}