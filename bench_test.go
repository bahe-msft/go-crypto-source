package main

import "testing"

func BenchmarkSecureStringGenerator_Length16(b *testing.B) {
	const length = 16
	g := NewSecureStringGenerator()

	for n := 0; n < b.N; n++ {
		_, _ = g.Generate(length)
	}
}

func BenchmarkMathRandStringGenerator(b *testing.B) {
	const length = 16
	g := NewMathRandStringGenerator()

	for n := 0; n < b.N; n++ {
		_, _ = g.Generate(length)
	}
}

func BenchmarkFixedStringGenerator(b *testing.B) {
	const length = 16
	g := &fixedStringGenerator{}

	for n := 0; n < b.N; n++ {
		_, _ = g.Generate(length)
	}
}
