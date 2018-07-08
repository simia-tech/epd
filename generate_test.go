package epd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd"
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/pb"
)

func TestGenerate(t *testing.T) {
	document, privateKey := makeDocument(t)

	assert.NotEmpty(t, document.Id)
	assert.NotEmpty(t, document.PublicKey)
	assert.Empty(t, document.Contacts)
	assert.Empty(t, document.Sections)

	assert.NotEmpty(t, privateKey)
}

func makeDocument(tb testing.TB) (*pb.UnlockedDocument, asymmetric.PrivateKey) {
	document, privateKey, err := epd.Generate()
	require.NoError(tb, err)
	return document, privateKey
}
