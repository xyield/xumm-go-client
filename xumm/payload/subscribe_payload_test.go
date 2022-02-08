package payload

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm"
)

var upgrader = websocket.Upgrader{}

func TestSubscribe(t *testing.T) {
	// done := make(chan string)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade:", err)
		}
		defer c.Close()

		d := anyjson.AnyJson{
			"message": "Welcome 10e94f5f-caa5-4030-8a58-6d9f3cbd9ac5",
		}

		c.WriteJSON(d)
		// c.WriteJSON(&models.XummPayload{
		// 	Meta: models.PayloadMeta{
		// 		Exists:              true,
		// 		UUID:                "f94fc5d2-0dfe-4123-9182-a9f3b5addc8a",
		// 		Multisign:           false,
		// 		Submit:              false,
		// 		Destination:         "",
		// 		ResolvedDestination: "",
		// 		Resolved:            false,
		// 		Signed:              false,
		// 		Cancelled:           false,
		// 		Expired:             false,
		// 		Pushed:              false,
		// 		AppOpened:           false,
		// 		OpenedByDeeplink:    nil,
		// 		ReturnURLApp:        "test",
		// 		ReturnURLWeb:        nil,
		// 		IsXapp:              false,
		// 	},
		// 	Application: models.PayloadApplication{
		// 		Name:            "test",
		// 		Description:     "test",
		// 		Disabled:        0,
		// 		Uuidv4:          "27AC8810-F458-4386-8ED9-2B9A4D9BE212",
		// 		IconURL:         "https://test.com",
		// 		IssuedUserToken: "test",
		// 	},
		// 	Payload: models.Payload{
		// 		TxType:           "SignIn",
		// 		TxDestination:    "",
		// 		TxDestinationTag: 0,
		// 		RequestJSON: anyjson.AnyJson{
		// 			"TransactionType": "SignIn",
		// 			"SignIn":          true,
		// 		},
		// 		Origintype:       "test",
		// 		Signmethod:       "test",
		// 		CreatedAt:        "2021-11-23T21:22:22Z",
		// 		ExpiresAt:        "2021-11-24T21:22:22Z",
		// 		ExpiresInSeconds: 86239,
		// 	},
		// 	Response: models.PayloadResponse{
		// 		Hex:                "test",
		// 		Txid:               "test",
		// 		ResolvedAt:         "test",
		// 		DispatchedTo:       "test",
		// 		DispatchedResult:   "test",
		// 		DispatchedNodetype: "test",
		// 		MultisignAccount:   "test",
		// 		Account:            "test",
		// 	},
		// })
		// <-done
	}))

	defer s.Close()

	cfg, _ := xumm.NewConfig()

	// defer close(ch)
	wsURL, _ := convertHttpToWS(s.URL)
	p := &Payload{
		Cfg: cfg,
		WSCfg: WSCfg{
			url: wsURL,
		},
	}

	_, err := p.Subscribe("10e94f5f-caa5-4030-8a58-6d9f3cbd9ac5")
	assert.NoError(t, err)

	var msgs []anyjson.AnyJson
	for v := range p.WSCfg.msgs {
		fmt.Println(v)
		msgs = append(msgs, v)
	}
	assert.Equal(t, []anyjson.AnyJson{{"message": "Welcome 10e94f5f-caa5-4030-8a58-6d9f3cbd9ac5"}}, msgs)
}

// type testServer struct{}

// func (t *testServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// }

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
