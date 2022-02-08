package payload

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/xyield/xumm-go-client/utils"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func (p *Payload) Subscribe(uuid string) (*models.XummPayload, error) {
	ws, _, err := websocket.DefaultDialer.Dial(p.WSCfg.url, nil)
	if err != nil {
		log.Println("Error connecting to websocket:", err)
		return nil, err
	}
	// defer ws.Close()

	p.WSCfg.msgs = make(chan anyjson.AnyJson)

	go recieveMessage(ws, p.WSCfg.msgs)

	// TO DO: Write a for loop that accepts msgs

	return nil, nil
}

// Recieves messages from an open connection reads them and fires them into a channel
func recieveMessage(conn *websocket.Conn, c chan anyjson.AnyJson) {
	defer close(c)
	defer conn.Close()

	for {
		var msg anyjson.AnyJson
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Could not read message!")
			log.Println(err)
		}
		c <- msg
		if checkMessage(msg) {
			return
		}
		// break
	}
}

// Check if message contains payload uuid or expired true
func checkMessage(m anyjson.AnyJson) bool {
	utils.PrettyPrintJson(m)
	if _, ok := m["payload_uuidv4"]; ok {
		return ok
	}

	if _, ok := m["expired"]; ok {
		return true
	}

	return false
}
