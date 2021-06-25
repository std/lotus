package z_miner_ledger

import (
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/ipfs/go-cid"
)

type State struct {
	// Information not related to sectors.
	Info cid.Cid

	PreCommitDeposits abi.TokenAmount // Total funds locked as PreCommitDeposits
	LockedFunds       abi.TokenAmount // Total rewards and added funds locked in vesting table

	VestingFunds cid.Cid // VestingFunds (Vesting Funds schedule for the miner).

	FeeDebt abi.TokenAmount // Absolute value of debt this miner owes from unpaid fees

	InitialPledge abi.TokenAmount // Sum of initial pledge requirements of all active sectors

	// Sectors that have been pre-committed but not yet proven.
	PreCommittedSectors cid.Cid // Map, HAMT[SectorNumber]SectorPreCommitOnChainInfo

	// PreCommittedSectorsExpiry maintains the state required to expire PreCommittedSectors.
	PreCommittedSectorsExpiry cid.Cid // BitFieldQueue (AMT[Epoch]*BitField)

	// Allocated sector IDs. Sector IDs can never be reused once allocated.
	AllocatedSectors cid.Cid // BitField

	// Information for all proven and not-yet-garbage-collected sectors.
	//
	// Sectors are removed from this AMT when the partition to which the
	// sector belongs is compacted.
	Sectors cid.Cid // Array, AMT[SectorNumber]SectorOnChainInfo (sparse)

	// The first epoch in this miner's current proving period. This is the first epoch in which a PoSt for a
	// partition at the miner's first deadline may arrive. Alternatively, it is after the last epoch at which
	// a PoSt for the previous window is valid.
	// Always greater than zero, this may be greater than the current epoch for genesis miners in the first
	// WPoStProvingPeriod epochs of the chain; the epochs before the first proving period starts are exempt from Window
	// PoSt requirements.
	// Updated at the end of every period by a cron callback.
	ProvingPeriodStart abi.ChainEpoch

	// Index of the deadline within the proving period beginning at ProvingPeriodStart that has not yet been
	// finalized.
	// Updated at the end of each deadline window by a cron callback.
	CurrentDeadline uint64

	// The sector numbers due for PoSt at each deadline in the current proving period, frozen at period start.
	// New sectors are added and expired ones removed at proving period boundary.
	// Faults are not subtracted from this in state, but on the fly.
	Deadlines cid.Cid

	// Deadlines with outstanding fees for early sector termination.
	EarlyTerminations bitfield.BitField
}



type SectorOnChainInfo struct {
	SectorNumber          abi.SectorNumber
	SealProof             abi.RegisteredSealProof
	SealedCID             cid.Cid
	DealIDs               []abi.DealID
	Activation            abi.ChainEpoch
	Expiration            abi.ChainEpoch
	DealWeight            abi.DealWeight
	VerifiedDealWeight    abi.DealWeight
	InitialPledge         abi.TokenAmount
	ExpectedDayReward     abi.TokenAmount
	ExpectedStoragePledge abi.TokenAmount
}



type ProveCommitSectorParams struct {
	SectorNumber abi.SectorNumber
	Proof        []byte
}


//meth PublishStorageDeals

type ComputeDataCommitmentParams struct {
	DealIDs    []abi.DealID
	SectorType abi.RegisteredSealProof
}

type VerifyDealsForActivationParams struct {
	DealIDs      []abi.DealID
	SectorExpiry abi.ChainEpoch
	SectorStart  abi.ChainEpoch
}

type ActivateDealsParams struct {
	DealIDs      []abi.DealID
	SectorExpiry abi.ChainEpoch
}

type OnMinerSectorsTerminateParams struct {
	Epoch   abi.ChainEpoch
	DealIDs []abi.DealID
}



