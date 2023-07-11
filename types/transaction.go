package types

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

type RawTransaction struct {
	From   Address
	To     Address
	Amount int64
}

type SignedTransaction struct {
	From   Address
	To     Address
	Amount int64
	r      *big.Int
	s      *big.Int
}

func Itob(x int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(x))

	return b
}

func (tx *RawTransaction) Sign(key *ecdsa.PrivateKey) (*SignedTransaction, error) {
	reader := rand.Reader
	hash := tx.Hash()
	r, s, err := ecdsa.Sign(reader, key, hash[:])
	if err != nil {
		return nil, err
	}

	return &SignedTransaction{
		tx.From,
		tx.To,
		tx.Amount,
		r,
		s,
	}, nil
}

func (tx *RawTransaction) Hash() [32]byte {
	bytes := append(tx.From.Bytes(), tx.To.Bytes()...)
	bytes = append(bytes, Itob(tx.Amount)...)

	return sha256.Sum256(bytes)
}

func (tx *SignedTransaction) Hash() [32]byte {
	bytes := append(tx.From.Bytes(), tx.To.Bytes()...)
	bytes = append(bytes, Itob(tx.Amount)...)

	return sha256.Sum256(bytes)
}

func (tx *SignedTransaction) Signature() (r *big.Int, s *big.Int) {
	return tx.r, tx.s
}
