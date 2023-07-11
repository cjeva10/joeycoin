package state

import (
	"crypto/sha256"
	"math/big"
	"time"

	"github.com/cjeva10/joeycoin/types"
)

// build the genesis state of the chain

func Genesis() *State {
    now := big.NewInt(time.Now().Unix())
    digest := append(big.NewInt(0).Bytes(), now.Bytes()...)
    digest = append(digest, nil...)
    
    return &State{
        Accounts: nil,
        LatestBlock: types.Block{
            Body: nil,
            PrevHash: [32]byte{},
            Hash: sha256.Sum256(digest),
            Number: big.NewInt(0),
            Timestamp: now, 
        },
    }
}
