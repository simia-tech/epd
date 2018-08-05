package asymmetric_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/asymmetric"
)

func TestGenerateKeyPair(t *testing.T) {
	publicKey, privateKey := makeKeyPair(t)
	assert.NotEmpty(t, publicKey)
	assert.NotEmpty(t, privateKey)
}

func TestKeyJSONEncoding(t *testing.T) {
	publicKey, privateKey := makeKeyPair(t)
	assert.Equal(t, publicKey, *(decodeJSON(t, encodeJSON(t, publicKey), &asymmetric.PublicKey{})).(*asymmetric.PublicKey))
	assert.Equal(t, privateKey, *(decodeJSON(t, encodeJSON(t, privateKey), &asymmetric.PrivateKey{})).(*asymmetric.PrivateKey))
}

func makeKeyPair(tb testing.TB) (asymmetric.PublicKey, asymmetric.PrivateKey) {
	publicKey, privateKey, err := asymmetric.GenerateKeyPair()
	require.NoError(tb, err)
	return publicKey, privateKey
}

func encodeJSON(tb testing.TB, value interface{}) []byte {
	data, err := json.Marshal(value)
	require.NoError(tb, err)
	return data
}

func decodeJSON(tb testing.TB, data []byte, value interface{}) interface{} {
	require.NoError(tb, json.Unmarshal(data, value))
	return value
}
