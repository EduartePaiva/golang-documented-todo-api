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
	encNoPad := enc.WithPadding(-1)
	return encNoPad.EncodeToString(bytes)
}
