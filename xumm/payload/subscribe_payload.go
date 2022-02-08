package payload

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/xyield/xumm-go-client/utils"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type PayloadExpiredError struct {
	UUID string
}

func (e *PayloadExpiredError) Error() string {
	return fmt.Sprintf("Payload with uuid %v expired", e.UUID)
}

func (p *Payload) Subscribe(uuid string) (*models.XummPayload, error) {
	ws, _, err := websocket.DefaultDialer.Dial(p.WSCfg.url, nil)
	if err != nil {
		log.Println("Error connecting to websocket:", err)
		return nil, err
	}
	// defer ws.Close()

	msgsc := make(chan anyjson.AnyJson)
	done := make(chan string)
	expired := make(chan bool)
	// p.WSCfg.done = make(chan string)
	// p.WSCfg.expired = make(chan bool)

	go recieveMessage(ws, msgsc, done, expired)

	// TO DO: Write a for loop that accepts msgs
	for {
		select {
		case v := <-msgsc:
			utils.PrettyPrintJson(v)
			p.WSCfg.msgs = append(p.WSCfg.msgs, v)
		case <-done:
			fmt.Println("Payload resolved")
			return p.GetPayloadByUUID(uuid)
		case <-expired:
			fmt.Println("Payload expired")
			return nil, &PayloadExpiredError{UUID: uuid}
		}
	}

	// return nil, nil
}

// Recieves messages from an open connection reads them and fires them into a channel
func recieveMessage(conn *websocket.Conn, msgs chan anyjson.AnyJson, done chan string, expired chan bool) {
	// defer close(msgs)
	defer conn.Close()

	for {
		var msg anyjson.AnyJson
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Could not read message!")
			log.Println(err)
		}
		// utils.PrettyPrintJson(msg)
		msgs <- msg
		if checkMessage(msg, "payload_uuidv4") {
			done <- msg["payload_uuidv4"].(string)
			return
		}
		if checkMessage(msg, "expired") {
			expired <- msg["expired"].(bool)
			return
		}
	}
}

// Check if message contains payload uuid or expired true
func checkMessage(m anyjson.AnyJson, k string) bool {
	if _, ok := m[k]; ok {
		return ok
	}

	return false
}
