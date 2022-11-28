package face

type VerifyRequest struct {
	Face1 string `validate:"required" json:"face1"`
	Face2 string `validate:"required" json:"face2"`
}
type VerifyResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
	IsIdentical   bool   `json:"isIdentical"`
	Confidence    string `json:"confidence"`
}
