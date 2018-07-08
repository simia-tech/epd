package epd

import (
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/crypto/symmetric"
	"github.com/simia-tech/epd/pb"
	"github.com/simia-tech/errx"
)

type LockedFacade struct {
	document *pb.LockedDocument
}

func NewLockedFacade(document *pb.LockedDocument) *LockedFacade {
	return &LockedFacade{document: document}
}

func (lf *LockedFacade) Verify() error {
	signature := lf.document.Signature
	lf.document.Signature = nil

	documentBytes, err := lf.document.MarshalBinary()
	lf.document.Signature = signature
	if err != nil {
		return errx.Annotatef(err, "marshal locked document")
	}

	return asymmetric.Verify(documentBytes, lf.document.PublicKey, signature)
}

func (lf *LockedFacade) Unlock(privateKey asymmetric.PrivateKey) (*pb.UnlockedDocument, error) {
	document := &pb.UnlockedDocument{
		Id:        lf.document.Id,
		PublicKey: lf.document.PublicKey,
		Name:      lf.document.Name,
		Sections:  map[string]*pb.UnlockedSection{},
	}

	ownContact, err := unlockContact(lf.document.Contacts[lf.document.Id], privateKey)
	if err != nil {
		return nil, errx.Annotatef(err, "own contact")
	}

	for sectionID, lockedSection := range lf.document.Sections {
		section, err := unlockSection(lockedSection, ownContact.SectionKeys[sectionID])
		if err != nil {
			return nil, errx.Annotatef(err, "section %s", sectionID)
		}

		document.Sections[sectionID] = section
	}

	return document, nil
}

func unlockSection(section *pb.LockedSection, key symmetric.Key) (*pb.UnlockedSection, error) {
	bytes, err := symmetric.Decrypt(section.Encrypted, key)
	if err != nil {
		return nil, errx.Annotatef(err, "decrypt")
	}

	unlockedSection := &pb.UnlockedSection{}
	if err := unlockedSection.UnmarshalBinary(bytes); err != nil {
		return nil, errx.Annotatef(err, "unmarshal binary")
	}
	return unlockedSection, nil
}

func unlockContact(contact *pb.LockedContact, privateKey asymmetric.PrivateKey) (*pb.UnlockedContact, error) {
	bytes, err := asymmetric.Decrypt(contact.Encrypted, privateKey)
	if err != nil {
		return nil, errx.Annotatef(err, "decrypt")
	}

	unlockedContact := &pb.UnlockedContact{}
	if err := unlockedContact.UnmarshalBinary(bytes); err != nil {
		return nil, errx.Annotatef(err, "unmarshal binary")
	}
	return unlockedContact, nil
}
