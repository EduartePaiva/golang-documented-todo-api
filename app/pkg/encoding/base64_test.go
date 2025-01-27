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

func TestDecodeBase64urlIgnorePadding(t *testing.T) {
	decoded, err := DecodeBase64urlIgnorePadding("TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdCwgc2VkIGRvIGVpdXNtb2QgdGVtcG9yIGluY2lkaWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZSBtYWduYSBhbGlxdWEuIFV0IGVuaW0gYWQgbWluaW0gdmVuaWFtLCBxdWlzIG5vc3RydWQgZXhlcmNpdGF0aW9uIHVsbGFtY28gbGFib3JpcyBuaXNpIHV0IGFsaXF1aXAgZXggZWEgY29tbW9kbyBjb25zZXF1YXQuIER1aXMgYXV0ZSBpcnVyZSBkb2xvciBpbiByZXByZWhlbmRlcml0IGluIHZvbHVwdGF0ZSB2ZWxpdCBlc3NlIGNpbGx1bSBkb2xvcmUgZXUgZnVnaWF0IG51bGxhIHBhcmlhdHVyLiBFeGNlcHRldXIgc2ludCBvY2NhZWNhdCBjdXBpZGF0YXQgbm9uIHByb2lkZW50LCBzdW50IGluIGN1bHBhIHF1aSBvZmZpY2lhIGRlc2VydW50IG1vbGxpdCBhbmltIGlkIGVzdCBsYWJvcnVtLg")
	assert.NoError(t, err)
	assert.Equal(
		t,
		decoded,
		[]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
	)

	// test the ignore padding, if a pedding exists it'll error. That's expected
	decoded, err = DecodeBase64urlIgnorePadding("cGFkZGluZw==")
	assert.Error(t, err)
	assert.NotEqual(t, decoded, []byte("padding"))
}
