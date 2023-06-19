package benchmark_test

import (
	"crypto/sha1"
	"crypto/sha256"
	"testing"
)

func BenchmarkSum256(b *testing.B) {
	data := []byte("Digital House impulsionando a transformação digital")
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}
func BenchmarkSum(b *testing.B) {
	data := []byte("Digital House impulsionando a transformação digital")
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}
