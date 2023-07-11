package types

// define the structure of a block

import (
	"crypto/sha256"
	"math/big"
)

type Block struct {
	Body      []SignedTransaction
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
		hash := transaction.Hash()
		txBytes = append(txBytes, hash[:]...)
	}

	computedHash := sha256.Sum256(txBytes)

	if computedHash != block.Hash {
		return false
	}

	return true
}
