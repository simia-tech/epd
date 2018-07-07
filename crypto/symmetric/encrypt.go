package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/simia-tech/errx"
)

func Encrypt(plaintext []byte, key Key) ([]byte, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return nil, errx.Annotatef(err, "new aes cipher")
	}

	aesGCM, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, errx.Annotatef(err, "new cipher gcm")
	}

	nonce := make([]byte, aesGCM.NonceSize())
	rand.Read(nonce)

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	return append(nonce, ciphertext...), nil
}
