package pb

import "github.com/golang/protobuf/proto"

func (uc *UnlockedContact) MarshalBinary() ([]byte, error) {
	return proto.Marshal(uc)
}
