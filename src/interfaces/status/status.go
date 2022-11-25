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
