package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testMold struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Iat  int64  `json:"iat"`
}

func TestDecodeJWT(t *testing.T) {
	decoded := testMold{}
	err := DecodeJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", &decoded)
	assert.NoError(t, err)

	assert.Equal(t, decoded.Sub, "1234567890")
	assert.Equal(t, decoded.Name, "John Doe")
	assert.Equal(t, decoded.Iat, int64(1516239022))

}
