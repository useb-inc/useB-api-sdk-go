package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/openbank"
)

const (
	URLOpenbankRealname string = "/realname"
	URLOpenbankSend     string = "/send"
	URLOpenbankVerify   string = "/verify"
)

type Openbank struct {
	Realname func(*openbank.RealnameRequest) (*openbank.RealnameResponse, error)
	Send     func(*openbank.SendRequest) (*openbank.SendResponse, error)
	Verify   func(*openbank.VerifyRequest) (*openbank.VerifyResponse, error)
}

func NewOpenbank(auth *Auth) *Openbank {
	base := &Base{
		Host: BANKING_BASE_HOST,
	}
	realname := callAPIMethod[openbank.RealnameRequest, openbank.RealnameResponse](base, auth, http.MethodPost, URLOpenbankRealname)
	send := callAPIMethod[openbank.SendRequest, openbank.SendResponse](base, auth, http.MethodPost, URLOpenbankSend)
	verify := callAPIMethod[openbank.VerifyRequest, openbank.VerifyResponse](base, auth, http.MethodPost, URLOpenbankRealname)

	return &Openbank{
		Realname: realname,
		Send:     send,
		Verify:   verify,
	}
}
