package xumm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
)

func TestKycAccountStatusTest(t *testing.T) {
	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")
	tt := []struct {
		description    string
		input          string
		json           string
		expectedOutput *models.KycAccountStatusResponse
	}{
		{
			description: "Valid account with kyc status true",
			input:       "rGBP1ZYpgiArYbDSvqu7Ps8AmWrD6hiqwe",
			json: `{
				"account": "rGBP1ZYpgiArYbDSvqu7Ps8AmWrD6hiqwe",
				"kycApproved": true
			  }`,
			expectedOutput: &models.KycAccountStatusResponse{
				Account:     "rGBP1ZYpgiArYbDSvqu7Ps8AmWrD6hiqwe",
				KycApproved: true,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(test.json, 200, m)

			c, _ := NewClient(WithHttpClient(m))

			customer, _ := c.KycAccountStatus(test.input)
			assert.Equal(t, test.expectedOutput, customer)
		})
	}
}

func TestKycStatusState(t *testing.T) {
	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")
	tt := []struct {
		description    string
		input          models.KycStatusStateRequest
		json           string
		expectedOutput *models.KycStatusStateResponse
	}{
		{
			description: "Valid account with kyc status none",
			input: models.KycStatusStateRequest{
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
			expectedOutput: &models.KycStatusStateResponse{
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
			c, _ := NewClient(WithHttpClient(m))

			customer, _ := c.KycStatusState(test.input)
			assert.Equal(t, test.expectedOutput, customer)
		})
	}
}
