package asymmetric

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	"github.com/simia-tech/errx"
)

func Verify(message []byte, publicKeyBytes PublicKey, signature Signature) error {
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		return errx.Annotatef(err, "parse pkcs1 public key")
	}

	hashed := sha256.Sum256(message)

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}
