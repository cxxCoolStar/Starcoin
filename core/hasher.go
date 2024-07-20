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

func (BlockHasher) Hash(b *Block) types.Hash {
	h := sha256.Sum256(b.HeaderByte())
	return h
}
