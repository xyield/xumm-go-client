package models

import anyjson "github.com/xyield/xumm-go-client/utils/json"

type XrpTxResponse struct {
	Txid           string                      `json:"txid"`
	BalanceChanges map[string][]BalanceDetails `json:"balanceChanges"`
	Node           string                      `json:"node"`
	Transaction    anyjson.AnyJson             `json:"transaction"`
}

type BalanceDetails struct {
	CounterParty string    `json:"counterParty"`
	Currency     string    `json:"currency"`
	Value        string    `json:"value"`
	Formatted    Formatted `json:"formatted"`
}

type Formatted struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}
