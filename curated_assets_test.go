package xumm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
)

func TestCuratedAssets(t *testing.T) {

	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")

	ci := &models.CurrencyInfo{
		Id:       178,
		IssuerId: 185,
		Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		Currency: "USD",
		Name:     "US Dollar",
		Avatar:   "https://nd4d3do.dlvr.cloud/fiat-dollar.png",
	}
	ci2 := &models.CurrencyInfo{
		Id:       169,
		IssuerId: 182,
		Issuer:   "rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq",
		Currency: "EUR",
		Name:     "Euro",
		Avatar:   "https://nd4d3do.dlvr.cloud/fiat-euro.png",
	}
	ci3 := &models.CurrencyInfo{
		Id:       170,
		IssuerId: 182,
		Issuer:   "rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq",
		Currency: "USD",
		Name:     "US Dollar",
		Avatar:   "https://nd4d3do.dlvr.cloud/fiat-dollar.png",
	}

	i := &models.Issuer{
		Id:     185,
		Name:   "Bitstamp",
		Domain: "bitstamp.net",
		Avatar: "https://nd4d3do.dlvr.cloud/ex-bitstamp.png",
		Currencies: map[string]models.CurrencyInfo{
			"USD": *ci,
		},
	}
	i2 := &models.Issuer{
		Id:     182,
		Name:   "GateHub",
		Domain: "gatehub.net",
		Avatar: "https://nd4d3do.dlvr.cloud/ex-gatehub.png",
		Currencies: map[string]models.CurrencyInfo{
			"EUR": *ci2,
			"USD": *ci3,
		},
	}
	car := &models.CuratedAssetsResponse{
		Issuers:    []string{"Bitstamp", "GateHub"},
		Currencies: []string{"USD", "EUR", "BTC", "ETH"},
		Details: map[string]models.Issuer{
			"Bitstamp": *i,
			"GateHub":  *i2,
		},
	}

	json1 := testutils.ConvertJsonFileToString("static-test-data/curated_assets_test.json")
	json3 := `{
		"error": {
		  "reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
		  "code": 812
		}
	  }`

	var tests = []struct {
		testName       string
		inputValue     string
		expectedOutput *models.CuratedAssetsResponse
		expectedError  error
		httpStatusCode int
	}{

		{testName: "correctData", inputValue: json1, expectedOutput: car, expectedError: nil, httpStatusCode: 200},
		{testName: "errorResponse", inputValue: json3, expectedOutput: nil, expectedError: &ErrorResponse{ErrorResponseInternal{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, httpStatusCode: 403},
	}
	for _, tt := range tests {

		t.Run(tt.testName, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.inputValue, tt.httpStatusCode, m)
			c, _ := NewClient(WithHttpClient(m))

			ca, err := c.CuratedAssets()

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
