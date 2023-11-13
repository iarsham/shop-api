package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func GeneratePassword() string {
	length := 20
	charset := letterBytes + specialBytes + numBytes
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		password[i] = charset[randomIndex.Int64()]
	}
	return string(password)
}
