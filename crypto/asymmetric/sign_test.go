package asymmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/asymmetric"
)

func TestSign(t *testing.T) {
	_, privateKey := makeKeyPair(t)

	signature, err := asymmetric.Sign([]byte("ciphertext"), privateKey)
	require.NoError(t, err)

	assert.Len(t, signature, 64)
}
