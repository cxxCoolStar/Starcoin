package core

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	store     Storage
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
	if height > bc.Height() {
		return nil, fmt.Errorf("height (%d) too high", height)
	}

	return bc.headers[height], nil
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) AddBlockWithoutValidation(b *Block) error {
	logrus.WithFields(logrus.Fields{}).Info("adding new block")
	bc.headers = append(bc.headers, b.Header)

	logrus.WithFields(logrus.Fields{}).Info("adding new block")
	return bc.store.Put(b)
}

func (bc *Blockchain) addGenesisBlock(b *Block) {

}
