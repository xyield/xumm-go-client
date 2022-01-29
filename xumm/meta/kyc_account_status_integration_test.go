//go:build integration
// +build integration

package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetKycStatusByAccountIntegration(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	ky, err := m.GetKycStatusByAccount("r9fVvKgMMPkBHQ3n28sifxi22zKphwkf8u")
	assert.NoError(t, err)

	assert.Equal(t, &models.KycStatusByAccountResponse{
		Account:     "r9fVvKgMMPkBHQ3n28sifxi22zKphwkf8u",
		KycApproved: false,
	}, ky)
}

func TestGetKycStatusByUserTokenIntegration(t *testing.T) {

	cfg, err := xumm.NewConfig()
	assert.NoError(t, err)
	m := &Meta{
		Cfg: cfg,
	}

	ky, err := m.GetKycStatusByUserToken(models.KycStatusByUserTokenRequest{"NoKYCAttemptMadeUserToken"})
	assert.NoError(t, err)

	assert.Equal(t, &models.KycStatusByUserTokenResponse{
		KycStatus: "NONE",
		PossibleStatuses: models.PossibleStatuses{
			None: "No KYC attempt has been made", InProgress: "KYC flow has been started, but did not finish (yet)", Rejected: "KYC flow has been started and rejected (NO SUCCESSFUL KYC)", Successful: "KYC flow has been started and was SUCCESSFUL :)"}}, ky)
}
