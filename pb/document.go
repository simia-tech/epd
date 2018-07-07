package pb

import "github.com/golang/protobuf/proto"

func (ud *UnlockedDocument) MarshalBinary() ([]byte, error) {
	return proto.Marshal(ud)
}

func (ld *LockedDocument) MarshalBinary() ([]byte, error) {
	return proto.Marshal(ld)
}
