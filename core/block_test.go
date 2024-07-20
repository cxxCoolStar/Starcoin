package core

import (
	"Starcoin/crypto"
	"Starcoin/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:      1,
		PreBlockHash: types.RandomHash(),
		Timestamp:    uint64(time.Now().UnixNano()),
		Height:       height,
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privateKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privateKey))
	assert.Nil(t, b.Verify())
	otherPrivateKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivateKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())

}
