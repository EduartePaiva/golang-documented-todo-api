package auth

import (
	"crypto/rand"
	"math/big"

	"github.com/golang-documented-todo-api/app/pkg/encoding"
)

func GenerateSessionToken() (string, error) {
	// Generate a random slice of byte of length 20
	bytes := make([]byte, 0, 20)
	for i := 0; i < 20; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(1<<8))
		if err != nil {
			return "", err
		}
		bytes = append(bytes, byte(num.Uint64()))
	}
	token := encoding.EncodeBase32LowerCaseNoPadding(bytes)
	return token, nil
}

func CreateSession(token string, userId string) {

}
