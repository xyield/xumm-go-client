package models

type XappResponse struct {
	Locale        string     `json:"locale"`
	Version       string     `json:"version"`
	Account       string     `json:"account"`
	Accountaccess string     `json:"accountaccess"`
	Accounttype   string     `json:"accounttype"`
	Style         string     `json:"style"`
	Origin        Origin     `json:"origin"`
	User          string     `json:"user"`
	UserDevice    UserDevice `json:"user_device"`
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
