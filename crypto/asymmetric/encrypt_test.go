package asymmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/asymmetric"
)

func TestEncrypt(t *testing.T) {
	publicKey, _ := generateKeyPair(t)

	ciphertext, err := asymmetric.Encrypt([]byte("plaintext"), publicKey)
	require.NoError(t, err)

	assert.Len(t, ciphertext, 101)
}
