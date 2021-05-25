package ledg_types

//func (e SectorEntry) MarshaBSON ()([]byte, error) {
//	ret:=bson.M {
//		"sectorEntry":"sectorEntry",
//	}
//	return bson.Marshal(ret)
//}


//func (s *Sector) UnmarshalBSON(data []byte) error {
//	var m bson.M
//	err := bson.Unmarshal(data, &m);	if err != nil {return err}
//
//	s.id=m["_id"].(string)
//	s.SectorNum=abi.SectorNumber(ledg_util.BsonM2Uint64(m["SectorNum"]))
//	s.SealProof=abi.RegisteredSealProof(ledg_util.BsonM2Uint64(m["SealProof"]))
//	s.SealedCID=ledg_util.BsonM2Cid(m["SealedCID"])
//
//	s.DealIDs=nil
//	s.Activation=abi.ChainEpoch(ledg_util.BsonM2Uint64(m["Activation"]))
//	s.Expiration=abi.ChainEpoch(ledg_util.BsonM2Uint64(m["Expiration"]))
//
//	s.PreCommitEpoch=abi.ChainEpoch(ledg_util.BsonM2Uint64(m["PreCommitEpoch"]))
//
//	s.PreCommitDeposit=ledg_util.BsonM2TokenAmount(m["PreCommitDeposit"])
//
//	s.DealWeight  =          ledg_util.BsonM2BigInt(m["DealWeight"])
//	s.VerifiedDealWeight=    ledg_util.BsonM2BigInt(m["VerifiedDealWeight"])
//	s.InitialPledge      =   ledg_util.BsonM2TokenAmount(m["InitialPledge"])
//	s.ExpectedDayReward   =  ledg_util.BsonM2TokenAmount(m["ExpectedDayReward"])
//	s.ExpectedStoragePledge =ledg_util.BsonM2TokenAmount(m["ExpectedStoragePledge"])
//
//
//	s.SectorPreCommitInfo=ledg_util.BsonM2Cid(m["SectorPreCommitInfo"])
//	s.SectorPreCommitOnChainInfo=ledg_util.BsonM2Cid(m["SectorPreCommitOnChainInfo"])
//	s.SectorOnChainInfo=ledg_util.BsonM2Cid(m["SectorOnChainInfo"])
//
//	s.Miner=ledg_util.BsonM2Address(m["Miner"])
//
//	//s.Entries=bson.un
//
//	//todo ledger deals bson unmarshal
//	//deal,err:=strconv.ParseUint(m["Deals"].(string),10,64);if err != nil {return err}
//	//e.Deals =abi.DealID(deal)
//
//	return nil
//}

//func (s Sector) MarshalBSON() ([]byte, error) {

	//sm:=bson.M{
	//	"_id"			:			  s.id,
	//	"SectorNumber":               s.SectorNum,
	//	"SealProof":                  s.SealProof,
	//	"SealedCID":                  s.SealedCID.String(),
	//	"DealIDs":                    nil,
	//	"Activation":                 s.Activation,
	//	"Expiration":                 s.Expiration,
	//	"PreCommitDeposit":           s.PreCommitDeposit.String(),
	//	"PreCommitEpoch":             s.PreCommitEpoch,
	//	"DealWeight":                 s.DealWeight.String(),
	//	"VerifiedDealWeight":         s.VerifiedDealWeight.String(),
	//	"InitialPledge":              s.InitialPledge.String(),
	//	"ExpectedDayReward":          s.ExpectedDayReward.String(),
	//	"ExpectedStoragePledge":      s.ExpectedStoragePledge.String(),
	//	"SectorPreCommitInfo":        s.SectorPreCommitInfo.String(),
	//	"SectorPreCommitOnChainInfo": s.SectorPreCommitOnChainInfo.String(),
	//	"SectorOnChainInfo":          s.SectorOnChainInfo.String(),
	//	"Miner":                      s.Miner.String(),
	//	"Entries":					  s.Entries,
	//}
	
	//d:=bson.M{
	//	//"_id":          e.id,
	//	"EntryCid":       cid,
	//	"_Address":     e._Address.String(),
	//	"Offset":      e.Offset.String(),
	//	"Method":   e.Method.String(),
	//	"Amount":       e.Amount.String(),
	//	"TotalAmount": e.TotalAmount,
	//	"Amount":      e.Amount,
	//	"balance":     e.balance,
	//	"GasFee":      e.GasFee.String(),
	//	"GasUsed":     strconv.FormatInt(e.GasUsed,10),
	//	"CallDepth":   strconv.FormatUint(e.CallDepth,10),
	//	"Sector":      bson.M{"miner":e.Sector.Miner.String(),"number":e.Sector.Number.String()},
	//	"Deals":        e.Deals,
	//	"Epoch":		   strconv.FormatUint(uint64(e.Epoch),10),
	//	"Nonce":		   strconv.FormatUint(uint64(e.Nonce),10),
	//	"MethodName": e.MethodName,
	//
	//}


	//return bson.Marshal(Addr{"Marshalled"})
	//return bson.Marshal(sm)
//}