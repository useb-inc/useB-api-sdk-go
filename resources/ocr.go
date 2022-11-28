package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/ocr"
)

const (
	URLOcrIdcard               string = "/ocr/idcard-driver"
	URLOcrDriver               string = "/ocr/idcard-driver"
	URLOcrPassport             string = "/ocr/passport"
	URLOcrPassportOverseas     string = "/ocr/passport-overseas"
	URLOcrAlien                string = "/ocr/alien"
	URLOcrBusinessRegistration string = "/ocr-doc/business-registration"
)

type Ocr struct {
	Idcard               func(*ocr.IdcardRequest) (*ocr.IdcardResponse, error)
	Driver               func(*ocr.DriverRequest) (*ocr.DriverResponse, error)
	Passport             func(*ocr.PassportRequest) (*ocr.PassportResponse, error)
	PassportOverseas     func(*ocr.PassportOverseasRequest) (*ocr.PassportOverseasResponse, error)
	Alien                func(*ocr.AlienRequest) (*ocr.AlienResponse, error)
	BusinessRegistration func(*ocr.BusinessRegistrationRequest) (*ocr.BusinessRegistrationResponse, error)
}

func NewOcr(auth *Auth) *Ocr {
	base := &Base{
		Host: API_BASE_HOST,
	}
	idcard := callAPIMethod[ocr.IdcardRequest, ocr.IdcardResponse](base, auth, http.MethodPost, URLOcrIdcard)
	driver := callAPIMethod[ocr.DriverRequest, ocr.DriverResponse](base, auth, http.MethodPost, URLOcrDriver)
	passport := callAPIMethod[ocr.PassportRequest, ocr.PassportResponse](base, auth, http.MethodPost, URLOcrPassport)
	passportOverseas := callAPIMethod[ocr.PassportOverseasRequest, ocr.PassportOverseasResponse](base, auth, http.MethodPost, URLOcrPassportOverseas)
	alien := callAPIMethod[ocr.AlienRequest, ocr.AlienResponse](base, auth, http.MethodPost, URLOcrAlien)
	businessRegistration := callAPIMethod[ocr.BusinessRegistrationRequest, ocr.BusinessRegistrationResponse](base, auth, http.MethodPost, URLOcrBusinessRegistration)

	return &Ocr{
		Idcard:               idcard,
		Driver:               driver,
		Passport:             passport,
		PassportOverseas:     passportOverseas,
		Alien:                alien,
		BusinessRegistration: businessRegistration,
	}
}
