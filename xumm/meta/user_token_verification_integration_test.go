//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestVerifyUserTokenIntegration(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	tks, err := m.VerifyUserToken("7e5d2547-4257-4487-afab-4a94bf07e92e")

	token := &models.UserTokenResponse{
		Tokens: []models.UserTokenValidity{
			{
				UserToken: "7e5d2547-4257-4487-afab-4a94bf07e92e",
				Active:    false,
				Issued:    0,
				Expires:   0,
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, token, tks)
}

func TestVerifyUserTokensIntegration_InvalidToken(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	tks, err := m.VerifyUserTokens("hello")

	token := &models.UserTokenResponse{Tokens: []models.UserTokenValidity{}}

	assert.NoError(t, err)
	assert.Equal(t, token, tks)
}

func TestVerifyUserTokensIntegration(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	tks, err := m.VerifyUserTokens("7e5d2547-4257-4487-afab-4a94bf07e92e")

	tokens := &models.UserTokenResponse{
		Tokens: []models.UserTokenValidity{
			{
				UserToken: "7e5d2547-4257-4487-afab-4a94bf07e92e",
				Active:    false,
				Issued:    0,
				Expires:   0,
			},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, tokens, tks)
}
