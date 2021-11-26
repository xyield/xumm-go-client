package payload

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/pkg/json"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestCancelPayloadByUuid(t *testing.T) {
	tt := []struct {
		description    string
		uuid           string
		jsonResponse   string
		expectedOutput *models.XummDeletePayloadResponse
		statusCode     int
		expectedError  error
	}{
		{
			description: "successfully cancelled",
			uuid:        "XXX",
			jsonResponse: `{
				"result": {
				  "cancelled": true,
				  "reason": "OK"
				},
				"meta": {
				  "exists": true,
				  "uuid": "XXX",
				  "multisign": false,
				  "submit": true,
				  "destination": "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
				  "resolved_destination": "XRP Tip Bot",
				  "finished": false,
				  "expired": true,
				  "pushed": true,
				  "app_opened": false,
				  "return_url_app": "<some-url-or-null>",
				  "return_url_web": "<some-url-or-null>"
				},
				"custom_meta": {
				  "identifier": "some_identifier_1337",
				  "blob": {},
				  "instruction": "Hey ❤️ ..."
				}
			  }`,
			expectedOutput: &models.XummDeletePayloadResponse{
				Result: models.XummCancelResult{
					Cancelled: true,
					Reason:    "OK",
				},
				Meta: models.PayloadMeta{
					Exists:              true,
					UUID:                "XXX",
					Multisign:           false,
					Submit:              true,
					Destination:         "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
					ResolvedDestination: "XRP Tip Bot",
					Finished:            false,
					Expired:             true,
					Pushed:              true,
					AppOpened:           false,
					ReturnURLApp:        "<some-url-or-null>",
					ReturnURLWeb:        "<some-url-or-null>",
				},
				CustomMeta: models.XummCustomMeta{
					Identifier:  "some_identifier_1337",
					Blob:        anyjson.AnyJson{},
					Instruction: "Hey ❤️ ...",
				},
			},
			statusCode:    200,
			expectedError: nil,
		},
		{
			description: "OK but not cancelled",
			uuid:        "XXX",
			jsonResponse: `{
				"result": {
				  "cancelled": false,
				  "reason": "<some-reason-see-note-below>"
				},
				"meta": {
				  "exists": true,
				  "uuid": "XXX",
				  "multisign": false,
				  "submit": true,
				  "destination": "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
				  "resolved_destination": "XRP Tip Bot",
				  "finished": false,
				  "expired": true,
				  "pushed": true,
				  "app_opened": false,
				  "return_url_app": "<some-url-or-null>",
				  "return_url_web": "<some-url-or-null>",
				  "custom_identifier": "some_identifier_1337",
				  "custom_blob": {},
				  "custom_instruction": "Hey ❤️ ..."
				}
			  }`,
			expectedOutput: &models.XummDeletePayloadResponse{
				Result: models.XummCancelResult{
					Cancelled: false,
					Reason:    "<some-reason-see-note-below>",
				},
				Meta: models.PayloadMeta{
					Exists:              true,
					UUID:                "XXX",
					Multisign:           false,
					Submit:              true,
					Destination:         "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
					ResolvedDestination: "XRP Tip Bot",
					Finished:            false,
					Expired:             true,
					Pushed:              true,
					AppOpened:           false,
					ReturnURLApp:        "<some-url-or-null>",
					ReturnURLWeb:        "<some-url-or-null>",
					CustomIdentifier:    "some_identifier_1337",
					CustomBlob:          anyjson.AnyJson{},
					CustomInstruction:   "Hey ❤️ ...",
				},
			},
			statusCode:    200,
			expectedError: nil,
		},
		{
			description: "Cancel payload error",
			uuid:        "XXX",
			jsonResponse: `{
				"error": {
				  "reference": "d1ad8cf2-1e4a-4d7d-b1f5-5c692770bd28",
				  "code": 404
				}
			  }`,
			expectedOutput: nil,
			statusCode:     404,
			expectedError:  &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "d1ad8cf2-1e4a-4d7d-b1f5-5c692770bd28", Code: 404}},
		},
		{
			description: "UUID error",
			uuid:        "XXX",
			jsonResponse: `{
				"error": true,
				"message": "Endpoint unknown or method invalid for given endpoint",
				"reference": "",
				"code": 404,
				"req": "/v1/platform/payload/xxx",
				"method": "DELETE"
			  }`,
			expectedOutput: nil,
			statusCode:     404,
			expectedError: &xumm.ErrorNotFound{Err: true, Message: "Endpoint unknown or method invalid for given endpoint", Reference: "",
				Code: 404, Req: "/v1/platform/payload/xxx", Method: "DELETE"},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(test.jsonResponse, test.statusCode, m)
			c, _ := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			p := &Payload{
				Cfg: c,
			}
			res, err := p.CancelPayloadByUUID(test.uuid)

			if test.expectedError != nil {
				assert.Nil(t, res)
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectedError.Error())
			} else {
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
				assert.Equal(t, test.expectedOutput, res)
			}
		})
	}
}
