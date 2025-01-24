package arctic

import (
	"context"
	"net/url"
	"strings"
)

type gitHub struct {
	CreateAuthorizationURL    func(state string, scopes []string) string
	ValidateAuthorizationCode func(ctx context.Context, code string) (OAuth2Tokens, error)
}

type GithubUserData struct {
	ID        int64  `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}

const (
	authorizationEndpoint = "https://github.com/login/oauth/authorize"
	tokenEndpoint         = "https://github.com/login/oauth/access_token"
)

func GitHub(clientId string, clientSecret string, redirectURI string) gitHub {
	return gitHub{
		CreateAuthorizationURL: func(state string, scopes []string) string {
			parsedURL, err := url.Parse(authorizationEndpoint)
			if err != nil {
				panic("invalid authorizationEndpoint")
			}
			queryParams := url.Values{}
			queryParams.Add("response_type", "code")
			queryParams.Add("client_id", clientId)
			queryParams.Add("redirect_uri", redirectURI)
			queryParams.Add("scope", strings.Join(scopes, " "))
			queryParams.Add("state", state)
			parsedURL.RawQuery = queryParams.Encode()
			return parsedURL.String()
		},
		ValidateAuthorizationCode: func(ctx context.Context, code string) (OAuth2Tokens, error) {
			body := url.Values{}
			body.Add("grant_type", "authorization_code")
			body.Add("code", code)
			body.Add("redirect_uri", redirectURI)
			request, err := createOAuth2Request(ctx, tokenEndpoint, body)
			if err != nil {
				return nil, err
			}
			encodedCredentials := encodeBasicCredentials(clientId, clientSecret)
			request.Header.Add("Authorization", "Basic "+encodedCredentials)
			return sendTokenRequest(request)
		},
	}
}
