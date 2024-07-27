package core

import (
	"Starcoin/crypto"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privateKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privateKey))
	assert.Nil(t, tx.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivateKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func TestTxEncodeDecode(t *testing.T) {
	tx := randomTxWithSignature(t)
	buf := &bytes.Buffer{}

	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))

	txDecoded := new(Transaction)

	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))

	assert.Equal(t, tx, txDecoded)

}

func randomTxWithSignature(t *testing.T) *Transaction {
	privateKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("foo"),
	}
	assert.Nil(t, tx.Sign(privateKey))

	return &tx
}
