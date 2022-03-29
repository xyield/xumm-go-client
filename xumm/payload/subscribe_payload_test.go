//go:build unit
// +build unit

package payload

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"syscall"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestSubscribe(t *testing.T) {

	tt := []struct {
		description      string
		messages         []anyjson.AnyJson
		uuid             string
		jsonResponse     string
		httpResponseCode int
		expectedOutput   *models.XummPayload
		expectedError    error
		interrupt        bool
	}{
		{
			description: "Successful subscribe and payload grab",
			messages: []anyjson.AnyJson{
				{"message": "Welcome f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
				{"payload_uuidv4": "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a6"},
			},
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
			httpResponseCode: 200,
			expectedError:    nil,
			interrupt:        false,
		},
		{
			description:      "Payload UUID does not exist",
			messages:         []anyjson.AnyJson{{"message": "..."}},
			uuid:             "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
			jsonResponse:     "",
			expectedOutput:   nil,
			httpResponseCode: 200,
			expectedError:    &PayloadUuidError{UUID: "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
			interrupt:        false,
		},
		{
			description: "Payload expired",
			messages: []anyjson.AnyJson{
				{"message": "Welcome f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
				{"expires_in_seconds": 10},
				{"expires_in_seconds": 5},
				{"expires_in_seconds": 1},
				{"expired": true},
			},
			uuid:             "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
			jsonResponse:     "",
			expectedOutput:   nil,
			httpResponseCode: 200,
			expectedError:    &PayloadExpiredError{UUID: "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
			interrupt:        false,
		},
		{
			description: "Connection interrupted",
			messages: []anyjson.AnyJson{
				{"message": "Welcome f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
				{"expires_in_seconds": 10},
			},
			uuid:             "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
			jsonResponse:     "",
			expectedOutput:   nil,
			httpResponseCode: 200,
			expectedError:    &ConnectionError{UUID: "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
			interrupt:        true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			ms := &testutils.MockWebSocketServer{
				Msgs: tc.messages,
			}

			s := ms.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tc.messages {
					err := c.WriteJSON(m)
					if err != nil {
						println("error writing message")
					}
				}
				if tc.interrupt == true {
					time.Sleep(time.Microsecond * 20)
					err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
					if err != nil {
						println("interrupt failed")
					}
				}
			})

			defer s.Close()

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tc.jsonResponse, tc.httpResponseCode, m)
			cfg, _ := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))

			wsURL, _ := testutils.ConvertHttpToWS(s.URL)
			p := &Payload{
				Cfg: cfg,
				WSCfg: WSCfg{
					BaseURL: wsURL + "/",
				},
			}

			actual, err := p.Subscribe(tc.uuid)

			if tc.expectedError != nil {
				assert.Nil(t, actual)
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.messages, p.WSCfg.msgs)
				assert.Equal(t, tc.expectedOutput, actual)
			}
		})
	}
}

func TestCreateAndSubscribe(t *testing.T) {

	type httpResponse struct {
		response http.Response
		error    error
	}

	tt := []struct {
		description             string
		payloadRequest          *models.XummPostPayload
		responses               []httpResponse
		jsonPostPayloadRequest  string
		messages                []anyjson.AnyJson
		uuid                    string
		expectedSubscribeOutput *models.XummPayload
		interrupt               bool
	}{
		{
			description: "successful create and subscribe",
			payloadRequest: &models.XummPostPayload{
				UserToken: "token",
				TxJson: anyjson.AnyJson{
					"TransactionType": "SignIn",
				},
			},
			responses: []httpResponse{
				{
					response: http.Response{
						StatusCode: 200,
						Body: ioutil.NopCloser(bytes.NewBufferString(`{
							"uuid": "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
							"next": {
							  "always": "https://xumm.app/sign/f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
							  "no_push_msg_received": "https://xumm.app/sign/f94fc5d2-0dfe-4123-9182-a9f3b5addc8a/qr"
							},
							"refs": {
							  "qr_png": "https://xumm.app/sign/f94fc5d2-0dfe-4123-9182-a9f3b5addc8a_q.png",
							  "qr_matrix": "https://xumm.app/sign/f94fc5d2-0dfe-4123-9182-a9f3b5addc8a_q.json",
							  "qr_uri_quality_opts": [
								"m",
								"q",
								"h"
							  ],
							  "websocket_status": "wss://xumm.app/sign/f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"
							},
							"pushed": true
						  }`)),
					},
					error: nil,
				},
				{
					response: http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBufferString(testutils.ConvertJsonFileToJsonString("static-test-data/valid_get_payload_response.json"))),
					},
					error: nil,
				},
			},

			jsonPostPayloadRequest: `{
				"user_token": "token",
				"txjson": {
					"TransactionType": "SignIn"
				}
			}`,
			messages: []anyjson.AnyJson{
				{"message": "Welcome f94fc5d2-0dfe-4123-9182-a9f3b5addc8a"},
				{"payload_uuidv4": "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a6"},
			},
			uuid: "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
			expectedSubscribeOutput: &models.XummPayload{
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
			interrupt: false,
		},
	}

	for _, tcs := range tt {
		t.Run(tcs.description, func(t *testing.T) {

			ms := &testutils.MockWebSocketServer{
				Msgs: tcs.messages,
			}

			s := ms.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tcs.messages {
					err := c.WriteJSON(m)
					if err != nil {
						println("error writing message")
					}
				}
				if tcs.interrupt == true {
					err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
					if err != nil {
						println("interrupt failed")
					}
				}
			})

			defer s.Close()

			m := &testutils.MockClient{}

			responseCounter := 0
			m.DoFunc = func(req *http.Request) (*http.Response, error) {
				responseCounter++
				return &tcs.responses[responseCounter-1].response, tcs.responses[responseCounter-1].error
			}

			cfg, _ := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			wsURL, _ := testutils.ConvertHttpToWS(s.URL)
			p := &Payload{
				Cfg: cfg,
				WSCfg: WSCfg{
					baseUrl: wsURL + "/",
				},
			}

			actual, err := p.CreateAndSubscribe(*tcs.payloadRequest)

			if tcs.responses[1].error != nil {
				assert.Nil(t, actual)
				assert.Error(t, err)
				assert.EqualError(t, err, tcs.responses[1].error.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tcs.messages, p.WSCfg.msgs)
				assert.Equal(t, tcs.expectedSubscribeOutput, actual)
			}
		})
	}
}

func TestCheckMessage(t *testing.T) {
	tt := []struct {
		description string
		input       anyjson.AnyJson
		key         string
		expected    bool
	}{
		{
			description: "Message contains payload uuid field",
			input: anyjson.AnyJson{
				"payload_uuidv4": "ccb0ca8e-d498-4aa8-bed0-d55d9015f556",
			},
			key:      "payload_uuidv4",
			expected: true,
		},
		{
			description: "Message contains expired field",
			input: anyjson.AnyJson{
				"expired": "true",
			},
			key:      "expired",
			expected: true,
		},
		{
			description: "Message doesn't contain a required field",
			input: anyjson.AnyJson{
				"message": "Welcome ccb0ca8e-d498-4aa8-bed0-d55d9015f556",
			},
			key:      "expired",
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, checkMessage(tc.input, tc.key))
		})
	}
}
