package asymmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/asymmetric"
)

func TestVerify(t *testing.T) {
	publicKey, privateKey := generateKeyPair(t)

	signature, err := asymmetric.Sign([]byte("ciphertext"), privateKey)
	require.NoError(t, err)

	err = asymmetric.Verify([]byte("ciphertext"), publicKey, signature)
	assert.NoError(t, err)
}
