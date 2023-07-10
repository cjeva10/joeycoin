package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
    "encoding/hex"
	"fmt"
	"math/big"

	"github.com/cjeva10/joeycoin/types"
)

func main() {
	reader := rand.Reader
	priv, err := ecdsa.GenerateKey(elliptic.P224(), reader)
	if err != nil {
		panic("Failed to generate private key")
	}

	pub := priv.PublicKey

	acc := types.Account{
		Address: &pub,
		Balance: big.NewInt(0),
		Nonce:   big.NewInt(0),
	}

	fmt.Println(hex.EncodeToString(acc.Address.X.Bytes()))
	fmt.Println(acc.Balance)
	fmt.Println(acc.Nonce)
}
