package ocr

type IdcardRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	MaskMode    bool   `json:"mask_mode"`
	IsColor     bool   `json:"is_color"`
	IsSecret    bool   `json:"secret_mode"`
}
type IdcardData struct {
	IdType         string `json:"idType"`
	UserName       string `json:"userName"`
	JuminNo1       string `json:"juminNo1"`
	JuminNo2       string `json:"juminNo2"`
	MaskedJuminNo2 string `json:"_juminNo2"`
	IssueDate      string `json:"issueDate"`
}
type IdcardResponse struct {
	Success       bool        `json:"success"`
	Message       string      `json:"message"`
	TransactionId string      `json:"transaction_id"`
	Data          *IdcardData `json:"data"`
}

type DriverRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	MaskMode    bool   `json:"mask_mode"`
	IsColor     bool   `json:"is_color"`
	IsSecret    bool   `json:"secret_mode"`
}
type DriverData struct {
	IdType         string `json:"idType"`
	UserName       string `json:"userName"`
	DriverNo       string `json:"driverNo"`
	JuminNo1       string `json:"juminNo1"`
	JuminNo2       string `json:"juminNo2"`
	MaskedJuminNo2 string `json:"_juminNo2"`
	IssueDate      string `json:"issueDate"`
}
type DriverResponse struct {
	Success       bool        `json:"success"`
	Message       string      `json:"message"`
	TransactionId string      `json:"transaction_id"`
	Data          *DriverData `json:"data"`
}

type PassportRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	MaskMode    bool   `json:"mask_mode"`
	IsColor     bool   `json:"is_color"`
	IsSecret    bool   `json:"secret_mode"`
}
type PassportData struct {
	IdType         string `json:"idType"`
	UserName       string `json:"userName"`
	UserNameEng    string `json:"userNameEng"`
	PassportNo     string `json:"passportNo"`
	IssueNo1       string `json:"issueNo1"`
	IssueNo2       string `json:"issueNo2"`
	MaskedJuminNo2 string `json:"_juminNo2"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birthDate"`
	IssueDate      string `json:"issueDate"`
	ExpiryDate     string `json:"expiryDate"`
	Mrz1           string `json:"mrz1"`
	Mrz2           string `json:"mrz2"`
}
type PassportResponse struct {
	Success       bool          `json:"success"`
	Message       string        `json:"message"`
	TransactionId string        `json:"transaction_id"`
	Data          *PassportData `json:"data"`
}

type PassportOverseasRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	IsColor     bool   `json:"is_color"`
	IsSecret    bool   `json:"secret_mode"`
}
type PassportOverseasData struct {
	IdType      string `json:"idType"`
	UserName    string `json:"userName"`
	PassportNo  string `json:"passportNo"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birthDate"`
	ExpiryDate  string `json:"expiryDate"`
	Mrz1        string `json:"mrz1"`
	Mrz2        string `json:"mrz2"`
}
type PassportOverseasResponse struct {
	Success       bool                  `json:"success"`
	Message       string                `json:"message"`
	TransactionId string                `json:"transaction_id"`
	Data          *PassportOverseasData `json:"data"`
}

type AlienRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	MaskMode    bool   `json:"mask_mode"`
	IsColor     bool   `json:"is_color"`
	IsSecret    bool   `json:"secret_mode"`
}
type AlienData struct {
	IdType         string `json:"idType"`
	Country        string `json:"country"`
	CountryCode    string `json:"countryCode"`
	Visa           string `json:"visa"`
	UserName       string `json:"userName"`
	IssueNo        string `json:"issueNo"`
	IssueNo1       string `json:"issueNo1"`
	IssueNo2       string `json:"issueNo2"`
	MaskedJuminNo2 string `json:"_juminNo2"`
	IssueDate      string `json:"issueDate"`
}
type AlienResponse struct {
	Success       bool       `json:"success"`
	Message       string     `json:"message"`
	TransactionId string     `json:"transaction_id"`
	Data          *AlienData `json:"data"`
}

type BusinessRegistrationRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
}
type BusinessRegistrationData struct {
	DocSize         string `json:"docSize"`
	DocType         string `json:"docType"`
	CompanyName     string `json:"companyName"`
	OwnerName       string `json:"ownerName"`
	BusinessRegNum  string `json:"businessRegNum"`
	BusinessCorpNum string `json:"businessCorpNum"`
	CompanyAddr     string `json:"companyAddr"`
	HQAddr          string `json:"HQAddr"`
	OpenDate        string `json:"openDate"`
	BusinessType1   string `json:"businessType1"`
	BusinessType2   string `json:"businessType2"`
}
type BusinessRegistrationResponse struct {
	Success       bool                      `json:"success"`
	Message       string                    `json:"message"`
	TransactionId string                    `json:"transaction_id"`
	Data          *BusinessRegistrationData `json:"data"`
}
