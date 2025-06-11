package hashing

import (
	"crypto/sha512"
	"encoding/hex"
)

func hashData(data []byte) string {
	hashBytes := sha512.Sum512(data)
	return hex.EncodeToString(hashBytes[:])
}

func Compare(data []byte, expected string) bool {
	return hashData(data) == expected
}

func HashString(data string) string {
	return hashData([]byte(data))
}
