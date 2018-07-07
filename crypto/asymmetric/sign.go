package asymmetric

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	"github.com/simia-tech/errx"
)

type Signature []byte

func Sign(message []byte, privateKeyBytes PrivateKey) (Signature, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, errx.Annotatef(err, "parse pkcs1 private key")
	}

	hashed := sha256.Sum256(message)

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
}
