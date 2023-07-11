package types

import (
	"crypto/ecdsa"
)

type Address ecdsa.PublicKey

func (addr *Address) Bytes() []byte {
    return addr.Bytes()
}

type Account struct {
	Address Address 
	Balance int64 
	Nonce   int64 
}

func (acc *Account) Bytes() []byte {
    return acc.Address.X.Bytes()
}
