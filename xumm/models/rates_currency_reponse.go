package models

type RatesCurrencyResponse struct {
	Usd  float64 `json:"USD"`
	Xrp  float64 `json:"XRP"`
	Meta Meta    `json:"__meta"`
}

type Meta struct {
	Currency Currency `json:"currency"`
}

type Currency struct {
	En     string `json:"en"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
}
