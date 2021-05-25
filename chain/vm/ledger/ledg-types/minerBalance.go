package ledg_types

type PowerBalance struct {
	StoragePower StoragePower
	VerifiedStoragePower StoragePower
}


type SectorCounts struct {
	PreCommitted int
	Committed int
	Active int
	Faults int
	DeclaredFault int
	Recovered int
	Expired int
}


