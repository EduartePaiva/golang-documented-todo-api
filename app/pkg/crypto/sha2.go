package crypto

import "crypto/sha256"

// Hashes with SHA-256 and returns the hashed data
func Sha256(token []byte) []byte {
	hash := sha256.New()
	hash.Write(token)
	result := hash.Sum(nil)
	return result
}
