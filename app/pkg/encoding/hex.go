package encoding

import "encoding/hex"

func EncodeHexLowerCase(data []byte) string {
	dst := make([]byte, 0, 100)
	hex.Encode(dst, data)
	return string(dst)
}
