package symmetric

import "crypto/rand"

type Key []byte

func GenerateKey(size int) Key {
	key := make([]byte, size)
	rand.Read(key)
	return key
}
