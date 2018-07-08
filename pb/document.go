package pb

import "github.com/golang/protobuf/proto"

func (ud *UnlockedDocument) MarshalBinary() ([]byte, error) {
	return proto.Marshal(ud)
}

func (ud *UnlockedDocument) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, ud)
}

func (ld *LockedDocument) MarshalBinary() ([]byte, error) {
	return proto.Marshal(ld)
}

func (ld *LockedDocument) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, ld)
}
