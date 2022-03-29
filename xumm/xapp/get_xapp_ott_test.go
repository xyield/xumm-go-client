//go:build unit
// +build unit

package xapp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetXappOtt(t *testing.T) {

	outputXappResponse := &models.XappOttResponse{
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
		description    string
		ottInput       string
		jsonResponse   string
		expectedOutput *models.XappOttResponse
		expectedError  error
		httpStatusCode int
	}{
		{
			description: "valid get ott",
			ottInput:    "token",
			jsonResponse: `{
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
			}`,
			expectedOutput: outputXappResponse,
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description:    "check ottInput isn't empty",
			ottInput:       "",
			jsonResponse:   "",
			expectedOutput: nil,
			expectedError:  &InvalidToken{},
			httpStatusCode: 0,
		},
		{
			description: "error response",
			ottInput:    "token",
			jsonResponse: `{
				"error": {
				  "reference": "000000-81ba-4b3c-baa4-b2ff3c1b445e",
				  "code": 400
				}
			  }`,
			expectedOutput: nil,
			expectedError:  &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "000000-81ba-4b3c-baa4-b2ff3c1b445e", Code: 400}},
			httpStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.jsonResponse, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			xapp := &Xapp{
				Cfg: cfg,
			}

			xr, err := xapp.GetXappOtt(tt.ottInput)

			if tt.expectedError != nil {
				assert.Nil(t, xr)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, xr)
			}
		})
	}
}
