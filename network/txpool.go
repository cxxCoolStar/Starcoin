package network

import (
	"Starcoin/core"
	"Starcoin/types"
	"sync"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
	lock         sync.RWMutex
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Add(tx *core.Transaction) error {
	hasher := core.TxHasher{}
	hash := tx.Hash(hasher)

	p.lock.Lock()
	defer p.lock.Unlock()

	if p.Has(hash) {
		return nil
	}

	p.transactions[hash] = tx
	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	_, ok := p.transactions[hash]
	return ok
}

func (p *TxPool) Len() int {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.transactions = make(map[types.Hash]*core.Transaction)
}
