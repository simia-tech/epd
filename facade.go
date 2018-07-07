package epd

import (
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/crypto/symmetric"
	"github.com/simia-tech/epd/pb"
	"github.com/simia-tech/epd/resolver"
	"github.com/simia-tech/errx"
)

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

func (uf *UnlockedFacade) Lock(privateKey asymmetric.PrivateKey, resolver resolver.Interface) (*pb.LockedDocument, error) {
	document := &pb.LockedDocument{
		Id:        uf.document.Id,
		PublicKey: uf.document.PublicKey,
		Contacts:  map[string]*pb.LockedContact{},
		Sections:  map[string]*pb.LockedSection{},
	}

	sectionKeys := map[string]symmetric.Key{}

	for sectionID, unlockedSection := range uf.document.Sections {
		sectionKeys[sectionID] = symmetric.GenerateKey(32)

		sectionBytes, err := unlockedSection.MarshalBinary()
		if err != nil {
			return nil, errx.Annotatef(err, "marshal section %s", sectionID)
		}

		encryptedBytes, err := symmetric.Encrypt(sectionBytes, sectionKeys[sectionID])
		if err != nil {
			return nil, errx.Annotatef(err, "encrypt section %s", sectionID)
		}

		document.Sections[sectionID] = &pb.LockedSection{Encrypted: encryptedBytes}
	}

	for contactID, unlockedContact := range uf.document.Contacts {
		for sectionID := range unlockedContact.SectionKeys {
			unlockedContact.SectionKeys[sectionID] = sectionKeys[sectionID]
		}

		foreignPublicKey, err := resolver.ResolvePublicKey(contactID)
		if err != nil {
			return nil, errx.Annotatef(err, "resolve public key for contact %s", contactID)
		}
		if foreignPublicKey == nil {
			return nil, errx.NotFoundf("could not resolve public key for %s", contactID)
		}

		contactBytes, err := unlockedContact.MarshalBinary()
		if err != nil {
			return nil, errx.Annotatef(err, "marshal contact %s", contactID)
		}

		encryptedBytes, err := asymmetric.Encrypt(contactBytes, foreignPublicKey)
		if err != nil {
			return nil, errx.Annotatef(err, "encrypt contact %s", contactID)
		}

		document.Contacts[contactID] = &pb.LockedContact{Encrypted: encryptedBytes}
	}

	documentBytes, err := document.MarshalBinary()
	if err != nil {
		return nil, errx.Annotatef(err, "marshal locked document")
	}

	signature, err := asymmetric.Sign(documentBytes, privateKey)
	if err != nil {
		return nil, errx.Annotatef(err, "sign locked document")
	}
	document.Signature = signature

	return document, nil
}
