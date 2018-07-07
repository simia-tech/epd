package resolver

import "github.com/simia-tech/epd/crypto/asymmetric"

type Interface interface {
	ResolvePublicKey(string) (asymmetric.PublicKey, error)
}
