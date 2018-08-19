package epd

import (
	"crypto/rand"
	"encoding/base32"
	"strings"
)

func RandomDocumentID() string {
	id := make([]byte, 10)
	rand.Read(id)
	return strings.ToLower(base32.StdEncoding.EncodeToString(id))
}

func RandomSectionID() string {
	id := make([]byte, 10)
	rand.Read(id)
	return strings.ToLower(base32.StdEncoding.EncodeToString(id))
}
