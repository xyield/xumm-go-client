//go:build unit
// +build unit

package payload

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetPayloadUuid(t *testing.T) {

	tt := []struct {
		description      string
		uuid             string
		jsonResponse     string
		expectedOutput   *models.XummPayload
		expectedError    error
		httpResponseCode int
	}{
		{
			description:  "valid uuid",
			uuid:         "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
			jsonResponse: testutils.ConvertJsonFileToJsonString("static-test-data/valid_get_payload_response.json"),
			expectedOutput: &models.XummPayload{
				Meta: models.PayloadMeta{
					Exists:              true,
					UUID:                "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
					Multisign:           false,
					Submit:              false,
					Destination:         "",
					ResolvedDestination: "",
					Resolved:            false,
					Signed:              false,
					Cancelled:           false,
					Expired:             false,
					Pushed:              false,
					AppOpened:           false,
					OpenedByDeeplink:    nil,
					ReturnURLApp:        "test",
					ReturnURLWeb:        nil,
					IsXapp:              false,
				},
				Application: models.PayloadApplication{
					Name:            "test",
					Description:     "test",
					Disabled:        0,
					Uuidv4:          "27AC8810-F458-4386-8ED9-2B9A4D9BE212",
					IconURL:         "https://test.com",
					IssuedUserToken: "test",
				},
				Payload: models.Payload{
					TxType:           "SignIn",
					TxDestination:    "",
					TxDestinationTag: 0,
					RequestJSON: anyjson.AnyJson{
						"TransactionType": "SignIn",
						"SignIn":          true,
					},
					Origintype:       "test",
					Signmethod:       "test",
					CreatedAt:        "2021-11-23T21:22:22Z",
					ExpiresAt:        "2021-11-24T21:22:22Z",
					ExpiresInSeconds: 86239,
				},
				Response: models.PayloadResponse{
					Hex:                "test",
					Txid:               "test",
					ResolvedAt:         "test",
					DispatchedTo:       "test",
					DispatchedResult:   "test",
					DispatchedNodetype: "test",
					MultisignAccount:   "test",
					Account:            "test",
				},
			},
			expectedError:    nil,
			httpResponseCode: 200,
		},
		{
			description:      "Empty uuid provided",
			uuid:             "",
			jsonResponse:     "",
			expectedOutput:   nil,
			expectedError:    &EmptyUuidError{},
			httpResponseCode: 0,
		},
		{
			description: "uuid not found",
			uuid:        "XXX",
			jsonResponse: `{
				"error": {
					"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
					"code": 404
				}
			}`,
			expectedOutput:   nil,
			expectedError:    &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 404}},
			httpResponseCode: 404,
		},
		{
			description: "uuid error",
			uuid:        "XXX",
			jsonResponse: `{
				"error": true,
				"message": "Endpoint unknown or method invalid for given endpoint",
				"reference": "",
				"code": 404,
				"req": "/v1/platform/payload/<some-uuid>",
				"method": "GET"
			  }`,
			expectedOutput: nil,
			expectedError: &xumm.ErrorNotFound{Err: true, Message: "Endpoint unknown or method invalid for given endpoint", Reference: "",
				Code: 404, Req: "/v1/platform/payload/<some-uuid>", Method: "GET"},
			httpResponseCode: 404,
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(test.jsonResponse, test.httpResponseCode, m)

			c, _ := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			p := &Payload{
				Cfg: c,
			}

			pr, err := p.GetPayloadByUUID(test.uuid)

			if test.expectedError != nil {
				assert.Nil(t, pr)
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectedError.Error())
			} else {
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
				assert.Equal(t, test.expectedOutput, pr)
			}
		})
	}
}
