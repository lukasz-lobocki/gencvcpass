package cmd

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

func runRootCommand(_ *cobra.Command, _ []string) error {
	password, err := GeneratePassword(appConfig)
	if err != nil {
		return err
	}

	fmt.Println(password)
	return nil
}

func GeneratePassword(cfg config) (string, error) {
	if err := cfg.validate(); err != nil {
		return "", err
	}

	cvcsString, err := getCVCCVCsString(cfg)
	if err != nil {
		return "", err
	}
	if loggingLevel >= 2 {
		logInfo.Printf("cvcsString= %s", cvcsString)
	}

	passString, err := swapUpperAndDigits(cvcsString, cfg)
	if err != nil {
		return "", err
	}
	if loggingLevel >= 2 {
		logInfo.Printf("passString= %s", passString)
	}

	sets, err := splitIntoSets(passString, CHUNKSIZE*CHUNKSPERSET)
	if err != nil {
		return "", err
	}
	if loggingLevel >= 2 {
		logInfo.Printf("sets= %v", sets)
	}

	return strings.Join(sets, cfg.separator), nil
}

func (cfg config) validate() error {
	if cfg.setsNum < 1 {
		return fmt.Errorf("sets=%d must be equal or greater than 1", cfg.setsNum)
	}
	if cfg.upperNum < 0 || cfg.upperNum > cfg.setsNum*2 {
		return fmt.Errorf("upper=%d must be between 0 and the number of sets=%d * 2", cfg.upperNum, cfg.setsNum)
	}
	if cfg.digitsNum < 0 || cfg.digitsNum > cfg.setsNum*2 {
		return fmt.Errorf("digits=%d must be between 0 and the number of sets=%d * 2", cfg.digitsNum, cfg.setsNum)
	}

	return nil
}

func swapUpperAndDigits(s string, cfg config) (string, error) {
	upper := func(x int) int {
		return CHUNKSIZE * x
	}
	digits := func(x int) int {
		return CHUNKSIZE*x + (CHUNKSIZE - 1)
	}

	upperSlice := getRandomPiece(getSlice(upper, cfg.setsNum*2), cfg.upperNum)
	digitSlice := getRandomPiece(getSlice(digits, cfg.setsNum*2), cfg.digitsNum)

	runes := []rune(s)
	for _, index := range upperSlice {
		runes[index] = unicode.ToUpper(runes[index])
	}

	for _, index := range digitSlice {
		randomDigitIndex, err := crand.Int(crand.Reader, big.NewInt(int64(len(DIGITS))))
		if err != nil {
			return "", err
		}

		runes[index] = rune(DIGITS[randomDigitIndex.Int64()])
	}

	return string(runes), nil
}

func splitIntoSets(s string, setSize int) (sets []string, err error) {
	if setSize <= 0 {
		return sets, fmt.Errorf("size of a set must be greater than 0")
	}

	for i := 0; i < len(s); i += setSize {
		end := min(i+setSize, len(s))
		sets = append(sets, s[i:end])
	}

	return sets, nil
}

func getCVCCVCsString(cfg config) (string, error) {
	var builder strings.Builder
	builder.Grow(cfg.setsNum * CHUNKSPERSET * CHUNKSIZE)

	for i := 0; i < cfg.setsNum*CHUNKSPERSET; i++ {
		cvc, err := getCVC(cfg)
		if err != nil {
			return "", err
		}
		builder.WriteString(cvc)
	}

	return builder.String(), nil
}

func getCVC(cfg config) (string, error) {
	var builder strings.Builder
	builder.Grow(CHUNKSIZE)

	useConsonant := false
	for i := 0; i < CHUNKSIZE; i++ {
		useConsonant = !useConsonant

		charSet := pickCharSet(cfg, useConsonant)
		randomIndex, err := crand.Int(crand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}

		builder.WriteByte(charSet[randomIndex.Int64()])
	}

	return builder.String(), nil
}

func pickCharSet(cfg config, useConsonant bool) string {
	if useConsonant {
		if cfg.lessNonPolish {
			return CONSONANTS_PL
		}
		return CONSONANTS
	}

	if cfg.lessNonPolish {
		return VOWELS_PL
	}

	return VOWELS
}
