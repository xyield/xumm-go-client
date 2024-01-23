//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetHookHashIntegrationTest(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	hookHash := "31C3EC186C367DA66DFBD0E576D6170A2C1AB846BFC35FC0B49D202F2A8CDFD8"

	hh, err := m.GetHookHash(hookHash)

	assert.NoError(t, err)
	assert.IsType(t, &models.HookHashResponse{}, hh)
	assert.NotEmpty(t, hh.Name)
	assert.NotEmpty(t, hh.Description)
}
