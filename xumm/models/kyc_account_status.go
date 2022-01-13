package models

type GetKycStatusByAccountResponse struct {
	Account     string `json:"account"`
	KycApproved bool   `json:"kycApproved"`
}
