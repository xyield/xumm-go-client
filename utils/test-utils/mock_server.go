package testutils

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gorilla/websocket"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
)

type MockWebSocketServer struct {
	Msgs []anyjson.AnyJson
}

type connFn func(*websocket.Conn)

func (ms *MockWebSocketServer) TestWebSocketServer(writeFunc connFn) *httptest.Server {
	var upgrader = websocket.Upgrader{}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade:", err)
		}

		writeFunc(c)
	}))

	return s
}

func ConvertHttpToWS(u string) (string, error) {
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
