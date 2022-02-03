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

func MockWebSocket(wsFunc http.HandlerFunc)

func TestSubscribe(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade:", err)
		}
		defer c.Close()

		d := anyjson.AnyJson{
			"message": "Hello",
		}

		c.WriteJSON(d)
	}))

	// res, _ := http.Get(s.URL)
	// b, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(b))

	cfg, _ := xumm.NewConfig()

	wsURL, _ := convertHttpToWS(s.URL)
	p := &Payload{
		Cfg: cfg,
		WSCfg: WSCfg{
			url: wsURL,
		},
	}

	_, err := p.Subscribe("adjpdjdads")

	assert.NoError(t, err)
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
