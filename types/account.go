package types

import (
	"crypto/ecdsa"
)

type Address ecdsa.PublicKey

type AccountState struct {
	Balance int64
	Nonce   int64
}

type AccountMap struct {
    state map[Address]AccountState
}

func (addr *Address) Bytes() []byte {
    b := append(addr.X.Bytes(), addr.Y.Bytes()...)
    return b
}
