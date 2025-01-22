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
