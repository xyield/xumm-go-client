package xapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
)

func TestGetXappOtt(t *testing.T) {

	testJson := `{
		"locale": "en",
		"version": "1.0.1",
		"account": "r...",
		"accountaccess": "FULL",
		"accounttype": "REGULAR",
		"style": "LIGHT",
		"origin": {
		"type": "TX",
		"data": {
			"txid": "..."
		}
		},
		"user": "XUMM-App-UserUUID",
		"user_device": {
		"currency": "EUR"
		}
	}`

	outputXappResponse := &models.XappResponse{
		Locale:        "en",
		Version:       "1.0.1",
		Account:       "r...",
		Accountaccess: "FULL",
		Accounttype:   "REGULAR",
		Style:         "LIGHT",
		Origin: models.Origin{
			Type: "TX",
			Data: models.Data{
				Txid: "...",
			},
		},
		User: "XUMM-App-UserUUID",
		UserDevice: models.UserDevice{
			Currency: "EUR",
		},
	}

	var tests = []struct {
		testName       string
		ottInput       string
		jsonResponse   string
		expectedOutput *models.XappResponse
		expectedError  error
		httpStatusCode int
	}{
		{testName: "valid get ott", ottInput: "token", jsonResponse: testJson, expectedOutput: outputXappResponse, expectedError: nil, httpStatusCode: 200},
		// {testName: "check ottInput isn't empty", ottInput: "", jsonResponse: testJson, expectedOutput: nil, expectedError: invalidToken, httpStatusCode: 0},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.jsonResponse, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			xapp := &Xapp{
				Cfg: cfg,
			}

			xr, _ := xapp.GetXappOtt(tt.ottInput)

			assert.Equal(t, tt.expectedOutput, xr)
		})
	}
}
