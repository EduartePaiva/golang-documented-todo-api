package arctic

import (
	"context"
	"net/url"
	"strings"
)

const (
	githubAuthorizationEndpoint   = "https://github.com/login/oauth/authorize"
	githubTokenEndpoint           = "https://github.com/login/oauth/access_token"
	googleAuthorizationEndpoint   = "https://accounts.google.com/o/oauth2/v2/auth"
	googleTokenEndpoint           = "https://oauth2.googleapis.com/token"
	googleTokenRevocationEndpoint = "https://oauth2.googleapis.com/revoke"
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

func GitHub(clientId string, clientSecret string, redirectURI string) gitHub {
	return gitHub{
		CreateAuthorizationURL: func(state string, scopes []string) string {
			parsedURL, err := url.Parse(githubAuthorizationEndpoint)
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
			request, err := createOAuth2Request(ctx, githubTokenEndpoint, body)
			if err != nil {
				return nil, err
			}
			encodedCredentials := encodeBasicCredentials(clientId, clientSecret)
			request.Header.Add("Authorization", "Basic "+encodedCredentials)
			return sendTokenRequest(request)
		},
	}
}

type google struct {
	CreateAuthorizationURL    func(state string, codeVerifier string, scopes []string) string
	ValidateAuthorizationCode func(ctx context.Context, code string) (OAuth2Tokens, error)
}

func Google(clientId string, clientSecret string, redirectURI string) google {
	client := OAuth2Client(clientId, clientSecret, &redirectURI)
	return google{
		CreateAuthorizationURL: func(state string, codeVerifier string, scopes []string) string {
			return client.CreateAuthorizationURLWithPKCE(
				googleAuthorizationEndpoint,
				state,
				S256,
				codeVerifier,
				scopes,
			)
		},
		ValidateAuthorizationCode: func(ctx context.Context, code string) (OAuth2Tokens, error) {},
	}
}

type codeChallengeMethod uint8

const (
	S256 codeChallengeMethod = iota
	Plain
)

type oAuth2Client struct {
	CreateAuthorizationURL         func(authorizationEndpoint string, state string, scopes []string) string
	CreateAuthorizationURLWithPKCE func(
		authorizationEndpoint string,
		state string,
		codeChallenge codeChallengeMethod,
		codeVerifier string,
		scopes []string,
	) string
}

func OAuth2Client(clientId string, clientSecret string, redirectURI *string) oAuth2Client {
	return oAuth2Client{
		CreateAuthorizationURL: func(authorizationEndpoint, state string, scopes []string) string {
			url := url.Values{}
			url.Add("response_type", "code")
			url.Add("client_id", clientId)
			if redirectURI != nil {
				url.Add("redirect_uri", *redirectURI)
			}
			url.Add("state", state)
			if len(scopes) > 0 {
				url.Add("scope", strings.Join(scopes, " "))
			}

			return url.Encode()
		},
		CreateAuthorizationURLWithPKCE: func(
			authorizationEndpoint string,
			state string,
			codeChallenge codeChallengeMethod,
			codeVerifier string,
			scopes []string,
		) string {
			url := url.Values{}
			url.Add("response_type", "code")
			url.Add("client_id", clientId)
			if redirectURI != nil {
				url.Add("redirect_uri", *redirectURI)
			}
			url.Add("state", state)
			if codeChallenge == S256 {
				codeChallenge := CreateS256CodeChallenge(codeVerifier)
				url.Add("code_challenge_method", "S256")
				url.Add("code_challenge", codeChallenge)
			} else if codeChallenge == Plain {
				url.Add("code_challenge_method", "plain")
				url.Add("code_challenge", codeVerifier)
			}
			if len(scopes) > 0 {
				url.Add("scope", strings.Join(scopes, " "))
			}
			return url.Encode()
		},
	}
}
