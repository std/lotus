package ledg_types

import (
	"bytes"
)

func UnmarshalSectorPreCommitInfo(b []byte) (SectorPreCommitInfo,error) {
	ret:= SectorPreCommitInfo{}
	err:=ret.UnmarshalCBOR(bytes.NewReader(b))
	return ret,err
}

func CreateMinerParams(b []byte) (SectorPreCommitInfo,error) {
	ret:= SectorPreCommitInfo{}
	err:=ret.UnmarshalCBOR(bytes.NewReader(b))
	return ret,err
}
