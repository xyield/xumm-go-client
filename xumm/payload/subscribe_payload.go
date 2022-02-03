package payload

import (
	"fmt"
	"log"

	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm/models"
	"golang.org/x/net/websocket"
)

type WSConn interface {
	Read(msg []byte) (n int, err error)
	Write(msg []byte) (n int, err error)
}

func (p *Payload) Subscribe(uuid string) (*models.XummPayload, error) {
	ws, err := websocket.Dial(p.WSCfg.url, "", "http://localhost")
	if err != nil {
		log.Println(err)
	}
	var j anyjson.AnyJson
	websocket.JSON.Receive(ws, &j)
	fmt.Printf("Mesage: %+v", j)
	// fmt.Printf("%+v", ws)
	return nil, nil
}
