package models

type NftokenDetail struct {
	Issuer string `json:"issuer,omitempty"`
	Token  string `json:"token,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Name   string `json:"name,omitempty"`
	Image  string `json:"image,omitempty"`
}
