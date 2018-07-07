package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/simia-tech/errx"
)

type PublicKey []byte
type PrivateKey []byte

func GenerateKeyPair() (PublicKey, PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, nil, errx.Annotatef(err, "generate rsa key pair")
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

	return publicKeyBytes, privateKeyBytes, nil
}
