package cmd

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

func getRandomPiece(slice []int, length int) []int {
	r := mrand.New(newCryptoRandSource())
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice[:length]
}

func newCryptoRandSource() *cryptoRandSource {
	return &cryptoRandSource{}
}

func (s *cryptoRandSource) Seed(seed int64) {}

func (s *cryptoRandSource) Int63() int64 {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		logError.Fatalf("crypto/rand: failed to read random bytes. %v", err)
	}

	return int64(binary.LittleEndian.Uint64(b[:]) & (1<<63 - 1))
}

func getSlice(function operation, iterations int) (seq []int) {
	for i := 0; i < iterations; i++ {
		seq = append(seq, compute(function, i))
	}

	return seq
}

func compute(op operation, a int) int {
	return op(a)
}
