package crypto

import "crypto/sha256"

// Synchronously hashes data with SHA-256
func Sha256(token string) []byte {
	hash := sha256.New()
	hash.Write([]byte(token))
	result := hash.Sum(nil)
	return result
}
