package rpc

import (
	"net"

	"github.com/cjeva10/joeycoin/types"
)

// methods for communicating directly with peers

type Peer struct {
	ip      net.IP
	address *types.Address
}

type BlockMessage struct {
	Body      []types.SignedTransaction
	Hash      [32]byte
	PrevHash  [32]byte
	Number    int64 
	Timestamp int64
    Miner     *types.Address
    Work      int64
}

// send a new block head to peer
func (peer *Peer) SendBlock(message *BlockMessage) bool {
	return true
}
