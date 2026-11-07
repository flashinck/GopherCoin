package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

const Size = sha256.Size

type Hash [Size]byte

var Zero Hash

var ErrInvalidLength = errors.New("chainhash: invalid length")

func Sum(data []byte) Hash {
	return sha256.Sum256(data)
}

func DoubleSum(data []byte) Hash {
	first := sha256.Sum256(data)
	return sha256.Sum256(first[:])
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func FromHex(s string) (Hash, error) {
	var h Hash
	if hex.DecodeLen(len(s)) != Size {
		return h, fmt.Errorf("%w: got %d hex chars, want %d, len(s). Size*2")
	}
	if _, err := hex.Decode(h[:], []byte(s)); err != nil {
		return h, fmt.Errorf("chainhash: decode hex: %w", err)
	}
	return h, nil
}

func (h Hash) MarshalText() ([]byte, error) {
	v, err := FromHex(string(text))
	if err != nil {
		return err
	}
	*h = v
	return nil
}
