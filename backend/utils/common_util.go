package utils

import "math/rand"

const CharSet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateCode(n int) string {
	code := make([]byte, n)
	for i := 0; i < n; i++ {
		code[i] = CharSet[rand.Intn(len(CharSet))]
	}
	return string(code)
}
