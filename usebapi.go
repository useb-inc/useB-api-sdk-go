package usebapi

import (
	"errors"

	"github.com/useb-inc/useb-api-sdk-go/resources"
)

const (
	ErrClientIDRequired     string = "usebapi: Client ID is required"
	ErrClientSecretRequired string = "usebapi: Client Secret is required"
)

type UsebAPI struct {
	AuthAPI   *resources.Auth
	StatusAPI *resources.Status
}

func NewUsebAPI(clientId, clientSecret string) (*UsebAPI, error) {
	if clientId == "" {
		return nil, errors.New(ErrClientIDRequired)
	}

	if clientSecret == "" {
		return nil, errors.New(ErrClientSecretRequired)
	}

	auth := resources.NewAuth(clientId, clientSecret)
	status := resources.NewStatus(auth)

	usebapi := &UsebAPI{
		AuthAPI:   auth,
		StatusAPI: status,
	}

	return usebapi, nil
}
