package models

type HookHashResponse struct {
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Creator          HookHashCreator `json:"creator,omitempty"`
	Xapp             string          `json:"xapp,omitempty"`
	AppUUID          string          `json:"appuuid,omitempty"`
	Icon             string          `json:"icon,omitempty"`
	VerifiedAccounts []string        `json:"verifiedAccounts,omitempty"`
	Audits           []string        `json:"audits,omitempty"`
}

type HookHashCreator struct {
	Name string `json:"name,omitempty"`
	Mail string `json:"mail,omitempty"`
	Site string `json:"site,omitempty"`
}
