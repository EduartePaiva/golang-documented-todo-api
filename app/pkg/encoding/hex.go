package encoding

import "encoding/hex"

func EncodeHexLowerCase(data []byte) string {
	return hex.EncodeToString(data)
}
