package ledger

import (
	"github.com/filecoin-project/lotus/chain/vm/ledger/ledg-util"
	"gorm.io/gorm"
)

type GL interface {
 	GetPgDatabase() *gorm.DB
}

type PowerMethods interface {
	PowerConstructor(params ledg_util.ActorMethodParams)         //1
	CreateMiner(params ledg_util.ActorMethodParams)              //2
	UpdateClaimedPower(params ledg_util.ActorMethodParams)       //3
	EnrollCronEvent(params ledg_util.ActorMethodParams)          //4
	OnEpochTickEnd(params ledg_util.ActorMethodParams)           //5
	UpdatePledgeTotal(params ledg_util.ActorMethodParams)        //6
	Deprecated1(params ledg_util.ActorMethodParams)              //7
	SubmitPoRepForBulkVerify(params ledg_util.ActorMethodParams) //8
	CurrentTotalPower(params ledg_util.ActorMethodParams)        //9

}

type MinerMethods interface {
	MinerActorConstructor(params ledg_util.ActorMethodParams) //1
	ControlAddresses(params ledg_util.ActorMethodParams)      //2
	ChangeWorkerAddress(params ledg_util.ActorMethodParams)   //3

	ChangePeerID(params ledg_util.ActorMethodParams)             //4
	SubmitWindowedPoSt(params ledg_util.ActorMethodParams)       //5
	PreCommitSector(params ledg_util.ActorMethodParams)          //6
	ProveCommitSector(params ledg_util.ActorMethodParams)        //7
	ExtendSectorExpiration(params ledg_util.ActorMethodParams)   //8
	TerminateSectors(params ledg_util.ActorMethodParams)         //9
	DeclareFaults(params ledg_util.ActorMethodParams)            //10
	DeclareFaultsRecovered(params ledg_util.ActorMethodParams)   //11
	OnDeferredCronEvent(params ledg_util.ActorMethodParams)      //12
	CheckSectorProven(params ledg_util.ActorMethodParams)        //13
	ApplyRewards(params ledg_util.ActorMethodParams)             //14
	ReportConsensusFault(params ledg_util.ActorMethodParams)     //15
	WithdrawBalance(params ledg_util.ActorMethodParams)          //16
	ConfirmSectorProofsValid(params ledg_util.ActorMethodParams) //17
	ChangeMultiaddrs(params ledg_util.ActorMethodParams)         //18
	CompactPartitions(params ledg_util.ActorMethodParams)        //19
	CompactSectorNumbers(params ledg_util.ActorMethodParams)     //20
	ConfirmUpdateWorkerKey(params ledg_util.ActorMethodParams)   //21
	RepayDebt(params ledg_util.ActorMethodParams)                //22
	ChangeOwnerAddress(params ledg_util.ActorMethodParams)       //23
	DisputeWindowedPoSt(params ledg_util.ActorMethodParams)      //24

}
type MarketMethods interface {
	MarketActorConstructor(params ledg_util.ActorMethodParams)   //1
	AddBalance(params ledg_util.ActorMethodParams)               //2
	WithdrawMarketBalance(params ledg_util.ActorMethodParams)    //3 original name WithdrawBalance the same as miner meth.
	PublishStorageDeals (params ledg_util.ActorMethodParams)     //4
	VerifyDealsForActivation(params ledg_util.ActorMethodParams) //5
	ActivateDeals(params ledg_util.ActorMethodParams)            //6
	OnMinerSectorsTerminate(params ledg_util.ActorMethodParams)  //7
	ComputeDataCommitment(params ledg_util.ActorMethodParams)    //8
	CronTick(params ledg_util.ActorMethodParams)                 //9
}

type RewardActor interface {
	RewardActorConstructor(params ledg_util.ActorMethodParams) //1
	AwardBlockReward(params ledg_util.ActorMethodParams)       //2
	ThisEpochReward(params ledg_util.ActorMethodParams)        //3
	UpdateNetworkKPI(params ledg_util.ActorMethodParams)       //4
}

type CronActor interface {
	CronActorConstructor(params ledg_util.ActorMethodParams) //1
	EpochTick(params ledg_util.ActorMethodParams)            //2
}
type AccountActor interface {
	AccountActorConstructor(params ledg_util.ActorMethodParams) //1
	PubkeyAddress(params ledg_util.ActorMethodParams)           //2
}

type SystemActor interface {
	SystemActorConstructor(params ledg_util.ActorMethodParams) //1
}

type MultisigActor interface {
	MultisigActorConstructor(params ledg_util.ActorMethodParams)    //1
	Propose(params ledg_util.ActorMethodParams)                     //2
	Approve(params ledg_util.ActorMethodParams)                     //3
	Cancel(params ledg_util.ActorMethodParams)                      //4
	AddSigner(params ledg_util.ActorMethodParams)                   //5
	RemoveSigner(params ledg_util.ActorMethodParams)                //6
	SwapSigner(params ledg_util.ActorMethodParams)                  //7
	ChangeNumApprovalsThreshold(params ledg_util.ActorMethodParams) //8
	LockBalance(params ledg_util.ActorMethodParams)                 //9

}

type PaychActor interface {
	PaychActorConstructor(params ledg_util.ActorMethodParams) //1
	UpdateChannelState(params ledg_util.ActorMethodParams)    //2
	Settle(params ledg_util.ActorMethodParams)                //3
	Collect(params ledg_util.ActorMethodParams)               //4
}

type VerifiedRegistryActor interface {
	VerifiedRegistryActorConstructor(params ledg_util.ActorMethodParams) //1
	AddVerifier(params ledg_util.ActorMethodParams)                      //3
	RemoveVerifier(params ledg_util.ActorMethodParams)                   //3
	AddVerifiedClient(params ledg_util.ActorMethodParams)                //3
	UseBytes(params ledg_util.ActorMethodParams)                         //3
	RestoreBytes(params ledg_util.ActorMethodParams)                     //3
}

