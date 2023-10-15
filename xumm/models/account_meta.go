package models

type AccountMetaResponse struct {
	Avatar             string              `json:"avatar,omitempty"`
	XummProfile        XummProfile         `json:"xummProfile,omitempty"`
	ThirdPartyProfiles []ThirdPartyProfile `json:"thirdPartyProfiles,omitempty"`
	GlobalID           GlobalID            `json:"globalid,omitempty"`
}

type XummProfile struct {
	AccountAlias string `json:"accountAlias,omitempty"`
	OwnerAlias   string `json:"ownerAlias,omitempty"`
}

type ThirdPartyProfile struct {
	AccountAlias string `json:"accountAlias,omitempty"`
	Source       string `json:"source,omitempty"`
}

type GlobalID struct {
	Linked     string `json:"linked,omitempty"`
	ProfileURL string `json:"profileUrl,omitempty"`
}
