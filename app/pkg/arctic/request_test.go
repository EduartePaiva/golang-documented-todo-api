package arctic

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createOAuth2Request(t *testing.T) {
	body := url.Values{}
	body.Add("grant_type", "authorization_code")
	body.Add("code", "code")
	body.Add("redirect_uri", "redirectURI")
	body.Add("code_verifier", "codeVerifier")
	body.Add("client_id", "this is a client id")

	request, err := createOAuth2Request(context.Background(), "https://test.com", body)
	request.Header.Add("Authorization", "Basic clientid:clientpass")
	assert.NoError(t, err)
	t.Log("logging the header:\n")
	t.Log(fmt.Sprint(request.Header))
	data := make([]byte, len(body.Encode()))
	request.Body.Read(data)
	t.Log(fmt.Sprint(string(data)))
}
