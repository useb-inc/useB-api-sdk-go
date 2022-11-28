package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/firmbank"
)

const (
	URLFirmbankSend   string = "/firmbank/send"
	URLFirmbankVerify string = "/firmbank/verify"
)

type Firmbank struct {
	Send   func(*firmbank.SendRequest) (*firmbank.SendResponse, error)
	Verify func(*firmbank.VerifyRequest) (*firmbank.VerifyResponse, error)
}

func NewFirmbank(auth *Auth) *Firmbank {
	base := &Base{
		Host: BANKING_BASE_HOST,
	}
	send := callAPIMethod[firmbank.SendRequest, firmbank.SendResponse](base, auth, http.MethodPost, URLFirmbankSend)
	verify := callAPIMethod[firmbank.VerifyRequest, firmbank.VerifyResponse](base, auth, http.MethodPost, URLFirmbankVerify)

	return &Firmbank{
		Send:   send,
		Verify: verify,
	}
}
