package epd

import (
	"github.com/simia-tech/errx"

	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/pb"
)

func Generate() (*pb.UnlockedDocument, asymmetric.PrivateKey, error) {
	id := randomID()

	publicKey, privateKey, err := asymmetric.GenerateKeyPair()
	if err != nil {
		return nil, nil, errx.Annotatef(err, "generate key pair")
	}

	return &pb.UnlockedDocument{
		Id:        id,
		PublicKey: publicKey,
		Contacts:  map[string]*pb.UnlockedContact{},
		Sections:  map[string]*pb.UnlockedSection{},
	}, privateKey, nil
}
