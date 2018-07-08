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

func TestUnlockedFacadeAddSection(t *testing.T) {
	document, _ := makeDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")

	require.NotNil(t, document.Sections[sectionID])
	assert.Equal(t, "Test", document.Sections[sectionID].Name)
}

func TestUnlockedFacadeAddMember(t *testing.T) {
	foreignDocument, _ := makeDocument(t)
	contactID := foreignDocument.Id

	document, _ := makeDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")
	facade.AddMember(sectionID, contactID)

	require.NotNil(t, document.Contacts[contactID])
	assert.NotNil(t, document.Contacts[contactID].SectionKeys[sectionID])
}

func TestUnlockedFacadeLock(t *testing.T) {
	foreignDocument, _ := makeDocument(t)
	memoryResolver := &resolver.Memory{Documents: []*pb.UnlockedDocument{foreignDocument}}
	contactID := foreignDocument.Id

	document, privateKey := makeDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")
	facade.AddMember(sectionID, contactID)

	lockedDocument, err := facade.Lock(privateKey, memoryResolver)
	require.NoError(t, err)

	assert.Equal(t, document.Id, lockedDocument.Id)
	assert.Equal(t, document.PublicKey, lockedDocument.PublicKey)
	assert.Len(t, lockedDocument.Contacts, 1)
	assert.Len(t, lockedDocument.Sections, 1)
	assert.NotEmpty(t, lockedDocument.Signature)
}

func TestLockedFacadeVerify(t *testing.T) {
	document, _ := makeLockedDocument(t)
	facade := epd.NewLockedFacade(document)
	assert.NoError(t, facade.Verify())
}

func makeLockedDocument(tb testing.TB) (*pb.LockedDocument, asymmetric.PrivateKey) {
	foreignDocument, _ := makeDocument(tb)
	memoryResolver := &resolver.Memory{Documents: []*pb.UnlockedDocument{foreignDocument}}
	contactID := foreignDocument.Id

	document, privateKey := makeDocument(tb)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")
	facade.AddMember(sectionID, contactID)

	lockedDocument, err := facade.Lock(privateKey, memoryResolver)
	require.NoError(tb, err)

	return lockedDocument, privateKey
}
