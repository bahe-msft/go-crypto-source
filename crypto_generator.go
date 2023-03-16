package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
)

const secureStringLetters = "0123456789abcdefghijklmnopqrstuvwxyz"

type secureStringGenerator struct{}

// NewSecureStringGenerator creates a secure string generator with letter set 0123456789abcdefghijklmnopqrstuvwxyz.
func NewSecureStringGenerator() *secureStringGenerator {
	return &secureStringGenerator{}
}

func (r *secureStringGenerator) Generate(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("invalid length")
	}

	// ref: https://github.com/kubernetes/cluster-bootstrap/blob/v0.26.2/token/util/helpers.go
	// len("0123456789abcdefghijklmnopqrstuvwxyz") = 36 which doesn't evenly divide
	// the possible values of a byte: 256 mod 36 = 4. Discard any random bytes we
	// read that are >= 252 so the bytes we evenly divide the character set.
	const maxByteValue = 252

	var (
		b     byte
		err   error
		token = make([]byte, length)
	)

	reader := bufio.NewReaderSize(rand.Reader, length*2)
	for i := range token {
		for {
			if b, err = reader.ReadByte(); err != nil {
				return "", err
			}
			if b < maxByteValue {
				break
			}
		}

		token[i] = secureStringLetters[int(b)%len(secureStringLetters)]
	}

	return string(token), nil
}
