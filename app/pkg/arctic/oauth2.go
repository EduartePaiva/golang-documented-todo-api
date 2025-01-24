package arctic

import (
	"github.com/golang-documented-todo-api/app/pkg/crypto"
	"github.com/golang-documented-todo-api/app/pkg/encoding"
)

func GenerateState() (string, error) {
	randomValues, err := crypto.GetRandomValues(32)
	if err != nil {
		return "", err
	}
	return encoding.EncodeBase64urlNoPadding(randomValues), nil
}

func CreateS256CodeChallenge(codeVerifier string) string {
	codeChallengeBytes := crypto.Sha256([]byte(codeVerifier))
	return encoding.EncodeBase64urlNoPadding(codeChallengeBytes)
}

func GenerateCodeVerifier() (string, error) {
	randomValues, err := crypto.GetRandomValues(32)
	if err != nil {
		return "", err
	}
	return encoding.EncodeBase64urlNoPadding(randomValues), nil
}
