package utils

import (
	"math/rand"
	"time"
)

func GenerateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[gen.Intn(len(charset))]
	}

	return string(code)
}
