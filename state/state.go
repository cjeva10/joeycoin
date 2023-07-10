package state

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/cjeva10/joeycoin/types"
)

// Definition of current chain state
// Basically the state of all accounts (k-v storage)
// and the latest block on the chain
type State struct {
	Accounts    map[*ecdsa.PublicKey]*types.Account
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
		from := tx.From.Address
		to := tx.To.Address
		amount := tx.Amount

		// change balances and update sender nonce
		next.Accounts[from].Balance.Sub(next.Accounts[tx.From.Address].Balance, amount)
		next.Accounts[to].Balance.Add(next.Accounts[tx.To.Address].Balance, amount)
		next.Accounts[from].Nonce.Add(next.Accounts[from].Nonce, big.NewInt(1))
	}

	return next, nil
}

// is the given transaction valid on a given state
func (curr *State) ValidTx(tx *types.Transaction) bool {
	from := tx.From.Bytes()
	to := tx.To.Bytes()
	amount := tx.Amount.Bytes()

	// verify hash
	txBytes := append(from, to...)
	txBytes = append(txBytes, amount...)
	hash := sha256.Sum256(txBytes)
	if hash != tx.Hash {
		return false
	}

	// verify signature
	r, s := tx.Signature()
	valid := ecdsa.Verify(tx.From.Address, hash[:], r, s)

	// verify from account balance
	if -1 == tx.From.Balance.Cmp(tx.Amount) {
		return false
	}

	// if the state nonce + 1 doesnt equal submitted nonce then it is invalid
	// may need some better logic to handle a nonce gap
	// i.e. what happens if a higher nonce transaction reaches us first before the next transaction does?
	stateNonce := curr.Accounts[tx.From.Address].Nonce
	if big.NewInt(0).Add(stateNonce, big.NewInt(1)) != tx.From.Nonce {
		return false
	}

	return valid
}
