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
	jsonData := json.NewDecoder(res.Body)
	data := make(map[string]interface{})
	err = jsonData.Decode(&data)
	if err != nil {
		return nil, err
	}
	if _, ok := data["error"]; ok {
		return nil, createOAuth2RequestError(data)
	}
	return &oAuth2Tokens{data: data}, nil
}

func createOAuth2RequestError(result map[string]interface{}) error {
	code, ok := result["error"].(string)
	if !ok {
		return fmt.Errorf("invalid error response")
	}
	description, _ := result["error_description"].(string)
	uri, _ := result["error_uri"].(string)
	state, _ := result["state"].(string)

	return OAuth2RequestError{
		code,
		description,
		uri,
		state,
	}
}
