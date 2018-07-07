package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/simia-tech/errx"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func Decrypt(ciphertext []byte, privateKeyBytes PrivateKey) ([]byte, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, errx.Annotatef(err, "parse pkcs1 private key")
	}

	key, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext[:64])
	if err != nil {
		return nil, errx.Annotatef(err, "key decrypt pkcs1v15")
	}

	data, err := symmetric.Decrypt(ciphertext[64:], key)
	if err != nil {
		return nil, errx.Annotatef(err, "data decrypt")
	}

	return data, nil
}
