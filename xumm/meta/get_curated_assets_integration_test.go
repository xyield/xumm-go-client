//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestCuratedAssetsIntegration(t *testing.T) {
	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	ca, err := m.GetCuratedAssets()

	assert.NoError(t, err)
	assert.NotNil(t, ca)
	assert.Contains(t, ca.Issuers, "Bitstamp")
	assert.Contains(t, ca.Currencies, "BTC")

	bitstamp, ok := ca.Details["Bitstamp"]

	assert.True(t, ok)

	assert.Equal(t, models.Issuer{
		Id:     185,
		Name:   "Bitstamp",
		Domain: "bitstamp.net",
		Avatar: "https://xumm.app/assets/icons/currencies/ex-bitstamp.png",
		Currencies: map[string]models.CurrencyInfo{
			"BTC": {
				Id:       492,
				IssuerId: 185,
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Currency: "BTC",
				Name:     "Bitcoin",
				Avatar:   "https://xumm.app/assets/icons/currencies/crypto-btc.png",
			},
			"USD": {
				Id:       178,
				IssuerId: 185,
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Currency: "USD",
				Name:     "US Dollar",
				Avatar:   "https://xumm.app/assets/icons/currencies/fiat-dollar.png",
			},
		},
	}, bitstamp)

}
