package main

import (
	"fmt"
	"math/rand"
)

func NewMathRandStringGenerator() *mathRandStringGenerator {
	return &mathRandStringGenerator{}
}

type mathRandStringGenerator struct{}

func (r *mathRandStringGenerator) Generate(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("invalid length")
	}

	s := ""
	for i := 0; i < length; i++ {
		s += string(secureStringLetters[rand.Intn(len(secureStringLetters))])
	}

	return s, nil
}

type fixedStringGenerator struct{}

func (f *fixedStringGenerator) Generate(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("invalid length")
	}

	s := ""
	for i := 0; i < length; i++ {
		s += "a"
	}

	return s, nil
}
