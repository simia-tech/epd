package pb

import "github.com/golang/protobuf/proto"

func (uc *UnlockedContact) MarshalBinary() ([]byte, error) {
	return proto.Marshal(uc)
}

func (uc *UnlockedContact) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, uc)
}
