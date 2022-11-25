package auth

type TokenResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Jwt           string `json:"jwt"`
	ExpiresIn     string `json:"expires_in"`
	TransactionId string `json:"transaction_id"`
}
