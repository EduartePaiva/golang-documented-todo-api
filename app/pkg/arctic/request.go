package arctic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/golang-documented-todo-api/app/pkg/encoding"
)

func createOAuth2Request(ctx context.Context, tokenEndpoint string, body url.Values) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, "POST", tokenEndpoint, bytes.NewReader([]byte(body.Encode())))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "arctic")
	return request, nil
}

func encodeBasicCredentials(username string, password string) string {
	bytes := []byte(fmt.Sprintf("%v:%v", username, password))
	return encoding.EncodeBase64(bytes)
}

func sendTokenRequest(req *http.Request) (OAuth2Tokens, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	data := json.NewDecoder(res.Body)
	v := make(map[string]string)
	data.Decode(&v)

	return nil, nil
}
