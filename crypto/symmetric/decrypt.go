package symmetric

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/simia-tech/errx"
)

func Decrypt(ciphertext []byte, key Key) ([]byte, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return nil, errx.Annotatef(err, "new aes cipher")
	}

	aesGCM, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, errx.Annotatef(err, "new cipher gcm")
	}

	nonceSize := aesGCM.NonceSize()
	return aesGCM.Open(nil, ciphertext[:nonceSize], ciphertext[nonceSize:], nil)
}
