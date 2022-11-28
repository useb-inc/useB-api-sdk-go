package masking

type IdcardRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	IsSecret    bool   `json:"secret_mode"`
}
type IdcardData struct {
	ImageBase64Mask string `json:"image_base64_mask"`
}
type IdcardResponse struct {
	Success       bool        `json:"success"`
	Message       string      `json:"message"`
	TransactionId string      `json:"transaction_id"`
	Data          *IdcardData `json:"data"`
}

type DriverRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	IsSecret    bool   `json:"secret_mode"`
}
type DriverData struct {
	ImageBase64Mask string `json:"image_base64_mask"`
}
type DriverResponse struct {
	Success       bool        `json:"success"`
	Message       string      `json:"message"`
	TransactionId string      `json:"transaction_id"`
	Data          *DriverData `json:"data"`
}

type PassportRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	IsSecret    bool   `json:"secret_mode"`
}
type PassportData struct {
	ImageBase64Mask string `json:"image_base64_mask"`
}
type PassportResponse struct {
	Success       bool          `json:"success"`
	Message       string        `json:"message"`
	TransactionId string        `json:"transaction_id"`
	Data          *PassportData `json:"data"`
}

type AlienRequest struct {
	ImageBase64 string `validate:"required" json:"image_base64"`
	IsSecret    bool   `json:"secret_mode"`
}
type AlienData struct {
	ImageBase64Mask string `json:"image_base64_mask"`
}
type AlienResponse struct {
	Success       bool       `json:"success"`
	Message       string     `json:"message"`
	TransactionId string     `json:"transaction_id"`
	Data          *AlienData `json:"data"`
}
