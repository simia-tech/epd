package epd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd"
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/pb"
)

func TestLockedFacadeVerify(t *testing.T) {
	document, _ := makeLockedDocument(t)
	facade := epd.NewLockedFacade(document)

	assert.NoError(t, facade.Verify())
}

func TestLockedFacadeUnlock(t *testing.T) {
	document, privateKey := makeLockedDocument(t)
	facade := epd.NewLockedFacade(document)

	unlockedDocument, err := facade.Unlock(privateKey)
	require.NoError(t, err)

	assert.Equal(t, document.Id, unlockedDocument.Id)
	assert.Equal(t, document.PublicKey, unlockedDocument.PublicKey)
	assert.Equal(t, document.Name, unlockedDocument.Name)
	assert.Len(t, unlockedDocument.Sections, 1)
}

func makeLockedDocument(tb testing.TB) (*pb.LockedDocument, asymmetric.PrivateKey) {
	foreignDocument, _, resolver := makeForeignDocument(tb)
	document, privateKey := makeDocument(tb)
	facade := epd.NewUnlockedFacade(document)
	addTestSectionAndMember(facade, foreignDocument.Id)

	lockedDocument, err := facade.Lock(privateKey, resolver)
	require.NoError(tb, err)

	return lockedDocument, privateKey
}
