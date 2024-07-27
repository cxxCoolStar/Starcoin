package network

import (
	"Starcoin/core"
	"github.com/stretchr/testify/assert"
	"strconv"
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

func TestSortTransactions(t *testing.T) {
	p := NewTxPool()
	txLen := 1000

	for i := 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte(strconv.FormatInt(int64(i), 10)))
		tx.SetFirstSeen(int64(i))
		assert.Nil(t, p.Add(tx))
	}

	assert.Equal(t, txLen, p.Len())

	txx := p.Transaction()
	for i := 0; i < len(txx)-1; i++ {
		assert.True(t, txx[i].FirstSeen() <= txx[i+1].FirstSeen())
	}
}
