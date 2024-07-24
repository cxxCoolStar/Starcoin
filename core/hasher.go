package core

import (
	"Starcoin/types"
	"crypto/sha256"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(h *Header) types.Hash {
	return sha256.Sum256(h.Bytes())
}

type TxHasher struct {
}

func (TxHasher) Hash(tx *Transaction) types.Hash {
	return sha256.Sum256(tx.Data)
}
