package arctic

import (
	"net/url"
	"strings"
)

type gitHub struct {
	CreateAuthorizationURL func(state string, scopes []string) string
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
			queryParams.Add("state", state)
			queryParams.Add("scope", strings.Join(scopes, " "))
			queryParams.Add("redirect_uri", redirectURI)
			parsedURL.RawQuery = queryParams.Encode()
			return parsedURL.String()
		},
	}
}
