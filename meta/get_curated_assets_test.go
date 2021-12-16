package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/models"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
)

func TestGetCuratedAssets(t *testing.T) {

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

	validJson := testutils.ConvertJsonFileToJsonString("static-test-data/curated_assets_test.json")
	errorJson := `{
		"error": {
		  "reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
		  "code": 812
		}
	  }`

	tests := []struct {
		testName       string
		inputValue     string
		expectedOutput *models.CuratedAssetsResponse
		expectedError  error
		httpStatusCode int
	}{

		{testName: "correctData", inputValue: validJson, expectedOutput: car, expectedError: nil, httpStatusCode: 200},
		{testName: "errorResponse", inputValue: errorJson, expectedOutput: nil, expectedError: &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, httpStatusCode: 403},
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
			ca, err := meta.GetCuratedAssets()

			if tt.expectedError != nil {
				assert.Nil(t, ca)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, ca)
			}
		})
	}
}
