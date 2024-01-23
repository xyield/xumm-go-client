//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetHookHashesIntegrationTest(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	hh, err := m.GetHookHashes()

	assert.NoError(t, err)
	assert.IsType(t, &models.HookHashesResponse{}, hh)
	assert.NotEmpty(t, hh)

	for _, hookHash := range *hh {
		assert.NotEmpty(t, hookHash.Name)
		assert.NotEmpty(t, hookHash.Description)
	}

}
