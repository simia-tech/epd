package epd

import (
	"github.com/simia-tech/epd/crypto/asymmetric"
	"github.com/simia-tech/epd/crypto/symmetric"
	"github.com/simia-tech/epd/pb"
	"github.com/simia-tech/epd/resolver"
	"github.com/simia-tech/epd/util/stringset"
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
		Name:       name,
		ContentKey: symmetric.GenerateKey(),
	}
	return sectionID
}

func (uf *UnlockedFacade) AddMember(sectionID, contactID string) error {
	section, ok := uf.document.Sections[sectionID]
	if !ok {
		return errx.NotFoundf("could not find section %s", sectionID)
	}

	section.MemberIds = stringset.With(section.MemberIds, contactID)

	return nil
}

func (uf *UnlockedFacade) Lock(privateKey asymmetric.PrivateKey, resolver resolver.Interface) (*pb.LockedDocument, error) {
	document := &pb.LockedDocument{
		Id:        uf.document.Id,
		PublicKey: uf.document.PublicKey,
		Name:      uf.document.Name,
		Contacts:  map[string]*pb.LockedContact{},
		Sections:  map[string]*pb.LockedSection{},
	}

	contacts := map[string]*pb.UnlockedContact{}
	ownContact := &pb.UnlockedContact{SectionKeys: map[string][]byte{}}

	for sectionID, unlockedSection := range uf.document.Sections {
		sectionKey := symmetric.GenerateKey()
		ownContact.SectionKeys[sectionID] = sectionKey

		for _, memberID := range unlockedSection.MemberIds {
			if unlockedContact, ok := contacts[memberID]; ok {
				unlockedContact.SectionKeys[sectionID] = sectionKey
			} else {
				contacts[memberID] = &pb.UnlockedContact{
					SectionKeys: map[string][]byte{sectionID: sectionKey},
				}
			}
		}

		lockedSection, err := lockSection(unlockedSection, sectionKey)
		if err != nil {
			return nil, errx.Annotatef(err, "section %s", sectionID)
		}
		document.Sections[sectionID] = lockedSection
	}

	for contactID, unlockedContact := range contacts {
		foreignPublicKey, err := resolver.ResolvePublicKey(contactID)
		if err != nil {
			return nil, errx.Annotatef(err, "resolve public key for contact %s", contactID)
		}
		if foreignPublicKey == nil {
			return nil, errx.NotFoundf("could not resolve public key for %s", contactID)
		}

		lockedContact, err := lockContact(unlockedContact, foreignPublicKey)
		if err != nil {
			return nil, errx.Annotatef(err, "contact %s", contactID)
		}
		document.Contacts[contactID] = lockedContact
	}

	lockedOwnContact, err := lockContact(ownContact, uf.document.PublicKey)
	if err != nil {
		return nil, errx.Annotatef(err, "own contact")
	}
	document.Contacts[uf.document.Id] = lockedOwnContact

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

func lockSection(section *pb.UnlockedSection, key symmetric.Key) (*pb.LockedSection, error) {
	bytes, err := section.MarshalBinary()
	if err != nil {
		return nil, errx.Annotatef(err, "marshal binary")
	}

	encryptedBytes, err := symmetric.Encrypt(bytes, key)
	if err != nil {
		return nil, errx.Annotatef(err, "encrypt")
	}

	return &pb.LockedSection{Encrypted: encryptedBytes}, nil
}

func lockContact(contact *pb.UnlockedContact, publicKey asymmetric.PublicKey) (*pb.LockedContact, error) {
	bytes, err := contact.MarshalBinary()
	if err != nil {
		return nil, errx.Annotatef(err, "marshal binary")
	}

	encryptedBytes, err := asymmetric.Encrypt(bytes, publicKey)
	if err != nil {
		return nil, errx.Annotatef(err, "encrypt")
	}

	return &pb.LockedContact{Encrypted: encryptedBytes}, nil
}
