//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	statictestdata "github.com/xyield/xumm-go-client/xumm/meta/static-test-data"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetRailsIntegrationTest(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	r, err := m.GetRails()
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	assert.IsType(t, &models.RailsResponse{}, r)
	assert.Equal(t, statictestdata.RailsResponseResult, r)

}
