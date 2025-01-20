package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	// Test if this function works the same way that the Oslo.js package
	assert.Equal(t, Sha256([]byte("abc")), []byte{186, 120, 22, 191, 143, 1, 207, 234, 65, 65, 64, 222, 93, 174, 34, 35, 176, 3, 97, 163, 150, 23, 122, 156, 180, 16, 255, 97, 242, 0, 21, 173})

	assert.Equal(t, Sha256([]byte("testing")), []byte{207, 128, 205, 138, 237, 72, 45, 93, 21, 39, 215, 220, 114, 252, 239, 248, 78, 99, 38, 89, 40, 72, 68, 125, 45, 192, 176, 232, 125, 252, 154, 144})
}
