package common

import (
	"crypto/sha256"
)

// Sha256 computes the SHA256 of an array of bytes
// *Wrapper of the Golang implementation sha256
func Sha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
