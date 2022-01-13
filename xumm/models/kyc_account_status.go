package models

type KycAccountStatusResponse struct {
	Account     string `json:"account"`
	KycApproved bool   `json:"kycApproved"`
}
