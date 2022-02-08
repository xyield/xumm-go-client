package payload

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func (p *Payload) Subscribe(uuid string) (*models.XummPayload, error) {
	ws, _, err := websocket.DefaultDialer.Dial(p.WSCfg.url, nil)
	if err != nil {
		log.Println("Error connecting to websocket:", err)
		return nil, err
	}
	defer ws.Close()

	p.WSCfg.msgs = make(chan anyjson.AnyJson)

	go recieveMessage(ws, p.WSCfg.msgs)

	// TO DO: Write a for loop that accepts msgs

	//hack to make websocket doesn't close before message is sent
	time.Sleep(time.Duration(time.Millisecond) * 50)

	return nil, nil
}

func recieveMessage(conn *websocket.Conn, c chan anyjson.AnyJson) {
	defer close(c)

	for {
		var msg anyjson.AnyJson
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Could not read message!")
			log.Println(err)
		}
		fmt.Println("message:", msg)
		c <- msg
		//hack until something can break out of the loop
		break
	}
}
