package network

import (
	"Starcoin/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTxPool_New(t *testing.T) {
	p := NewTxPool()

	assert.Equal(t, p.Len(), 0)
}

func TestTxPool_Add(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}
