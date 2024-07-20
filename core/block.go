package core

import (
	"Starcoin/crypto"
	"Starcoin/types"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

type Header struct {
	Version      uint32
	DataHash     types.Hash
	PreBlockHash types.Hash
	Timestamp    uint64
	Height       uint32
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	// Cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

func (b *Block) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(b.HeaderByte())
	if err != nil {
		return err
	}

	b.Validator = privateKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no sign")
	}

	if !b.Signature.Verify(b.Validator, b.HeaderByte()) {
		return fmt.Errorf("block has invalid signature")
	}

	return nil
}
func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) HeaderByte() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)
	return buf.Bytes()
}
