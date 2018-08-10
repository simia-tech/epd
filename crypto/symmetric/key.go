package symmetric

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
)

type Key []byte

var encoding = base64.RawStdEncoding.Strict()

func GenerateKey() Key {
	key := make([]byte, 32)
	rand.Read(key)
	return key
}

func (k Key) MarshalJSON() ([]byte, error) {
	buffer := make([]byte, encoding.EncodedLen(len(k)))
	encoding.Encode(buffer, k)
	return bytes.Join([][]byte{[]byte("\""), buffer, []byte("\"")}, []byte{}), nil
}

func (k *Key) UnmarshalJSON(data []byte) error {
	data = data[1 : len(data)-1]
	*k = make([]byte, encoding.DecodedLen(len(data)))
	_, err := encoding.Decode(*k, data)
	return err
}
