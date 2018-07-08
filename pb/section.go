package pb

import "github.com/golang/protobuf/proto"

func (us *UnlockedSection) MarshalBinary() ([]byte, error) {
	return proto.Marshal(us)
}

func (us *UnlockedSection) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, us)
}
