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

func TestPing(t *testing.T) {

	var tests = []struct {
		description    string
		jsonResponse   string
		expectedOutput *models.Pong
		expectedError  error
		httpStatusCode int
	}{
		{
			description: "successful ping request",
			jsonResponse: `{
				"pong": true,
				"auth": {
				  "quota": {},
				  "application": {
					"uuidv4": "8525e32b-1bd0-4839-af2f-f794874a80b0",
					"name": "test-application",
					"webhookurl": "https://test-webhook",
					"disabled": 0
				  },
				  "call": {
					"uuidv4": "4b97cf7a-1837-471f-baed-2ebebcf5adb4"
				  }
				}
			  }`,
			expectedOutput: &models.Pong{
				Pong: true,
				Auth: models.ApplicationDetails{
					Quota: map[string]interface{}{},
					Application: models.Application{
						UUIDV4:     "8525e32b-1bd0-4839-af2f-f794874a80b0",
						Name:       "test-application",
						WebhookUrl: "https://test-webhook",
						Disabled:   0,
					},
					Call: models.Call{
						UUIDV4: "4b97cf7a-1837-471f-baed-2ebebcf5adb4",
					},
				},
			},
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description: "forbidden ping request",
			jsonResponse: `{
				"error": {
				 	"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
				 	"code": 403
				}
			}`,
			expectedOutput: nil,
			expectedError: &xumm.ErrorResponse{
				ErrorResponseBody: xumm.ErrorResponseBody{
					Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
					Code:      403,
				},
			},
			httpStatusCode: 403,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.jsonResponse, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}
			pong, err := meta.Ping()
			assert.Equal(t, http.Header{
				"X-API-Key":    {"testApiKey"},
				"X-API-Secret": {"testApiSecret"},
				"Content-Type": {"application/json"},
			}, m.Spy.Header)

			if tt.expectedError != nil {
				assert.Nil(t, pong)
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, pong)
			}
		})
	}
}
