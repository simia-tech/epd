package symmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func TestEncrypt(t *testing.T) {
	key := makeKey()

	ciphertext, err := symmetric.Encrypt([]byte("plaintext"), key)
	require.NoError(t, err)

	assert.Len(t, ciphertext, 37)
}
