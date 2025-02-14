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

func GitHub(clientId string, clientSecret string, redirectURI string) gitHub {
	const (
		authorizationEndpoint string = "https://github.com/login/oauth/authorize"
		tokenEndpoint         string = "https://github.com/login/oauth/access_token"
	)
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

type google struct {
	CreateAuthorizationURL    func(state string, codeVerifier string, scopes []string) string
	ValidateAuthorizationCode func(ctx context.Context, code string, codeVerifier string) (OAuth2Tokens, error)
}
type GoogleUserData struct {
	ID        string `json:"sub"`
	AvatarURL string `json:"picture"`
	Name      string `json:"name"`
}

func Google(clientId string, clientSecret string, redirectURI string) google {
	const (
		authorizationEndpoint   string = "https://accounts.google.com/o/oauth2/v2/auth"
		tokenEndpoint           string = "https://oauth2.googleapis.com/token"
		tokenRevocationEndpoint string = "https://oauth2.googleapis.com/revoke"
	)

	client := OAuth2Client(clientId, &clientSecret, &redirectURI)

	return google{
		CreateAuthorizationURL: func(state string, codeVerifier string, scopes []string) string {
			return client.CreateAuthorizationURLWithPKCE(
				authorizationEndpoint,
				state,
				S256,
				codeVerifier,
				scopes,
			)
		},
		ValidateAuthorizationCode: func(ctx context.Context, code, codeVerifier string) (OAuth2Tokens, error) {
			return client.ValidateAuthorizationCode(ctx, tokenEndpoint, code, &codeVerifier)
		},
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
	ValidateAuthorizationCode func(
		ctx context.Context,
		tokenEndpoint string,
		code string,
		codeVerifier *string,
	) (OAuth2Tokens, error)
}

func OAuth2Client(clientId string, clientPassword *string, redirectURI *string) oAuth2Client {
	return oAuth2Client{
		CreateAuthorizationURL: func(authorizationEndpoint, state string, scopes []string) string {
			parsedURL, err := url.Parse(authorizationEndpoint)
			if err != nil {
				panic("invalid authorizationEndpoint")
			}
			queryParams := url.Values{}
			queryParams.Add("response_type", "code")
			queryParams.Add("client_id", clientId)
			if redirectURI != nil {
				queryParams.Add("redirect_uri", *redirectURI)
			}
			queryParams.Add("state", state)
			if len(scopes) > 0 {
				queryParams.Add("scope", strings.Join(scopes, " "))
			}

			parsedURL.RawQuery = queryParams.Encode()
			return parsedURL.String()
		},
		CreateAuthorizationURLWithPKCE: func(
			authorizationEndpoint string,
			state string,
			codeChallenge codeChallengeMethod,
			codeVerifier string,
			scopes []string,
		) string {
			parsedURL, err := url.Parse(authorizationEndpoint)
			if err != nil {
				panic("invalid authorizationEndpoint")
			}
			queryParams := url.Values{}
			queryParams.Add("response_type", "code")
			queryParams.Add("client_id", clientId)
			if redirectURI != nil {
				queryParams.Add("redirect_uri", *redirectURI)
			}
			queryParams.Add("state", state)
			if codeChallenge == S256 {
				codeChallenge := CreateS256CodeChallenge(codeVerifier)
				queryParams.Add("code_challenge_method", "S256")
				queryParams.Add("code_challenge", codeChallenge)
			} else if codeChallenge == Plain {
				queryParams.Add("code_challenge_method", "plain")
				queryParams.Add("code_challenge", codeVerifier)
			}
			if len(scopes) > 0 {
				queryParams.Add("scope", strings.Join(scopes, " "))
			}
			parsedURL.RawQuery = queryParams.Encode()
			return parsedURL.String()
		},
		ValidateAuthorizationCode: func(ctx context.Context, tokenEndpoint, code string, codeVerifier *string) (OAuth2Tokens, error) {
			body := url.Values{}
			body.Add("grant_type", "authorization_code")
			body.Add("code", code)
			if redirectURI != nil {
				body.Add("redirect_uri", *redirectURI)
			}
			if codeVerifier != nil {
				body.Add("code_verifier", *codeVerifier)
			}
			if clientPassword == nil {
				body.Add("client_id", clientId)
			}

			request, err := createOAuth2Request(ctx, tokenEndpoint, body)
			if err != nil {
				return nil, err
			}
			if clientPassword != nil {
				encodedCredentials := encodeBasicCredentials(clientId, *clientPassword)
				request.Header.Add("Authorization", "Basic "+encodedCredentials)
			}
			return sendTokenRequest(request)
		},
	}
}
