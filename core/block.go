package core

import (
	"Starcoin/crypto"
	"Starcoin/types"
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

func (b *Block) Sign(privateKey crypto.PrivateKey) *crypto.Signature {
	//sig, err := privateKey.Sign(b.)
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

func (b *Block) hashableData() {

}
