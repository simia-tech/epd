package symmetric_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func TestGenerateKey(t *testing.T) {
	key := makeKey()
	assert.Len(t, key, 32)
}

func TestKeyJSONEncoding(t *testing.T) {
	key := makeKey()
	assert.Equal(t, key, *(decodeJSON(t, encodeJSON(t, key), &symmetric.Key{})).(*symmetric.Key))
}

func makeKey() symmetric.Key {
	return symmetric.GenerateKey()
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
