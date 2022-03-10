// +build unit

package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetKycStatusByAccountTest(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		json           string
		expectedOutput *models.KycStatusByAccountResponse
	}{
		{
			description: "Valid account with kyc status true",
			input:       "rGBP1ZYpgiArYbDSvqu7Ps8AmWrD6hiqwe",
			json: `{
				"account": "rGBP1ZYpgiArYbDSvqu7Ps8AmWrD6hiqwe",
				"kycApproved": true
			  }`,
			expectedOutput: &models.KycStatusByAccountResponse{
				Account:     "rGBP1ZYpgiArYbDSvqu7Ps8AmWrD6hiqwe",
				KycApproved: true,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(test.json, 200, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}

			customer, _ := meta.GetKycStatusByAccount(test.input)
			assert.Equal(t, http.Header{
				"X-API-Key":    {"testApiKey"},
				"X-API-Secret": {"testApiSecret"},
				"Content-Type": {"application/json"},
			}, m.Spy.Header)
			assert.Equal(t, test.expectedOutput, customer)
		})
	}
}

func TestGetKycStatusByUserToken(t *testing.T) {
	tt := []struct {
		description    string
		input          models.KycStatusByUserTokenRequest
		json           string
		expectedOutput *models.KycStatusByUserTokenResponse
	}{
		{
			description: "Valid account with kyc status none",
			input: models.KycStatusByUserTokenRequest{
				UserToken: "test-token",
			},
			json: `{
				"kycStatus": "NONE",
				"possibleStatuses": {
				  "NONE": "No KYC attempt has been made",
				  "IN_PROGRESS": "KYC flow has been started, but did not finish (yet)",
				  "REJECTED": "KYC flow has been started and rejected (NO SUCCESSFUL KYC)",
				  "SUCCESSFUL": "KYC flow has been started and was SUCCESSFUL :)"
				}
			  }`,
			expectedOutput: &models.KycStatusByUserTokenResponse{
				KycStatus: "NONE",
				PossibleStatuses: models.PossibleStatuses{
					None:       "No KYC attempt has been made",
					InProgress: "KYC flow has been started, but did not finish (yet)",
					Rejected:   "KYC flow has been started and rejected (NO SUCCESSFUL KYC)",
					Successful: "KYC flow has been started and was SUCCESSFUL :)",
				},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(test.json, 200, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}
			customer, _ := meta.GetKycStatusByUserToken(test.input)
			assert.Equal(t, test.expectedOutput, customer)
			assert.Equal(t, http.Header{
				"X-API-Key":    {"testApiKey"},
				"X-API-Secret": {"testApiSecret"},
				"Content-Type": {"application/json"},
			}, m.Spy.Header)
		})
	}
}
