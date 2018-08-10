package asymmetric

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"github.com/simia-tech/errx"
)

type PublicKey []byte
type PrivateKey []byte

var encoding = base64.RawStdEncoding.Strict()

func GenerateKeyPair() (PublicKey, PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, nil, errx.Annotatef(err, "generate rsa key pair")
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

	return publicKeyBytes, privateKeyBytes, nil
}

func (pk PublicKey) MarshalJSON() ([]byte, error) {
	buffer := make([]byte, encoding.EncodedLen(len(pk)))
	encoding.Encode(buffer, pk)
	return bytes.Join([][]byte{[]byte("\""), buffer, []byte("\"")}, []byte{}), nil
}

func (pk *PublicKey) UnmarshalJSON(data []byte) error {
	data = data[1 : len(data)-1]
	*pk = make([]byte, encoding.DecodedLen(len(data)))
	_, err := encoding.Decode(*pk, data)
	return err
}

func (pk PrivateKey) MarshalJSON() ([]byte, error) {
	buffer := make([]byte, encoding.EncodedLen(len(pk)))
	encoding.Encode(buffer, pk)
	return bytes.Join([][]byte{[]byte("\""), buffer, []byte("\"")}, []byte{}), nil
}

func (pk *PrivateKey) UnmarshalJSON(data []byte) error {
	data = data[1 : len(data)-1]
	*pk = make([]byte, encoding.DecodedLen(len(data)))
	_, err := encoding.Decode(*pk, data)
	return err
}
