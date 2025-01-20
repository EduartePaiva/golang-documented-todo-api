package encoding

import "encoding/base32"

const base32UpperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
const base32LowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz234567"

func EncodeBase32LowerCaseNoPadding(bytes []byte) string {
	return encodeBase32_internal(bytes, base32LowerCaseAlphabet, false)
}

func encodeBase32_internal(bytes []byte, alphabet string, padding bool) string {
	enc := base32.NewEncoding(alphabet)
	if padding {
		return enc.EncodeToString(bytes)
	}
	encNoPad := enc.WithPadding(-1)
	return encNoPad.EncodeToString(bytes)
}
