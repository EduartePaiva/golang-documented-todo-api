package encoding

import (
	"testing"

	"github.com/golang-documented-todo-api/app/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestEncodeHexLowerCase(t *testing.T) {
	assert.Equal(t, EncodeHexLowerCase(crypto.Sha256([]byte("testing"))), "cf80cd8aed482d5d1527d7dc72fceff84e6326592848447d2dc0b0e87dfc9a90")
}
