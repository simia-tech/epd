package epd_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/simia-tech/epd"
)

func TestRandomDocumentID(t *testing.T) {
	id := epd.RandomDocumentID()
	assert.False(t, strings.HasSuffix(id, "="))
	assert.Regexp(t, "^[a-z0-9]{16}$", id)
}

func TestRandomSectionID(t *testing.T) {
	id := epd.RandomSectionID()
	assert.False(t, strings.HasSuffix(id, "="))
	assert.Regexp(t, "^[a-z0-9]{16}$", id)
}
