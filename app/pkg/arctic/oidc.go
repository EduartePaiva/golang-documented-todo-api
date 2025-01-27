package arctic

import "github.com/golang-documented-todo-api/app/pkg/encoding"

// The mold is the data that'll be filled. If mold is a struct, it must contain the json tags
func DecodeIdToken[T any](token string, receiver *T) error {
	return encoding.DecodeJWT(token, receiver)
}
