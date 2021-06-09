package models

type CurratedAssetsResponse struct {
	Issuers    []string          `json:"issuers"`
	Currencies []string          `json:"currencies"`
	Details    map[string]Issuer `json:"details"`
}

type Issuer struct {
	Id         int64                   `json:"id"`
	Name       string                  `json:"name"`
	Domain     string                  `json:"domain,omitempty"`
	Avatar     string                  `json:"avatar,omitempty"`
	Currencies map[string]CurrencyInfo `json:"currencies"`
}

type CurrencyInfo struct {
	Id       int64  `json:"id"`
	IssuerId int64  `json:"issuer_id"`
	Issuer   string `json:"issuer"`
	Currency string `json:"currency"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar,omitempty"`
}
