package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/status"
)

const (
	URLIdcard string = "/status/idcard"
)

type Status struct {
	Idcard func(*status.IdcardRequest) (*status.IdcardResponse, error)
}

func NewStatus(auth *Auth) *Status {
	base := &Base{
		Host: API_BASE_HOST,
	}
	idcard := callAPIMethod[status.IdcardRequest, status.IdcardResponse](base, auth, http.MethodPost, URLIdcard)

	return &Status{
		Idcard: idcard,
	}
}
