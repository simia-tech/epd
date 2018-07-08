package symmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func TestDecrypt(t *testing.T) {
	key := makeKey()

	ciphertext, err := symmetric.Encrypt([]byte("plaintext"), key)
	require.NoError(t, err)

	plaintext, err := symmetric.Decrypt(ciphertext, key)
	require.NoError(t, err)

	assert.Equal(t, "plaintext", string(plaintext))
}
