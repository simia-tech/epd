package symmetric

import "crypto/rand"

type Key []byte

func GenerateKey() Key {
	key := make([]byte, 32)
	rand.Read(key)
	return key
}
