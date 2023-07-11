package miner

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"testing"
	"time"

	"github.com/cjeva10/joeycoin/types"
)

func TestProof(t *testing.T) {
	now := big.NewInt(time.Now().Unix())
	digest := append(big.NewInt(0).Bytes(), now.Bytes()...)
	digest = append(digest, nil...)

	reader := rand.Reader
	priv1, err := ecdsa.GenerateKey(elliptic.P256(), reader)
	if err != nil {
		t.Fatalf("Failed to generate a private key")
	}

	var pub1 types.Address = types.Address(priv1.PublicKey)

	genesis := types.Block{
		Body:      nil,
		PrevHash:  [32]byte{},
		Hash:      sha256.Sum256(digest),
		Number:    big.NewInt(0),
		Timestamp: now,
	}

	// generate the proof
	genesisProof := &BlockProof{
		Block: &genesis,
		Miner: &pub1,
		Work:  0,
	}

	i := 0
	for !genesisProof.Valid() {
		// increase the nonce
		genesisProof.Work++
		i++
	}

	t.Logf("Generated valid proof after %d iterations", i)
}
