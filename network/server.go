package network

import (
	"Starcoin/core"
	"Starcoin/crypto"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type ServerOpts struct {
	Transports []Transport
	BlockTime  time.Duration
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts  ServerOpts
	blockTime   time.Duration
	memPool     *TxPool
	rpcCh       chan RPC
	quitCh      chan struct{}
	isValidator bool
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts:  opts,
		blockTime:   opts.BlockTime,
		memPool:     NewTxPool(),
		isValidator: opts.PrivateKey != nil,
		rpcCh:       make(chan RPC, 1024),
		quitCh:      make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * s.blockTime)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("Received RPC from %s: %s\n", rpc.From, rpc.Payload)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			if s.isValidator {
				s.createNewBlock()
			}
		}
	}

	fmt.Printf("Server stopped\n")
}

func (s *Server) initTransports() {
	for _, tr := range s.ServerOpts.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}

}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": tx.Hash(core.TxHasher{}),
		}).Info("mempool already contaions tx with hash")
		return nil
	}
	logrus.WithFields(logrus.Fields{
		"hash": tx.Hash(core.TxHasher{}),
	}).Info("adding new tx to the mempool")
	return s.memPool.Add(tx)
}

func (s *Server) createNewBlock() error {
	fmt.Println("creating a new block")
	return nil
}
