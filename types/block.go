package types

// define the structure of a block

import (
	"math/big"
    "crypto/sha256"
)

type Block struct {
	Body      []Transaction
	Hash      [32]byte
	PrevHash  [32]byte
	Number    *big.Int
	Timestamp *big.Int
}

func (block *Block) ValidLen() bool {
    return len(block.Body) <= 100 && len(block.Body) >= 0
}

// block hash must equal hash of concatenated number, timestamp, tx hashes
func (block *Block) ValidHash() bool {
    transactions := block.Body

    txBytes := append(block.Number.Bytes(), block.Timestamp.Bytes()...)
    for _, transaction := range transactions {
        txBytes = append(txBytes, transaction.Hash[:]...) 
    }

    computedHash := sha256.Sum256(txBytes)

    if computedHash != block.Hash {
        return false
    }

    return true
}
