package epd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd"
)

func TestUnlockedFacadeAddSection(t *testing.T) {
	document, _ := makeDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")

	require.NotNil(t, document.Sections[sectionID])
	assert.Equal(t, "Test", document.Sections[sectionID].Name)
	assert.Len(t, document.Sections[sectionID].ContentKey, 32)
}

func TestUnlockedFacadeAddMember(t *testing.T) {
	foreignDocument, _ := makeDocument(t)
	contactID := foreignDocument.Id

	document, _ := makeDocument(t)
	facade := epd.NewUnlockedFacade(document)

	sectionID := facade.AddSection("Test")
	require.NoError(t, facade.AddMember(sectionID, contactID))

	assert.Contains(t, document.Sections[sectionID].MemberIds, contactID)
}

func TestUnlockedFacadeLock(t *testing.T) {
	foreignDocument, _, resolver := makeForeignDocument(t)
	document, privateKey := makeDocument(t)
	facade := epd.NewUnlockedFacade(document)
	addTestSectionAndMember(facade, foreignDocument.Id)

	lockedDocument, err := facade.Lock(privateKey, resolver)
	require.NoError(t, err)

	assert.Equal(t, document.Id, lockedDocument.Id)
	assert.Equal(t, document.PublicKey, lockedDocument.PublicKey)
	assert.Equal(t, document.Name, lockedDocument.Name)
	assert.Len(t, lockedDocument.Contacts, 2)
	assert.Len(t, lockedDocument.Sections, 1)
	assert.NotEmpty(t, lockedDocument.Signature)
}

func addTestSectionAndMember(facade *epd.UnlockedFacade, contactID string) string {
	sectionID := facade.AddSection("Test")
	facade.AddMember(sectionID, contactID)
	return sectionID
}
