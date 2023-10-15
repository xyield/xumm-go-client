package models

type XappOttResponse struct {
	Locale        string      `json:"locale"`
	Version       string      `json:"version"`
	Account       string      `json:"account"`
	Accountaccess string      `json:"accountaccess"`
	Accounttype   string      `json:"accounttype"`
	Style         string      `json:"style"`
	Origin        Origin      `json:"origin"`
	User          string      `json:"user"`
	UserDevice    UserDevice  `json:"user_device"`
	AccountInfo   AccountInfo `json:"account_info"`
}

type Origin struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	Txid string `json:"txid"`
}

type UserDevice struct {
	Currency string `json:"currency"`
}

type AccountInfo struct {
	Account         string  `json:"account"`
	Name            string  `json:"name"`
	Domain          string  `json:"domain"`
	Blocked         bool    `json:"blocked"`
	Source          string  `json:"source"`
	KycApproved     bool    `json:"kycApproved"`
	ProSubscription bool    `json:"proSubscription"`
	Profile         Profile `json:"profile"`
}

type Profile struct {
	Slug        string `json:"slug"`
	ProfileURL  string `json:"profileUrl"`
	AccountSlug string `json:"accountSlug"`
	PayString   string `json:"payString"`
}
