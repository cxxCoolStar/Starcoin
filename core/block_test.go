package core

import (
	"Starcoin/types"
	"fmt"
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

func TestHashBlock(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}
