package payload

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

var upgrader = websocket.Upgrader{}

func TestSubscribe(t *testing.T) {

	tt := []struct {
		description      string
		messages         []anyjson.AnyJson
		uuid             string
		jsonResponse     string
		httpResponseCode int
		expectedOutput   *models.XummPayload
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				upgrader.CheckOrigin = func(r *http.Request) bool { return true }
				c, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					log.Println("Upgrade:", err)
				}

				for _, m := range tc.messages {
					c.WriteJSON(m)
				}
			}))

			defer s.Close()

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tc.jsonResponse, tc.httpResponseCode, m)
			cfg, _ := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))

			wsURL, _ := convertHttpToWS(s.URL)
			p := &Payload{
				Cfg: cfg,
				WSCfg: WSCfg{
					url: wsURL,
				},
			}

			actual, err := p.Subscribe(tc.uuid)
			assert.NoError(t, err)

			assert.Equal(t, tc.messages, p.WSCfg.msgs)
			assert.Equal(t, tc.expectedOutput, actual)
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

func convertHttpToWS(u string) (string, error) {
	s, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	switch s.Scheme {
	case "http":
		s.Scheme = "ws"
	case "https":
		s.Scheme = "wss"
	}

	return s.String(), nil
}
