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
		c.WriteJSON(anyjson.AnyJson{
			"payload_uuidv4": "ccb0ca8e-d498-4aa8-bed0-d55d9015f556",
		})
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
		msgs = append(msgs, v)
	}
	assert.Equal(t, []anyjson.AnyJson{{"message": "Welcome 10e94f5f-caa5-4030-8a58-6d9f3cbd9ac5"}, {"payload_uuidv4": "ccb0ca8e-d498-4aa8-bed0-d55d9015f556"}}, msgs)
}

func TestCheckMessage(t *testing.T) {
	tt := []struct {
		description string
		input       anyjson.AnyJson
		expected    bool
	}{
		{
			description: "Message contains payload uuid field",
			input: anyjson.AnyJson{
				"payload_uuidv4": "ccb0ca8e-d498-4aa8-bed0-d55d9015f556",
			},
			expected: true,
		},
		{
			description: "Message contains expired field",
			input: anyjson.AnyJson{
				"expired": "true",
			},
			expected: true,
		},
		{
			description: "Message doesn't contain a required field",
			input: anyjson.AnyJson{
				"message": "Welcome ccb0ca8e-d498-4aa8-bed0-d55d9015f556",
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, checkMessage(tc.input))
		})
	}
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
