package ledg_types

//type EntryType struct{
//	NewSectorOnChain EntryTypeConst
//}

type EntryTypeConst string
const (
	PrecommitSector EntryTypeConst = "PrecommitSector"
	CommitSector ="CommitSector"
	TerminateSector="TerminateSector"
	ExtendSectorExpiration="ExtendSectorExpiration"
	ExpireSector="ExpireSector"
	DeclareFaults="DeclareFaults"
	DeclareFaultsRecovered="DeclareFaultsRecovered"
	ConfirmSectorProofsValid="ConfirmSectorProofsValid"


)
var SectorEntryType=struct {
	PreCommit EntryTypeConst
	ProveCommit EntryTypeConst
	Terminate EntryTypeConst
	ExtendSectorExpiration EntryTypeConst
	DeclareFaults EntryTypeConst
	DeclareFaultsRecovered EntryTypeConst
	ConfirmSectorProofsValid EntryTypeConst //unlock precommit depo here???


}{
	PreCommit:                PrecommitSector,
	ProveCommit:              CommitSector,
	Terminate:                TerminateSector,
	ExtendSectorExpiration:   ExtendSectorExpiration,
	DeclareFaults:            DeclareFaults,
	DeclareFaultsRecovered:   DeclareFaultsRecovered,
	ConfirmSectorProofsValid: ConfirmSectorProofsValid,
}

var MinerEntryType=struct {
	PreCommit EntryTypeConst
	ProveCommit EntryTypeConst
	Terminate EntryTypeConst
}{
	PreCommit:                PrecommitSector,
	ProveCommit:              CommitSector,
	Terminate: 				  TerminateSector,
}
