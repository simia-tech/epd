package stringset_test

import (
	"testing"

	"github.com/simia-tech/epd/util/stringset"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.True(t, stringset.Contains([]string{"a", "b"}, "a"))
	assert.False(t, stringset.Contains([]string{"a", "b"}, "c"))
}

func TestWith(t *testing.T) {
	result := stringset.With([]string{"a", "b"}, "c")
	assert.Equal(t, result, []string{"a", "b", "c"})
}
