package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/masking"
)

const (
	URLMaskingIdcard   string = "/masking/idcard"
	URLMaskingDriver   string = "/masking/driver"
	URLMaskingPassport string = "/masking/passport"
	URLMaskingAlien    string = "/masking/alien"
)

type Masking struct {
	Idcard   func(*masking.IdcardRequest) (*masking.IdcardResponse, error)
	Driver   func(*masking.DriverRequest) (*masking.DriverResponse, error)
	Passport func(*masking.PassportRequest) (*masking.PassportResponse, error)
	Alien    func(*masking.AlienRequest) (*masking.AlienResponse, error)
}

func NewMasking(auth *Auth) *Masking {
	base := &Base{
		Host: API_BASE_HOST,
	}
	idcard := callAPIMethod[masking.IdcardRequest, masking.IdcardResponse](base, auth, http.MethodPost, URLOcrIdcard)
	driver := callAPIMethod[masking.DriverRequest, masking.DriverResponse](base, auth, http.MethodPost, URLOcrDriver)
	passport := callAPIMethod[masking.PassportRequest, masking.PassportResponse](base, auth, http.MethodPost, URLOcrPassport)
	alien := callAPIMethod[masking.AlienRequest, masking.AlienResponse](base, auth, http.MethodPost, URLOcrAlien)

	return &Masking{
		Idcard:   idcard,
		Driver:   driver,
		Passport: passport,
		Alien:    alien,
	}
}
