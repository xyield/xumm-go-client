package models

import "github.com/xyield/xumm-go-client/pkg/json"

type XrpTxResponse struct {
	Txid           string                      `json:"txid"`
	BalanceChanges map[string][]BalanceDetails `json:"balanceChanges"`
	Node           string                      `json:"node"`
	Transaction    json.AnyJson                `json:"transaction"`
}

type BalanceDetails struct {
	CounterParty string `json:"counterParty,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Value        string `json:"value,omitempty"`
}
