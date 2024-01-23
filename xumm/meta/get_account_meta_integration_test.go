package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	statictestdata "github.com/xyield/xumm-go-client/xumm/meta/static-test-data"
)

func TestGetAccountMetaIntegrationTest(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	am, err := m.GetAccountMeta(statictestdata.AccountMetaTestResult.Account)

	assert.NoError(t, err)
	assert.NotEmpty(t, am)
	assert.Equal(t, statictestdata.AccountMetaTestResult, am)
}
