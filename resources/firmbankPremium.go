package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/firmbankpremium"
)

const (
	URLFirmbankPremiumRealname string = "/firmbank-custom/realname"
	URLFirmbankPremiumSend     string = "/firmbank-custom/send"
	URLFirmbankPremiumVerify   string = "/firmbank-custom/verify"
)

type FirmbankPremium struct {
	Realname func(*firmbankpremium.RealnameRequest) (*firmbankpremium.RealnameResponse, error)
	Send     func(*firmbankpremium.SendRequest) (*firmbankpremium.SendResponse, error)
	Verify   func(*firmbankpremium.VerifyRequest) (*firmbankpremium.VerifyResponse, error)
}

func NewFirmbankPremium(auth *Auth) *FirmbankPremium {
	base := &Base{
		Host: BANKING_BASE_HOST,
	}
	realname := callAPIMethod[firmbankpremium.RealnameRequest, firmbankpremium.RealnameResponse](base, auth, http.MethodPost, URLFirmbankPremiumRealname)
	send := callAPIMethod[firmbankpremium.SendRequest, firmbankpremium.SendResponse](base, auth, http.MethodPost, URLFirmbankPremiumSend)
	verify := callAPIMethod[firmbankpremium.VerifyRequest, firmbankpremium.VerifyResponse](base, auth, http.MethodPost, URLFirmbankPremiumVerify)

	return &FirmbankPremium{
		Realname: realname,
		Send:     send,
		Verify:   verify,
	}
}
