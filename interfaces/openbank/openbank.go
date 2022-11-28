package openbank

// [오픈뱅킹 API]
// 계좌실명조회
type RealnameRequest struct {
	BankCode              string `validate:"required" json:"bank_code"`
	AccountNum            string `validate:"required" json:"account_num"`
	AccountHolderInfoType string `validate:"required" json:"account_holder_info_type"`
	AccountHolderInfo     string `validate:"required" json:"account_holder_info"`
	IsSecret              bool   `json:"secret_mode"`
}
type RealnameData struct {
	BankCode          string `json:"bank_code"`
	BankName          string `json:"bank_name"`
	AccountNum        string `json:"account_num"`
	AccountHolderName string `json:"account_holder_name"`
}
type RealnameResponse struct {
	Success       bool          `json:"success"`
	Message       string        `json:"message"`
	TransactionId string        `json:"transaction_id"`
	ExtraMessage  string        `json:"extra_message,omitempty"`
	Data          *RealnameData `json:"data,omitempty"`
}

// 1원 입금이체
type SendRequest struct {
	BankCode          string `validate:"required" json:"bank_code"`
	AccountNum        string `validate:"required" json:"account_num"`
	AccountHolderName string `validate:"required" json:"account_holder_name"`
	CodeType          string `validate:"required" json:"code_type"`
	ClientName        string `json:"client_name"`
	ClientBusinessNum string `json:"client_business_num"`
	CodePosition      string `json:"code_position"`
	IsSecret          bool   `json:"secret_mode"`
}
type SendResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

// 1원인증코드검증
type VerifyRequest struct {
	TransactionId string `validate:"required" json:"transaction_id"`
	PrintContent  string `validate:"required" json:"print_content"`
}
type VerifyData struct {
	PairTransactionId string `json:"pair_transaction_id"`
	PrintContent      string `json:"print_content"`
}
type VerifyResponse struct {
	Success       bool        `json:"success"`
	Message       string      `json:"message"`
	TransactionId string      `json:"transaction_id"`
	Data          *VerifyData `json:"data,omitempty"`
}
