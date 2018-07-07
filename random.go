package epd

import (
	"crypto/rand"

	"github.com/lytics/base62"
)

func randomID() string {
	id := make([]byte, 16)
	rand.Read(id)
	return base62.StdEncoding.EncodeToString(id)[:16]
}
