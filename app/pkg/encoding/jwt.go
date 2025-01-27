package encoding

import (
	"encoding/json"
	"fmt"
	"strings"
)

// get the jwt raw data
func DecodeJWT[T any](jwt string, receiver *T) error {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid JWT")
	}
	data, err := DecodeBase64urlIgnorePadding(parts[1])
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, receiver)
	if err != nil {
		return err
	}
	return nil
}
