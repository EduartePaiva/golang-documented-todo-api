package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSessionToken(t *testing.T) {
	v, err := GenerateSessionToken()
	t.Log(len(v), v)
	// This should never error
	assert.Nil(t, err)
	assert.NotEmpty(t, v)
}
