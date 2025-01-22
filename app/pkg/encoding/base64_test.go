package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBase64urlNoPadding(t *testing.T) {
	// Test if there are Padding "=" in this function
	bytes := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	v := EncodeBase64urlNoPadding(bytes)
	assert.NotContains(t, v, "=")
	t.Log(v)
	assert.Equal(t, v, "AQIDBAUGBwg")
}
