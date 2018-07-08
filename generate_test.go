package epd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd"
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/pb"
	"github.com/simia-tech/epd/resolver"
)

func TestGenerate(t *testing.T) {
	document, privateKey := makeDocument(t)

	assert.NotEmpty(t, document.Id)
	assert.NotEmpty(t, document.PublicKey)
	assert.Equal(t, "Test", document.Name)
	assert.Empty(t, document.Sections)

	assert.NotEmpty(t, privateKey)
}

func makeDocument(tb testing.TB) (*pb.UnlockedDocument, asymmetric.PrivateKey) {
	document, privateKey, err := epd.Generate("Test")
	require.NoError(tb, err)
	return document, privateKey
}

func makeForeignDocument(tb testing.TB) (*pb.UnlockedDocument, asymmetric.PrivateKey, *resolver.Memory) {
	document, privateKey := makeDocument(tb)
	resolver := &resolver.Memory{Documents: []*pb.UnlockedDocument{document}}
	return document, privateKey, resolver
}
