package models

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Transaction map[string]interface{}

type XrpTxResponse struct {
	Txid           string                      `json:"txid"`
	BalanceChanges map[string][]BalanceDetails `json:"balanceChanges"`
	Node           string                      `json:"node"`
	Transaction    Transaction                 `json:"transaction"`
}

type BalanceDetails struct {
	CounterParty string `json:"counterParty,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Value        string `json:"value,omitempty"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {

	if *t == nil {
		*t = make(Transaction)
	}

	var values map[string]interface{}
	if err := jsoniter.Unmarshal(data, &values); err != nil {
		return err
	}
	for k, v := range values {
		fmt.Printf("%v: %v, %T\n", k, v, v)
		if i, ok := v.(float64); ok {
			values[k] = int64(i)
		}
		// if k == "meta" {
		// 	for mk, mv := range values["meta"] {
		// 		if j, ok := mv.(float64); ok {
		// 			values["meta"][mk] = int64(j)
		// 		}
		// 	}
		// }
		(*t)[k] = values[k]
	}

	return nil
}