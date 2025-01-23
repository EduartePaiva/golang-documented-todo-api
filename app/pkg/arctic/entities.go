package arctic

import (
	"fmt"
	"strings"
	"time"
)

type OAuth2Tokens interface {
	TokenType() (string, error)
	AccessToken() (string, error)
	AccessTokenExpiresInSeconds() (int, error)
	AccessTokenExpiresAt() (time.Time, error)
	RefreshToken() (string, error)
	IdToken() (string, error)
	Scopes() ([]string, error)
}

type oAuth2Tokens struct {
	data map[string]interface{}
}

func (a *oAuth2Tokens) TokenType() (string, error) {
	v, ok := a.data["token_type"]
	if !ok {
		return "", fmt.Errorf("missing or invalid 'token_type' field")
	}
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("missing or invalid 'token_type' field")
	}
	return token, nil
}
func (a *oAuth2Tokens) AccessToken() (string, error) {
	v, ok := a.data["access_token"]
	if !ok {
		return "", fmt.Errorf("missing or invalid 'access_token' field")
	}
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("missing or invalid 'access_token' field")
	}
	return token, nil
}
func (a *oAuth2Tokens) RefreshToken() (string, error) {
	v, ok := a.data["refresh_token"]
	if !ok {
		return "", fmt.Errorf("missing or invalid 'refresh_token' field")
	}
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("missing or invalid 'refresh_token' field")
	}
	return token, nil
}
func (a *oAuth2Tokens) IdToken() (string, error) {
	v, ok := a.data["id_token"]
	if !ok {
		return "", fmt.Errorf("missing or invalid 'id_token' field")
	}
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("missing or invalid 'id_token' field")
	}
	return token, nil
}
func (a *oAuth2Tokens) Scopes() ([]string, error) {
	v, ok := a.data["scope"]
	if !ok {
		return []string{}, fmt.Errorf("missing or invalid 'scope' field")
	}
	token, ok := v.(string)
	if !ok {
		return []string{}, fmt.Errorf("missing or invalid 'scope' field")
	}
	return strings.Split(token, " "), nil
}
func (a *oAuth2Tokens) AccessTokenExpiresInSeconds() (int, error) {
	v, ok := a.data["expires_in"]
	if !ok {
		return 0, fmt.Errorf("missing or invalid 'access_token' field")
	}
	token, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("missing or invalid 'expires_in' field")
	}
	return token, nil
}
func (a *oAuth2Tokens) AccessTokenExpiresAt() (time.Time, error) {
	seconds, err := a.AccessTokenExpiresInSeconds()
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(time.Second * time.Duration(seconds)), nil
}

type OAuth2RequestError struct {
	code        string
	description string
	uri         string
	state       string
}

func (e OAuth2RequestError) Error() string {
	return fmt.Sprintf("oauth2 request error: %v", e.code)
}
