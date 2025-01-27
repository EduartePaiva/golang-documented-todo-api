package encoding

import (
	"encoding/base64"
)

const base64Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const base64urlAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func EncodeBase64urlNoPadding(bytes []byte) string {
	return encodeBase64_internal(bytes, base64urlAlphabet, false)
}

func EncodeBase64(bytes []byte) string {
	return encodeBase64_internal(bytes, base64Alphabet, true)
}

func encodeBase64_internal(bytes []byte, alphabet string, padding bool) string {
	enc := base64.NewEncoding(alphabet)
	if padding {
		return enc.EncodeToString(bytes)
	}
	encNoPad := enc.WithPadding(base64.NoPadding)
	return encNoPad.EncodeToString(bytes)
}

func DecodeBase64urlIgnorePadding(encoded string) ([]byte, error) {
	return decodeBase64_internal(encoded, base64urlAlphabet, false)
}

func decodeBase64_internal(encoded string, alphabet string, padding bool) ([]byte, error) {
	enc := base64.NewEncoding(alphabet)
	// if padding is false mutate to no padding
	if !padding {
		enc = enc.WithPadding(base64.NoPadding)
	}
	return enc.DecodeString(encoded)
}
