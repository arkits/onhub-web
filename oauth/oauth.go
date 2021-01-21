package oauth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var logger = logging.MustGetLogger("oauth")

const (
	OAUTH_ACCESS_TOKEN_URL = "https://www.googleapis.com/oauth2/v4/token"
	OAUTH_ISSUE_TOKEN_URL  = "https://oauthaccountmanager.googleapis.com/v1/issuetoken"
)

var httpClient = &http.Client{}

type PersistedToken struct {
	Token      string
	ExpiryDate time.Time
}

type TokenStore struct {
	mu    sync.Mutex
	Token PersistedToken
}

var tokenStore TokenStore

// GetToken return a Bearer Token that can be used to make OnHub API calls
func GetToken() string {

	if !isTokenCurrent(tokenStore.Token) {
		accessToken := GetAccessToken()
		persistedToken := GetIssuedToken(accessToken)

		tokenStore.mu.Lock()
		tokenStore.Token = persistedToken
		tokenStore.mu.Unlock()

		logger.Debug("Persisted new API Token")
	} else {
		// logger.Debug("Using Cached API Token")
	}

	// Token may have been updated... Get the new one
	token := tokenStore.Token

	return token.Token
}

// GetAccessToken gets the initial "access_token"
func GetAccessToken() string {

	// Build the payload
	payload := url.Values{}
	payload.Set("client_id", "936475272427.apps.googleusercontent.com")
	payload.Set("grant_type", "refresh_token")
	payload.Set("refresh_token", viper.GetString("auth.refresh_token"))

	request, err := http.NewRequest("POST", OAUTH_ACCESS_TOKEN_URL, strings.NewReader(payload.Encode()))
	if err != nil {
		logger.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := httpClient.Do(request)
	if err != nil {
		logger.Fatal(err)
	}

	defer response.Body.Close()

	type AccessTokenReponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
		IDToken     string `json:"id_token"`
	}

	var accessToken AccessTokenReponse
	json.NewDecoder(response.Body).Decode(&accessToken)

	return accessToken.AccessToken
}

// GetIssuedToken gets the issue API Token based on the accessToken
func GetIssuedToken(accessToken string) PersistedToken {

	// Build the payload
	payload := url.Values{}
	payload.Set("client_id", "586698244315-vc96jg3mn4nap78iir799fc2ll3rk18s.apps.googleusercontent.com")
	payload.Set("app_id", "com.google.OnHub")
	payload.Set("hl", "en-US")
	payload.Set("lib_ver", "3.3")
	payload.Set("response_type", "token")
	payload.Set("scope", "https://www.googleapis.com/auth/accesspoints https://www.googleapis.com/auth/clouddevices")

	request, err := http.NewRequest("POST", OAUTH_ISSUE_TOKEN_URL, strings.NewReader(payload.Encode()))
	if err != nil {
		logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		logger.Fatal(err)
	}

	defer response.Body.Close()

	type IssueTokenResponse struct {
		IssueAdvice   string `json:"issueAdvice"`
		Token         string `json:"token"`
		ExpiresIn     string `json:"expiresIn"`
		GrantedScopes string `json:"grantedScopes"`
	}

	var issuedToken IssueTokenResponse
	json.NewDecoder(response.Body).Decode(&issuedToken)

	var persistedToken PersistedToken
	persistedToken.Token = issuedToken.Token

	expiresIn, _ := strconv.Atoi(issuedToken.ExpiresIn)
	expiresIn = expiresIn - 5 // Reduce the expiry time... just in case :)
	persistedToken.ExpiryDate = time.Now().Add(time.Duration(expiresIn) * time.Second)

	return persistedToken
}

func isTokenCurrent(token PersistedToken) bool {

	if token.Token == "" {
		logger.Debug("isTokenCurrent: Token is not set")
		return false
	}

	timeNow := time.Now()
	expiry := token.ExpiryDate

	if timeNow.After(expiry) {
		diff := expiry.Sub(timeNow) / time.Nanosecond
		logger.Debugf("PersistedToken is expiring in %v", diff)
		return false
	}

	return true

}
