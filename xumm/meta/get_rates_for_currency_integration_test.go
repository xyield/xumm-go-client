// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetRatesForCurrencyIntegration(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	r, err := m.GetRatesForCurrency("USD")
	assert.NoError(t, err)

	assert.Equal(t, models.Currency{
		En:     "US Dollar",
		Code:   "USD",
		Symbol: "$",
	}, r.Meta.Currency)
}
