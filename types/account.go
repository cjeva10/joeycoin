package types

import (
	"crypto/ecdsa"
	"math/big"
)

type Account struct {
	Address *ecdsa.PublicKey
	Balance *big.Int
	Nonce   *big.Int
}

func (acc *Account) Bytes() []byte {
    return acc.Address.X.Bytes()
}
