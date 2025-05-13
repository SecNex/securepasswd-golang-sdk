package random

import (
	"math/rand"
	"time"
)

const (
	LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOLS = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	NUMBERS = "0123456789"
)

type RandomString []byte

func Generate(length int) RandomString {
	letters := []byte(LETTERS + SYMBOLS + NUMBERS)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return RandomString(b)
}

func (r *RandomString) String() string {
	return string(*r)
}
