package resolver

import (
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/pb"
)

type Memory struct {
	Documents []*pb.UnlockedDocument
}

func (m *Memory) ResolvePublicKey(contactID string) (asymmetric.PublicKey, error) {
	for _, document := range m.Documents {
		if document.Id == contactID {
			return document.PublicKey, nil
		}
	}
	return nil, nil
}
