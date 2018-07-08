package symmetric_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func TestGenerateKey(t *testing.T) {
	key := symmetric.GenerateKey()
	assert.Len(t, key, 32)
}

func makeKey() symmetric.Key {
	return symmetric.GenerateKey()
}
