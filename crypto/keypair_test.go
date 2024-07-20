package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	assert.NotNil(t, publicKey.Address())
}
