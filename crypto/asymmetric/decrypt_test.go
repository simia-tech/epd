package asymmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/asymmetric"
)

func TestDecrypt(t *testing.T) {
	publicKey, privateKey := generateKeyPair(t)

	ciphertext, err := asymmetric.Encrypt([]byte("plaintext"), publicKey)
	require.NoError(t, err)

	plaintext, err := asymmetric.Decrypt(ciphertext, privateKey)
	require.NoError(t, err)

	assert.Equal(t, "plaintext", string(plaintext))
}
