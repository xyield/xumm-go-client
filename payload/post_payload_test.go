package payload

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/models"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
)

func TestPostPayload(t *testing.T) {

	tests := []struct {
		description    string
		payloadRequest *models.XummPostPayload
		jsonRequest    string
		jsonResponse   string
		expectedOutput *models.CreatedPayload
		expectedError  error
		httpStatusCode int
	}{
		{
			description: "successful POST payload request",
			payloadRequest: &models.XummPostPayload{
				UserToken: "token",
				TxJson: anyjson.AnyJson{
					"TransactionType": "Payment",
				},
			},
			jsonRequest: testutils.ConvertJsonFileToJsonString("static-test-data/post_payload_request.json"),
			jsonResponse: `{
				"uuid": "95516771-5c9e-4b90-ab04-116c938ddba4",
				"next": {
				  "always": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4",
				  "no_push_msg_received": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4/qr"
				},
				"refs": {
				  "qr_png": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.png",
				  "qr_matrix": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.json",
				  "qr_uri_quality_opts": [
					"m",
					"q",
					"h"
				  ],
				  "websocket_status": "wss://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4"
				},
				"pushed": true
			  }`,
			expectedOutput: &models.CreatedPayload{
				UUID: "95516771-5c9e-4b90-ab04-116c938ddba4",
				Next: models.Next{
					Always:            "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4",
					NoPushMsgReceived: "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4/qr",
				},
				Refs: models.Refs{
					WebsocketStatus: "wss://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4",
					QrPng:           "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.png",
					QrMatrix:        "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.json",
					QrURIQualityOpts: []string{
						"m",
						"q",
						"h",
					},
				},
				Pushed: true,
			},
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description: "unsucessful POST paylaod request - no transactionType",
			payloadRequest: &models.XummPostPayload{
				UserToken: "token",
				TxJson: anyjson.AnyJson{
					"noTransactionType": "test",
				},
			},
			jsonRequest: testutils.ConvertJsonFileToJsonString("static-test-data/post_payload_invalid_request.json"),
			jsonResponse: `{
				"uuid": "95516771-5c9e-4b90-ab04-116c938ddba4",
				"next": {
				  "always": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4",
				  "no_push_msg_received": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4/qr"
				},
				"refs": {
				  "qr_png": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.png",
				  "qr_matrix": "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.json",
				  "qr_uri_quality_opts": [
					"m",
					"q",
					"h"
				  ],
				  "websocket_status": "wss://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4"
				},
				"pushed": true
			  }`,
			expectedOutput: &models.CreatedPayload{
				UUID: "95516771-5c9e-4b90-ab04-116c938ddba4",
				Next: models.Next{
					Always:            "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4",
					NoPushMsgReceived: "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4/qr",
				},
				Refs: models.Refs{
					WebsocketStatus: "wss://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4",
					QrPng:           "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.png",
					QrMatrix:        "https://xumm.app/sign/95516771-5c9e-4b90-ab04-116c938ddba4_q.json",
					QrURIQualityOpts: []string{
						"m",
						"q",
						"h",
					},
				},
				Pushed: true,
			},
			expectedError:  &TransactionTypeError{},
			httpStatusCode: 0,
		},
		{
			description: "unsucessful POST paylaod request - bad request/duplicate",
			payloadRequest: &models.XummPostPayload{
				UserToken: "token",
				TxJson: anyjson.AnyJson{
					"TransactionType": "Payment",
				},
			},
			jsonRequest: testutils.ConvertJsonFileToJsonString("static-test-data/post_payload_request.json"),
			jsonResponse: `{
				"error": {
				  "reference": "95516771-5c9e-4b90-ab04-116c938ddba4",
				  "code": 600
				}
			  }`,
			expectedOutput: nil,
			expectedError:  &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "95516771-5c9e-4b90-ab04-116c938ddba4", Code: 600}},
			httpStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.jsonResponse, tt.httpStatusCode, m)
			c, _ := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			p := &Payload{
				Cfg: c,
			}

			pr, err := p.PostPayload(*tt.payloadRequest)

			if tt.expectedError != nil {
				assert.Nil(t, pr)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				body, _ := ioutil.ReadAll(m.Spy.Body)
				assert.JSONEq(t, tt.jsonRequest, string(body))
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, pr)
			}
		})
	}
}
