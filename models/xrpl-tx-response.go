package models

type XrpTxResponse struct {
	Txid           string                      `json:"txid"`
	BalanceChanges map[string][]BalanceDetails `json:"balanceChanges"`
	Node           string                      `json:"node"`
	Transaction    map[string]interface{}      `json:"transaction"`
}

type BalanceDetails struct {
	CounterParty string `json:"counterParty,omitempty"`
	Cuurency     string `json:"currency,omitempty"`
	Value        string `json:"value,omitempty"`
}
