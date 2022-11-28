package firmbank

type SendRequest struct {
	BankCode          string `validate:"required" json:"bank_code"`
	AccountNum        string `validate:"required" json:"account_num"`
	AccountHolderName string `validate:"required" json:"account_holder_name"`
	IsSecret          bool   `json:"secret_mode"`
}
type SendData struct {
	Tid string `json:"tid"`
}
type SendResponse struct {
	Success       bool      `json:"success"`
	Message       string    `json:"message"`
	TransactionId string    `json:"transaction_id"`
	Data          *SendData `json:"data,omitempty"`
}

type VerifyRequest struct {
	Tid          string `validate:"required" json:"tid"`
	PrintContent string `validate:"required" json:"print_content"`
}
type VerifyResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}
