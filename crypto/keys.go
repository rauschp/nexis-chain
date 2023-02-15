package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

type PrivateKey struct {
	Key ed25519.PrivateKey
}

func GenerateNewPrivateKey() *PrivateKey {
	_, priv, _ := ed25519.GenerateKey(rand.Reader)

	return &PrivateKey{
		Key: priv,
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.Key
}

func (p *PrivateKey) Sign(message []byte) []byte {
	return ed25519.Sign(p.Key, message)
}

// Public Key
type PublicKey struct {
	Key ed25519.PublicKey
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, 32)
	copy(b, p.Key[32:])

	return &PublicKey{
		Key: b,
	}
}

func (p *PublicKey) Verify(message []byte, signature []byte) bool {
	return ed25519.Verify(p.Key, message, signature)
}
