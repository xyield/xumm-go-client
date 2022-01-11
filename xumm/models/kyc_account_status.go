package models

// by xrp address
type KycAccountStatusResponse struct {
	Account     string `json:"account"`
	KycApproved bool   `json:"kycApproved"`
}
