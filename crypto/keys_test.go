package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	privKey := GenerateNewPrivateKey()
	pubKey := privKey.Public()

	msg := []byte("Test message")
	sig := privKey.Sign(msg)

	assert.True(t, pubKey.Verify(msg, sig))
}

func TestVerifyWrongKey(t *testing.T) {
	privKey1 := GenerateNewPrivateKey()
	privKey2 := GenerateNewPrivateKey()

	msg := []byte("fefefe")
	sig := privKey1.Sign(msg)

	assert.True(t, privKey1.Public().Verify(msg, sig))
	assert.False(t, privKey2.Public().Verify(msg, sig))
}
