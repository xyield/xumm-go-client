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
		Avatar: "https://xumm.app/cdn-cgi/image/width=250,height=250,quality=75,fit=crop/assets/icons/currencies/ex-bitstamp.png",
		Currencies: map[string]models.CurrencyInfo{
			"BTC": {
				Id:       492,
				IssuerId: 185,
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Currency: "BTC",
				Name:     "Bitcoin",
				Avatar:   "https://xumm.app/cdn-cgi/image/width=250,height=250,quality=75,fit=crop/assets/icons/currencies/crypto-btc.png",
			},
			"EUR": {
				Id:       13854758,
				IssuerId: 185,
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Currency: "EUR",
				Name:     "Euro",
				Avatar:   "https://xumm.app/cdn-cgi/image/width=250,height=250,quality=75,fit=crop/assets/icons/currencies/fiat-euro.png",
			},
			"GBP": {
				Id:       13854774,
				IssuerId: 185,
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Currency: "GBP",
				Name:     "British Pound",
				Avatar:   "https://cdn.xumm.pro/cdn-cgi/image/width=250,height=250,quality=75,fit=crop/currencies-tokens/fiat-brpnd.png",
			},
			"USD": {
				Id:       178,
				IssuerId: 185,
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Currency: "USD",
				Name:     "US Dollar",
				Avatar:   "https://xumm.app/cdn-cgi/image/width=250,height=250,quality=75,fit=crop/assets/icons/currencies/fiat-dollar.png",
			},
		},
	}, bitstamp)

}
