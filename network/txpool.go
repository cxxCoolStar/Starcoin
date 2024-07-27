package network

import (
	"Starcoin/core"
	"Starcoin/types"
	"fmt"
	"sort"
	"sync"
)

type TxMapSorter struct {
	transactions []*core.Transaction
}

func (t TxMapSorter) Len() int {
	return len(t.transactions)
}

func (t TxMapSorter) Less(i, j int) bool {
	return t.transactions[i].FirstSeen() < t.transactions[j].FirstSeen()
}

func (t TxMapSorter) Swap(i, j int) {
	t.transactions[i], t.transactions[j] = t.transactions[j], t.transactions[i]
}

func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	transactions := make([]*core.Transaction, len(txMap))

	i := 0
	for _, val := range txMap {
		transactions[i] = val
		i++
	}

	s := &TxMapSorter{transactions}
	sort.Sort(s)

	return s
}

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
	lock         sync.RWMutex
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Transaction() []*core.Transaction {
	s := NewTxMapSorter(p.transactions)
	return s.transactions
}

func (p *TxPool) Add(tx *core.Transaction) error {
	hasher := core.TxHasher{}
	hash := tx.Hash(hasher)

	fmt.Printf("Adding transaction with hash: %x\n", hash)

	// 先检查是否已经存在
	if p.Has(hash) {
		return nil
	}

	// 获取写锁并添加交易
	p.lock.Lock()
	defer p.lock.Unlock()

	// 再次检查是否已经存在 (在获取写锁后进行双重检查)
	if _, ok := p.transactions[hash]; ok {
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
