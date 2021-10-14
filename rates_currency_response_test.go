package xumm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
)

func TestRatesCurrency(t *testing.T) {

	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")

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

	json1 := `{
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

	json2 := `{
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

		{testName: "correct data", testValue: "USD", inputValue: json1, expectedOutput: c, expectedError: nil, httpStatusCode: 200},
		{testName: "Incorrect length currency code", testValue: "USDaas", inputValue: json2, expectedOutput: nil, expectedError: &CurrencyCodeError{Code: "USDaas"}, httpStatusCode: -1},
		{testName: "Incorrect characters in currency code", testValue: "US$", inputValue: json2, expectedOutput: nil, expectedError: &CurrencyCodeError{Code: "US$"}, httpStatusCode: -1},
		{testName: "error response", testValue: "USD", inputValue: json2, expectedOutput: nil, expectedError: &ErrorResponse{ErrorResponseInternal{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, httpStatusCode: 403},
	}

	for _, tt := range tests {

		t.Run(tt.testName, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.inputValue, tt.httpStatusCode, m)
			c, _ := NewClient(WithHttpClient(m))

			ca, err := c.RatesCurrency(tt.testValue)

			if tt.expectedError != nil {
				assert.Nil(t, ca)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, ca)
			}
		})
	}
}
