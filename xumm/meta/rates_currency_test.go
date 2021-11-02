package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
)

func TestRatesCurrency(t *testing.T) {

	c := &models.RatesCurrencyResponse{
		Usd: 1,
		Xrp: 1.04,
		Meta: models.Meta{
			Currency: models.Currency{
				En:     "US Dollar",
				Code:   "USD",
				Symbol: "$",
			},
		},
	}

	validJson := `{
		"USD": 1,
		"XRP": 1.04,
		"__meta": {
		  "currency": {
			"en": "US Dollar",
			"code": "USD",
			"symbol": "$"
		  }
		}
	  }`

	errorJson := `{
		"error": {
		  "reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
		  "code": 812
		}
	  }`

	var tests = []struct {
		testName       string
		testValue      string
		inputValue     string
		expectedOutput *models.RatesCurrencyResponse
		expectedError  error
		httpStatusCode int
	}{

		{testName: "correct data", testValue: "USD", inputValue: validJson, expectedOutput: c, expectedError: nil, httpStatusCode: 200},
		{testName: "Incorrect length currency code", testValue: "USDaas", inputValue: errorJson, expectedOutput: nil, expectedError: &CurrencyCodeError{Code: "USDaas"}, httpStatusCode: -1},
		{testName: "Incorrect characters in currency code", testValue: "US$", inputValue: errorJson, expectedOutput: nil, expectedError: &CurrencyCodeError{Code: "US$"}, httpStatusCode: -1},
		{testName: "error response", testValue: "USD", inputValue: errorJson, expectedOutput: nil, expectedError: &xumm.ErrorResponse{ErrorResponseInternal: xumm.ErrorResponseInternal{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, httpStatusCode: 403},
	}

	for _, tt := range tests {

		t.Run(tt.testName, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.inputValue, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}
			ca, err := meta.RatesCurrency(tt.testValue)

			if tt.expectedError != nil {
				assert.Nil(t, ca)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, ca)
				assert.Equal(t, http.Header{
					"XUMM_API_KEY":    {"testApiKey"},
					"XUMM_API_SECRET": {"testApiSecret"},
					"Content-Type":    {"application/json"},
				}, m.Spy.Header)
			}
		})
	}
}
