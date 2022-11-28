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
	AuthAPI            *resources.Auth
	StatusAPI          *resources.Status
	OpenbankAPI        *resources.Openbank
	FirmbankAPI        *resources.Firmbank
	FirmbankPremiumAPI *resources.FirmbankPremium
	OcrAPI             *resources.Ocr
	MaskingAPI         *resources.Masking
	FaceAPI            *resources.Face
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
	openbank := resources.NewOpenbank(auth)
	firmbank := resources.NewFirmbank(auth)
	firmbankpremium := resources.NewFirmbankPremium(auth)
	ocr := resources.NewOcr(auth)
	masking := resources.NewMasking(auth)
	face := resources.NewFace(auth)

	usebapi := &UsebAPI{
		AuthAPI:            auth,
		StatusAPI:          status,
		OpenbankAPI:        openbank,
		FirmbankAPI:        firmbank,
		FirmbankPremiumAPI: firmbankpremium,
		OcrAPI:             ocr,
		MaskingAPI:         masking,
		FaceAPI:            face,
	}

	return usebapi, nil
}
