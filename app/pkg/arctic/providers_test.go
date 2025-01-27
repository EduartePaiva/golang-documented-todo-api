package arctic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GoogleProvider(t *testing.T) {
	google := Google("superClientId", "superSecret", "https://test.com")
	url := google.CreateAuthorizationURL("superState", "codeVerifier", []string{"scope1", "scope2"})
	assert.Equal(t, url, `https://accounts.google.com/o/oauth2/v2/auth?client_id=superClientId&code_challenge=N1E4yRMD7xixn_oFyO_W3htYN3rY7-HMDKJe6z6r928&code_challenge_method=S256&redirect_uri=https%3A%2F%2Ftest.com&response_type=code&scope=scope1+scope2&state=superState`)
}
