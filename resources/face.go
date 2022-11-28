package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/face"
)

const (
	URLFaceVerify string = "/face/verify"
)

type Face struct {
	Verify func(*face.VerifyRequest) (*face.VerifyResponse, error)
}

func NewFace(auth *Auth) *Face {
	base := &Base{
		Host: API_BASE_HOST,
	}
	verify := callAPIMethod[face.VerifyRequest, face.VerifyResponse](base, auth, http.MethodPost, URLFaceVerify)

	return &Face{
		Verify: verify,
	}
}
