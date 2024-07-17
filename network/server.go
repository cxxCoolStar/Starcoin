package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts ServerOpts
	rpcCh      chan RPC
	quitCh     chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC, 1024),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("Received RPC from %s: %s\n", rpc.From, rpc.Payload)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Printf("do stuff every x senconds")
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
