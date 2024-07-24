package core

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
)

type Blockchain struct {
	store     Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	if genesis == nil {
		return nil, errors.New("genesis block cannot be nil")
	}

	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)
	if err := bc.AddBlockWithoutValidation(genesis); err != nil {
		return nil, err
	}
	return bc, nil
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	err := bc.validator.ValidateBlock(b)
	if err != nil {
		return err
	}

	bc.AddBlockWithoutValidation(b)

	return nil
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	bc.lock.Lock()
	if height > bc.Height() {
		return nil, fmt.Errorf("height (%d) too high", height)
	}

	defer bc.lock.Unlock()

	return bc.headers[height], nil
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) AddBlockWithoutValidation(b *Block) error {
	bc.lock.Lock()
	bc.headers = append(bc.headers, b.Header)
	bc.lock.Unlock()

	logrus.WithFields(logrus.Fields{}).Info("adding new block")
	return bc.store.Put(b)
}

func (bc *Blockchain) addGenesisBlock(b *Block) {

}
