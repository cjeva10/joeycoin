package state

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cjeva10/joeycoin/types"
)

// Definition of current chain state
// Basically the state of all accounts (k-v storage)
// and the latest block on the chain
type State struct {
	Accounts    map[types.Address]types.AccountState
	LatestBlock types.Block
}

// compute the next state using the current state and the given block
func (curr *State) Next(block *types.Block) (*State, error) {
	next := curr
	if !block.ValidLen() {
		return nil, fmt.Errorf("Proposed block length invalid (must be >=0 and <=100)")
	}

	if !block.ValidHash() {
		return nil, fmt.Errorf("Proposed block has invalid hash")
	}

	// loop through all transactions and update the intermediate state
	for _, tx := range block.Body {
		if !next.ValidTx(&tx) {
			return nil, fmt.Errorf("Invalid tx in this block 0x%s", hex.EncodeToString(block.Hash[:]))
		}

		// change balances and update sender nonce
	}

	return next, nil
}

// is the given transaction valid on a given state
func (curr *State) ValidTx(tx *types.SignedTransaction) bool {
	from := tx.From.Bytes()
	to := tx.To.Bytes()
	amount := types.Itob(tx.Amount)

	// verify account balance
    balance := curr.Accounts[tx.From].Balance
    if balance < tx.Amount {
        return false
    }

	// if the state nonce + 1 doesnt equal submitted nonce then it is invalid
	stateNonce := curr.Accounts[tx.From].Nonce
	if stateNonce + 1 != tx.Nonce {
		return false
	}

	// verify hash
	txBytes := append(from, to...)
	txBytes = append(txBytes, amount...)
	hash := sha256.Sum256(txBytes)
	if hash != tx.Hash() {
		return false
	}
	// verify signature
	r, s := tx.Signature()
    pub := ecdsa.PublicKey(tx.From)
	valid := ecdsa.Verify(&pub, hash[:], r, s)

	return valid
}
