package cmd

const (
	MAX_LOGGING_LEVEL int = 3 // Maximum allowed logging level.
	CONSONANTS            = "bcdfghjklmnpqrstvwxyz"
	VOWELS                = "aeiu"
	DIGITS                = "23456789"
	CHUNKSIZE             = 6
)

type tConfig struct {
	setsNum   int
	digitsNum int
	upperNum  int
	separator string
}

// operation implements function calculating positions of UPPER and digits
type operation func(int) int

// cryptoRandSource implements math/rand.Source
type cryptoRandSource struct{}
