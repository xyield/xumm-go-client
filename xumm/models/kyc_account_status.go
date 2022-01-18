package models

type KycStatusByAccountResponse struct {
	Account     string `json:"account"`
	KycApproved bool   `json:"kycApproved"`
}
