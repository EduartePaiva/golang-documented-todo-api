package encoding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBase32LowerCaseNoPadding(t *testing.T) {
	// Test if there are Padding "=" in this function
	bytes := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	v := EncodeBase32LowerCaseNoPadding(bytes)
	assert.NotContains(t, v, "=")

	// Test if UpperCase characters don't exists
	for _, c := range v {
		if c >= 'A' && c <= 'Z' {
			t.Error(fmt.Errorf("The value '%v' contains the upper letter %v", v, string(c)))
		}
	}
}
