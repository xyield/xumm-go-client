//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestPingIntegration(t *testing.T) {
	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)

	m := &Meta{Cfg: cfg}

	pong, err := m.Ping()

	assert.NoError(t, err)

	assert.Equal(t, models.Application{
		UUIDV4:     "eda1fbb4-8641-47bd-91c8-3adca27cd6e3",
		Name:       "Test Xumm App",
		WebhookUrl: "https://32fc-115-189-103-75.ngrok-free.app/webhook",
		Disabled:   0,
	}, pong.Auth.Application)
	assert.True(t, pong.Pong)
	assert.True(t, testutils.IsValidUUID(pong.Auth.Call.UUIDV4))
}

func TestPingErrorIntegration(t *testing.T) {
	cfg, err := xumm.NewConfig(xumm.WithAuth("badKey", "badSecret"))
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	pong, err := m.Ping()
	assert.Nil(t, pong)
	assert.Error(t, err)
	assert.True(t, testutils.IsValidUUID(err.(*xumm.ErrorResponse).ErrorResponseBody.Reference))
}
