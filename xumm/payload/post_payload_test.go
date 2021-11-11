package payload

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/pkg/json"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestPostPayload(t *testing.T) {

	tests := []struct {
		description    string
		payloadRequest *models.XummPostPayload
		jsonRequest    string
		jsonResponse   string
		expectedOutput *models.XummPostPayloadResponse
		expectedError  error
		httpStatusCode int
	}{
		{
			description: "sucessful POST paylaod request",
			payloadRequest: &models.XummPostPayload{
				UserToken: "token",
				TxJson: json.AnyJson{
					"TransactionType": "payment",
				},
			},
			jsonRequest: testutils.ConvertJsonFileToJsonString("static-test-data/post_payload_request.json"),
			jsonResponse: `{
				"uuid": "<payload-uuid>",
				"next": {
				  "always": "https://xumm.app/sign/<payload-uuid>",
				  "no_push_msg_received": "https://xumm.app/sign/<payload-uuid>/qr"
				},
				"refs": {
				  "qr_png": "https://xumm.app/sign/<payload-uuid>_q.png",
				  "qr_matrix": "https://xumm.app/sign/<payload-uuid>_q.json",
				  "qr_uri_quality_opts": [
					"m",
					"q",
					"h"
				  ],
				  "websocket_status": "wss://xumm.app/sign/<payload-uuid>"
				},
				"pushed": true
			  }`,
			expectedOutput: &models.XummPostPayloadResponse{
				UUID: "<payload-uuid>",
				Next: models.Next{
					Always:            "https://xumm.app/sign/<payload-uuid>",
					NoPushMsgReceived: "https://xumm.app/sign/<payload-uuid>/qr",
				},
				Refs: models.Refs{
					WebsocketStatus: "wss://xumm.app/sign/<payload-uuid>",
					QrPng:           "https://xumm.app/sign/<payload-uuid>_q.png",
					QrMatrix:        "https://xumm.app/sign/<payload-uuid>_q.json",
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
				TxJson: json.AnyJson{
					"noTransactionType": "test",
				},
			},
			jsonRequest: testutils.ConvertJsonFileToJsonString("static-test-data/post_payload_invalid_request.json"),
			jsonResponse: `{
				"uuid": "<payload-uuid>",
				"next": {
				  "always": "https://xumm.app/sign/<payload-uuid>",
				  "no_push_msg_received": "https://xumm.app/sign/<payload-uuid>/qr"
				},
				"refs": {
				  "qr_png": "https://xumm.app/sign/<payload-uuid>_q.png",
				  "qr_matrix": "https://xumm.app/sign/<payload-uuid>_q.json",
				  "qr_uri_quality_opts": [
					"m",
					"q",
					"h"
				  ],
				  "websocket_status": "wss://xumm.app/sign/<payload-uuid>"
				},
				"pushed": true
			  }`,
			expectedOutput: &models.XummPostPayloadResponse{
				UUID: "<payload-uuid>",
				Next: models.Next{
					Always:            "https://xumm.app/sign/<payload-uuid>",
					NoPushMsgReceived: "https://xumm.app/sign/<payload-uuid>/qr",
				},
				Refs: models.Refs{
					WebsocketStatus: "wss://xumm.app/sign/<payload-uuid>",
					QrPng:           "https://xumm.app/sign/<payload-uuid>_q.png",
					QrMatrix:        "https://xumm.app/sign/<payload-uuid>_q.json",
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
				assert.Equal(t, tt.jsonRequest, string(body))
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
