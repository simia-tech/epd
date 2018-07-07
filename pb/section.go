package pb

import "github.com/golang/protobuf/proto"

func (us *UnlockedSection) MarshalBinary() ([]byte, error) {
	return proto.Marshal(us)
}
