package epd

import "github.com/simia-tech/epd/pb"

type UnlockedFacade struct {
	document *pb.UnlockedDocument
}

func NewUnlockedFacade(document *pb.UnlockedDocument) *UnlockedFacade {
	return &UnlockedFacade{document: document}
}

func (uf *UnlockedFacade) AddSection(name string) string {
	sectionID := randomID()
	uf.document.Sections[sectionID] = &pb.UnlockedSection{
		Name: name,
	}
	return sectionID
}

func (uf *UnlockedFacade) AddMember(sectionID, contactID string) {
	if contact, ok := uf.document.Contacts[sectionID]; ok {
		contact.SectionKeys[sectionID] = []byte{}
	} else {
		uf.document.Contacts[contactID] = &pb.UnlockedContact{
			SectionKeys: map[string][]byte{sectionID: []byte{}},
		}
	}
}
