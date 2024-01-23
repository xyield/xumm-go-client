package models

type AccountMetaResponse struct {
	Account            string              `json:"account"`
	KycApproved        bool                `json:"kycApproved"`
	XummPro            bool                `json:"xummPro"`
	Blocked            bool                `json:"blocked"`
	ForceDtag          bool                `json:"force_dtag"`
	Avatar             string              `json:"avatar,omitempty"`
	XummProfile        XummProfile         `json:"xummProfile,omitempty"`
	ThirdPartyProfiles []ThirdPartyProfile `json:"thirdPartyProfiles,omitempty"`
	GlobalID           GlobalID            `json:"globalid,omitempty"`
}

type XummProfile struct {
	AccountAlias string `json:"accountAlias,omitempty"`
	OwnerAlias   string `json:"ownerAlias,omitempty"`
	Slug         string `json:"slug,omitempty"`
	ProfileURL   string `json:"profileUrl,omitempty"`
	AccountSlug  string `json:"accountSlug,omitempty"`
	PayString    string `json:"payString,omitempty"`
}

type ThirdPartyProfile struct {
	AccountAlias string `json:"accountAlias,omitempty"`
	Source       string `json:"source,omitempty"`
}

type GlobalID struct {
	Linked          string `json:"linked,omitempty"`
	ProfileURL      string `json:"profileUrl,omitempty"`
	SufficientTrust bool   `json:"sufficientTrust"`
}
