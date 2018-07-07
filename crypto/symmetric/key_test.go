package symmetric_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/simia-tech/epd/crypto/symmetric"
)

func TestGenerateKey(t *testing.T) {
	sizes := []int{0, 1, 2, 3, 9, 10, 11}
	for _, size := range sizes {
		t.Run(fmt.Sprintf("Len%d", size), func(t *testing.T) {
			key := symmetric.GenerateKey(size)
			assert.Len(t, key, size)
		})
	}
}

func generateKey() symmetric.Key {
	return symmetric.GenerateKey(32)
}
