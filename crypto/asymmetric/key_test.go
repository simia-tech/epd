package asymmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/asymmetric"
)

func TestGenerateKeyPair(t *testing.T) {
	publicKey, privateKey := generateKeyPair(t)
	assert.NotEmpty(t, publicKey)
	assert.NotEmpty(t, privateKey)
}

func generateKeyPair(tb testing.TB) (asymmetric.PublicKey, asymmetric.PrivateKey) {
	publicKey, privateKey, err := asymmetric.GenerateKeyPair()
	require.NoError(tb, err)
	return publicKey, privateKey
}
