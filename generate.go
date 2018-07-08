package epd

import (
	"github.com/simia-tech/errx"

	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/pb"
)

func Generate(name string) (*pb.UnlockedDocument, asymmetric.PrivateKey, error) {
	id := randomID()

	publicKey, privateKey, err := asymmetric.GenerateKeyPair()
	if err != nil {
		return nil, nil, errx.Annotatef(err, "generate key pair")
	}

	return &pb.UnlockedDocument{
		Id:        id,
		PublicKey: publicKey,
		Name:      name,
		Sections:  map[string]*pb.UnlockedSection{},
	}, privateKey, nil
}
