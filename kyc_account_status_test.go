package xumm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
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
			m := &MockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					b := ioutil.NopCloser(bytes.NewBuffer([]byte(test.json)))
					return &http.Response{
						StatusCode: 200,
						Body:       b,
					}, nil
				},
			}
			c, _ := NewClient(WithHttpClient(m))

			customer, _ := c.KycAccountStatus(test.input)
			assert.Equal(t, test.expectedOutput, customer)
		})
	}
}
