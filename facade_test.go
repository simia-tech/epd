package epd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/epd"
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
