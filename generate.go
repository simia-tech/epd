package epd

import (
	"crypto/rand"

	"github.com/lytics/base62"
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

func randomID() string {
	id := make([]byte, 16)
	rand.Read(id)
	return base62.StdEncoding.EncodeToString(id)[:16]
}
