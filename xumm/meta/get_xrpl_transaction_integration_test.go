//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestXrplTransactionIntegrationTest(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	txid := "EA043BA795908EC70C93F65FB68700069BEA2DF7152BC7D5F2D8F3261A93509B"
	bc := map[string][]models.BalanceDetails{
		"rUjKdoe5XrHcbf51PHqnv3d7Q5aNiP2KC5": {
			models.BalanceDetails{
				CounterParty: "",
				Currency:     "XRP",
				Value:        "-0.589015",
				Formatted: models.Formatted{
					Value:    "-0.589015",
					Currency: "XRP",
				},
			},
		},
		"r9fVvKgMMPkBHQ3n28sifxi22zKphwkf8u": {
			models.BalanceDetails{
				CounterParty: "",
				Currency:     "XRP",
				Value:        "0.589",
				Formatted: models.Formatted{
					Value:    "0.589",
					Currency: "XRP",
				},
			},
		},
	}

	xt, err := m.GetXrplTransaction(txid)

	assert.NoError(t, err)
	assert.Equal(t, "EA043BA795908EC70C93F65FB68700069BEA2DF7152BC7D5F2D8F3261A93509B", xt.Txid)
	assert.Equal(t, bc, xt.BalanceChanges)
	assert.Equal(t, xt.Node, "wss://xrplcluster.com")
	assert.NotNil(t, xt.Transaction)

	assert.Contains(t, xt.Transaction, "Account")
	assert.Contains(t, xt.Transaction, "Amount")
	assert.Contains(t, xt.Transaction, "Destination")
	assert.Contains(t, xt.Transaction, "Fee")
	assert.Contains(t, xt.Transaction, "TransactionType")
}
