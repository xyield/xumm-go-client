package models

type GetKycStatusByUserTokenResponse struct {
	KycStatus        string           `json:"kycStatus"`
	PossibleStatuses PossibleStatuses `json:"possibleStatuses"`
}

type PossibleStatuses struct {
	None       string `json:"NONE"`
	InProgress string `json:"IN_PROGRESS"`
	Rejected   string `json:"REJECTED"`
	Successful string `json:"SUCCESSFUL"`
}
