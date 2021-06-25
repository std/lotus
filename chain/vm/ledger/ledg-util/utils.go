package ledg_util

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/types"
	ledg_types "github.com/filecoin-project/lotus/chain/vm/ledger/ledg-types"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)


var llog			= logging.Logger("gledger")

func Llog(s string){
	llog.Info(s)
}
func Llogf(template string, args ...interface{}){
	llog.Infof(template,args...)
}



func GetActorType(addr address.Address) string{

	ntwk:=address.MainnetPrefix
	if (address.CurrentNetwork==address.Testnet){
		ntwk=address.TestnetPrefix
	}
	ret:="account"
	if addr.String()==ntwk+"00" {ret="system"} else
	if addr.String()==ntwk+"01" {ret="init"} else
	if addr.String()==ntwk+"02" {ret="reward"} else
	if addr.String()==ntwk+"03" {ret="cron"} else
	if addr.String()==ntwk+"04" {ret="power"} else
	if addr.String()==ntwk+"05" {ret="market"} else
	if addr.String()==ntwk+"06" {ret="registry"} else
	if addr.String()==ntwk+"099" {ret="burnt"} else
	if IsMiner(addr) {ret="miner"}else
	if IsPaych(addr) {ret="paych"}

	return ret
}

func GetPowerEntryType(msg *types.Message) ledg_types.EntryTypeConst{

	methName:=GetMethodName(msg)
	entryMap := map[string]ledg_types.EntryTypeConst{
		"CreateMiner":   "CreateMiner",
		"PreCommitSector":"PreCommitSector",
	}
	return entryMap[methName]

}


func GetEntryType(msg *types.Message) ledg_types.EntryTypeConst{

	methName:=GetMethodName(msg)
	entryMap := map[string]ledg_types.EntryTypeConst{
		"PreCommitSector": ledg_types.MinerEntryType.PreCommit,
		"CreateMiner":   "CreateMiner",
	}
	return entryMap[methName]
}


func  GetMethodName( msg *types.Message) string{

	ntwk:=address.MainnetPrefix
	if (address.CurrentNetwork==address.Testnet){
		ntwk=address.TestnetPrefix
	}
	addr:=msg.To

	ret:=msg.Method.String()
	meth:=msg.Method

	if addr.String()==ntwk+"00" {
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="PubkeyAddress"
		}
	} else	if addr.String()==ntwk+"01" {
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="Exec"
		}
	} else if addr.String()==ntwk+"02" {
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="AwardBlockReward"
		case 3:ret="ThisEpochReward"
		case 4:ret="UpdateNetworkKPI"
		}
	} else	if addr.String()==ntwk+"03" {
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="EpochTick"
		}

	} else if addr.String()==ntwk+"04" {
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="CreateMiner"
		case 3:ret="UpdateClaimedPower"
		case 4:ret="EnrollCronEvent"
		case 5:ret="OnEpochTickEnd"
		case 6:ret="UpdatePledgeTotal"
		case 7:ret="Deprecated1"
		case 8:ret="SubmitPoRepForBulkVerify"
		case 9:ret="CurrentTotalPower"
		}
	}else if addr.String()==ntwk+"05"{
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="AddBalance"
		case 3:ret="WithdrawBalance"
		case 4:ret="PublishStorageDeals"
		case 5:ret="VerifyDealsForActivation"
		case 6:ret="ActivateDeals"
		case 7:ret="OnMinerSectorsTerminate"
		case 8:ret="ComputeDataCommitment"
		case 9:ret="CronTick"
		}
	}else if addr.String()==ntwk+"06"{
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="AddVerifier"
		case 3:ret="RemoveVerifier"
		case 4:ret="AddVerifiedClient"
		case 5:ret="UseBytes"
		case 6:ret="RestoreBytes"

		}
	}else if IsMiner(addr) {

		switch meth {
		case 1:ret="Constructor"
		case 2:ret="ControlAddresses"
		case 3:ret="ChangeWorkerAddress"
		case 4:ret="ChangePeerID"
		case 5:ret="SubmitWindowedPoSt"
		case 6:ret="PreCommitSector"
		case 7:ret="ProveCommitSector"
		case 8:ret="ExtendSectorExpiration"
		case 9:ret="TerminateSectors"
		case 10:ret="DeclareFaults"
		case 11:ret="DeclareFaultsRecovered"
		case 12:ret="OnDeferredCronEvent"
		case 13:ret="CheckSectorProven"
		case 14:ret="ApplyRewards"
		case 15:ret="ReportConsensusFault"
		case 16:ret="WithdrawBalance"
		case 17:ret="ConfirmSectorProofsValid"
		case 18:ret="ChangeMultiaddrs"
		case 19:ret="CompactPartitions"
		case 20:ret="CompactSectorNumbers"
		case 21:ret="ConfirmUpdateWorkerKey"
		case 22:ret="RepayDebt"
		case 23:ret="ChangeOwnerAddress"
		case 24:ret="DisputeWindowedPoSt"
		}
	}else if isMultisig(addr) {
		switch meth {
		case 1:ret="Constructor"
		case 2:ret="Propose"
		case 3:ret="Approve"
		case 4:ret="Cancel"
		case 5:ret="AddSigner"
		case 6:ret="RemoveSigner"
		case 7:ret="SwapSigner"
		case 8:ret="ChangeNumApprovalsThreshold"
		case 9:ret="LockBalance"
	}
} else if meth==0  { ret = "send"}

	return ret
}
func  IsMiner(addr address.Address) bool{
	isMiner:= addr.Protocol()==address.ID
	return isMiner
	//return addr.String()[0:2]=="t0" && ( len(addr.String()) ==7 || len(addr.String()) ==8)
}
func isMultisig(addr address.Address) bool {
	return false
	panic("isMultisig not implemented")
}
func IsPaych(addr address.Address) bool {
	return false
	panic("isPaych not implemented")
}




func  BsonM2Uint64(data interface{})uint64{
	ret,_:=strconv.ParseUint(data.(string),10,64)
	return ret
}

func  BsonM2BigInt(data interface{})big.Int{
	ret, err := big.FromString(data.(string))
	if err!=nil {return big.NewInt(0)}
	return ret
}

func  BsonM2Cid(data interface{})cid.Cid{
	ret,_:=data.(string)
	Cid,_:=cid.Parse(ret)
	return Cid
}

func  BsonM2Address(data interface{})address.Address{
	str,_:=data.(string)
	ret,_:=address.NewFromString(str)
	return ret
}

func  Log(m bson.M){
	mgo:=GetOrCreateMongoConnection()
	mgo.InsertLogEntry(m)
}




func MaxOf(i1,i2 int) int {
	max := i1
	if max<i2 {max=i2}

	return max
}