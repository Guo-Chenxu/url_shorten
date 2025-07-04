package utils

import (
	"fmt"
	"math/big"
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var base = big.NewInt(62)

func EncodeBase62Byte(input []byte) string {
	if len(input) == 0 {
		return ""
	}

	x := new(big.Int).SetBytes(input)
	return EncodeBase62(x)
}

func EncodeBase62Uint64(input uint64) string {
	return EncodeBase62(big.NewInt(int64(input)))
}

func EncodeBase62(x *big.Int) string {
	var result string

	zero := big.NewInt(0)
	mod := big.NewInt(0)

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = string(charset[mod.Int64()]) + result
	}

	return result
}

func DecodeBase62Byte(input string) ([]byte, error) {
	x, err := DecodeBase62(input)
	return x.Bytes(), err
}

func DecodeBase62Uint64(input string) (uint64, error) {
	x, err := DecodeBase62(input)
	return uint64(x.Uint64()), err
}

func DecodeBase62(input string) (*big.Int, error) {
	if len(input) == 0 {
		return big.NewInt(0), nil
	}

	x := big.NewInt(0)
	charIndex := make(map[rune]int, len(charset))

	for index, char := range charset {
		charIndex[char] = index
	}

	for _, char := range input {
		index, ok := charIndex[char]
		if !ok {
			return big.NewInt(0), fmt.Errorf("invalid character in Base62 string: %c", char)
		}

		x.Mul(x, base)
		x.Add(x, big.NewInt(int64(index)))
	}

	return x, nil
}
