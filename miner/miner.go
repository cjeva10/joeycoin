// functions for mining blocks

package miner

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/cjeva10/joeycoin/types"
)

// an interface for generating a proof of work
type Proof interface {
	Prove() [32]byte
	Valid() bool
}

type BlockProof struct {
	Block *types.Block
	Miner *types.Address
	Work  int64 // nonce
}

// generate a proof of work by hashing the block hash with the nonce
func (proof *BlockProof) Prove() [32]byte {
	workBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(workBytes, uint64(proof.Work))

	blockBytes := append(proof.Block.Hash[:], proof.Miner.X.Bytes()...)
    blockBytes = append(blockBytes, proof.Miner.Y.Bytes()...)
	blockBytes = append(blockBytes, workBytes...)

	return sha256.Sum256(blockBytes)
}

// check if a proof is valid
func (proof *BlockProof) Valid() bool {
	proofInt := new(big.Int)
	prove := proof.Prove()
	proofInt.SetBytes(prove[:])

	DIFFICULTY := []byte{0, 0, 128, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	diffInt := new(big.Int)
	diffInt.SetBytes(DIFFICULTY)

	if -1 == proofInt.Cmp(diffInt) {
		return true
	}

	return false
}

// create a valid proof of work for a given block and account
func BuildProof(block *types.Block, addr *types.Address) BlockProof {
    proof := BlockProof{
        Block: block,
        Miner: addr,
        Work: 0,
    }

    // simply keep incrementing the work until the proof is valid
    for !proof.Valid() {
        proof.Work++
    }

    return proof
}
