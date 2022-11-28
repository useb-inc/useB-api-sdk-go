package resources

import (
	"net/http"

	"github.com/useb-inc/useb-api-sdk-go/interfaces/status"
)

const (
	URLIdcard               string = "/status/idcard"
	URLDriver               string = "/status/driver"
	URLPassport             string = "/status/passport"
	URLPassportOverseas     string = "/status/passport-overseas"
	URLAlien                string = "/status/alien"
	URLBusinessRegistration string = "/status/business-registration"
)

type Status struct {
	Idcard               func(*status.IdcardRequest) (*status.IdcardResponse, error)
	Driver               func(*status.DriverRequest) (*status.DriverResponse, error)
	Passport             func(*status.PassportRequest) (*status.PassportResponse, error)
	PassportOverseas     func(*status.PassportOverseasRequest) (*status.PassportOverseasResponse, error)
	Alien                func(*status.AlienRequest) (*status.AlienResponse, error)
	BusinessRegistration func(*status.BusinessRegistrationRequest) (*status.BusinessRegistrationResponse, error)
}

func NewStatus(auth *Auth) *Status {
	base := &Base{
		Host: API_BASE_HOST,
	}
	idcard := callAPIMethod[status.IdcardRequest, status.IdcardResponse](base, auth, http.MethodPost, URLIdcard)
	driver := callAPIMethod[status.DriverRequest, status.DriverResponse](base, auth, http.MethodPost, URLDriver)
	passport := callAPIMethod[status.PassportRequest, status.PassportResponse](base, auth, http.MethodPost, URLPassport)
	passportOverseas := callAPIMethod[status.PassportOverseasRequest, status.PassportOverseasResponse](base, auth, http.MethodPost, URLPassportOverseas)
	alien := callAPIMethod[status.AlienRequest, status.AlienResponse](base, auth, http.MethodPost, URLAlien)
	businessRegistration := callAPIMethod[status.BusinessRegistrationRequest, status.BusinessRegistrationResponse](base, auth, http.MethodPost, URLBusinessRegistration)

	return &Status{
		Idcard:               idcard,
		Driver:               driver,
		Passport:             passport,
		PassportOverseas:     passportOverseas,
		Alien:                alien,
		BusinessRegistration: businessRegistration,
	}
}
