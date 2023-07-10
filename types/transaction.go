package types

import (
	"math/big"
)

type Transaction struct {
	From   Account
	To     Account
	Amount *big.Int
	Hash   [32]byte
	r      *big.Int
	s      *big.Int
}

func (tx *Transaction) Signature() (r *big.Int, s *big.Int) {
	return tx.r, tx.s
}
