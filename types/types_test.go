package types

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"fmt"
	"math/big"

	"testing"
)

func TestNewAccount(t *testing.T) {
	reader := rand.Reader
	priv, err := ecdsa.GenerateKey(elliptic.P256(), reader)
	if err != nil {
		t.Fatalf("Failed to generate a private key")
	}

	pub := priv.PublicKey

	account := Account{
		Address: &pub,
		Balance: big.NewInt(0),
		Nonce:   big.NewInt(0),
	}

	if !account.Address.Equal(&pub) {
		fmt.Println(account.Address)
		fmt.Println(pub)
		t.Fatalf("Address doesn't equal generated public key")
	}

	if account.Balance.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Starting balance wasn't zero, was %d", account.Balance)
	}

	if account.Nonce.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Starting Nonce must be zero")
	}
}

func TestNewTransaction(t *testing.T) {
	reader := rand.Reader
	priv1, err := ecdsa.GenerateKey(elliptic.P256(), reader)
	if err != nil {
		t.Fatalf("Failed to generate a private key")
	}

	pub1 := priv1.PublicKey

	priv2, err := ecdsa.GenerateKey(elliptic.P256(), reader)
	if err != nil {
		t.Fatalf("Failed to generate a private key")
	}

	pub2 := priv2.PublicKey

	account1 := Account{
		Address: &pub1,
		Balance: big.NewInt(1000),
		Nonce:   big.NewInt(0),
	}

	account2 := Account{
		Address: &pub2,
		Balance: big.NewInt(1000),
		Nonce:   big.NewInt(0),
	}

    _ = Transaction{
		From:   account1,
		To:     account2,
		Amount: big.NewInt(1),
		Hash:   [32]byte{},
		r:      big.NewInt(0),
		s:      big.NewInt(0),
	}
}
