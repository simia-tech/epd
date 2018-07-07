package epd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd"
	"github.com/simia-tech/epd/pb"
	"github.com/simia-tech/epd/resolver"
)

func TestUnlockedFacadeAddSection(t *testing.T) {
	document, _ := generateDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")

	require.NotNil(t, document.Sections[sectionID])
	assert.Equal(t, "Test", document.Sections[sectionID].Name)
}

func TestUnlockedFacadeAddMember(t *testing.T) {
	foreignDocument, _ := generateDocument(t)
	contactID := foreignDocument.Id

	document, _ := generateDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")
	facade.AddMember(sectionID, contactID)

	require.NotNil(t, document.Contacts[contactID])
	assert.NotNil(t, document.Contacts[contactID].SectionKeys[sectionID])
}

func TestUnlockedFacadeLock(t *testing.T) {
	foreignDocument, _ := generateDocument(t)
	memoryResolver := &resolver.Memory{Documents: []*pb.UnlockedDocument{foreignDocument}}
	contactID := foreignDocument.Id

	document, privateKey := generateDocument(t)
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
