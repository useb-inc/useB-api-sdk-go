package status

type IdcardRequest struct {
	Identity  string `validate:"required" json:"identity"`
	IssueDate string `validate:"required" json:"issueDate"`
	UserName  string `validate:"required" json:"userName"`
	IsSecret  bool   `json:"secret_mode"`
}

type IdcardResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

type DriverRequest struct {
	UserName  string `validate:"required" json:"userName"`
	BirthDate string `validate:"required" json:"birthDate"`
	LicenseNo string `validate:"required" json:"licenseNo"`
	JuminNo   string `json:"juminNo"`
	IsSecret  bool   `json:"secret_mode"`
}
type DriverResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

type PassportRequest struct {
	UserName       string `validate:"required" json:"userName"`
	BirthDate      string `validate:"required" json:"birthDate"`
	PassportNo     string `validate:"required" json:"passportNo"`
	IssueDate      string `validate:"required" json:"issueDate"`
	ExpirationDate string `validate:"required" json:"expirationDate"`
	IsSecret       bool   `json:"secret_mode"`
}
type PassportResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

type PassportOverseasRequest struct {
	PassportNo  string `validate:"required" json:"passportNo"`
	Nationality string `validate:"required" json:"nationality"`
	BirthDate   string `validate:"required" json:"birthDate"`
	IsSecret    bool   `json:"secret_mode"`
}
type PassportOverseasResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

type AlienRequest struct {
	IssueNo   string `validate:"required" json:"issueNo"`
	IssueDate string `validate:"required" json:"issueDate"`
	IsSecret  bool   `json:"secret_mode"`
}
type AlienResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

type BusinessRegistrationRequest struct {
	BizNo string `validate:"required" json:"biz_no"`
}

type BusinessRegistrationData struct {
	TaxTypeCode       string `json:"tax_type_code"`
	TaxTypeName       string `json:"tax_type_name"`
	ClosingDate       string `json:"closing_date"`
	TaxTypeChangeDate string `json:"tax_type_change_date"`
}

type BusinessRegistrationResponse struct {
	Success       bool                      `json:"success"`
	Message       string                    `json:"message"`
	TransactionId string                    `json:"transaction_id"`
	Data          *BusinessRegistrationData `json:"data,omitempty"`
}
