package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/simia-tech/errx"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func Encrypt(plaintext []byte, publicKeyBytes PublicKey) ([]byte, error) {
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		return nil, errx.Annotatef(err, "parse pkcs1 public key")
	}

	key := symmetric.GenerateKey(32)

	encryptedKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, key)
	if err != nil {
		return nil, errx.Annotatef(err, "key encrypt pkcs1v15")
	}

	encryptedData, err := symmetric.Encrypt(plaintext, key)
	if err != nil {
		return nil, errx.Annotatef(err, "data encrypt")
	}

	return append(encryptedKey, encryptedData...), nil
}
