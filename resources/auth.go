package resources

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/auth"
)

const (
	URLRequestToken string = "/oauth/token"
)

type Token struct {
	jwt       string
	expiresIn time.Time
}

type Auth struct {
	token        *Token
	clientId     string
	clientSecret string
}

func NewAuth(clientId, clientSecret string) *Auth {
	auth := &Auth{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
	return auth
}

func (a *Auth) getToken() (string, error) {
	now := time.Now()
	// 토큰 캐시 리턴
	if a.token != nil && a.token.jwt != "" && !a.token.expiresIn.IsZero() && a.token.expiresIn.After(now) {
		return a.token.jwt, nil
	}

	token, err := requestToken(a.clientId, a.clientSecret)
	if err != nil {
		return "", err
	}
	// auth 인스턴스에 토큰 캐시
	a.token = token

	return token.jwt, nil
}

func requestToken(clientId, clientSecret string) (*Token, error) {
	urls := []string{AUTH_BASE_HOST, URLRequestToken}
	getTokenURL := strings.Join(urls, "")

	req, err := http.NewRequest(http.MethodPost, getTokenURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(clientId, clientSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(bytes))
	}

	var tokenRes auth.TokenResponse
	err = json.Unmarshal(bytes, &tokenRes)
	if err != nil {
		return nil, err
	}

	// YYYY-MM-DD HH:MM:SS 형식의 문자열을 받아서 Unix Timestamp로 변환
	expiresAt, err := time.Parse(DateLayoutISO, tokenRes.ExpiresIn)
	if err != nil {
		return nil, err
	}
	token := &Token{
		jwt:       tokenRes.Jwt,
		expiresIn: expiresAt,
	}

	return token, nil
}
