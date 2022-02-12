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

type PayloadUuidError struct {
	UUID string
}

func (e *PayloadUuidError) Error() string {
	return fmt.Sprintf("Payload with uuid %v does not exist", e.UUID)
}

// Subscribes to paylaod websocket to recieve messages and returns payload if it is resolved
func (p *Payload) Subscribe(uuid string) (*models.XummPayload, error) {
	ws, _, err := websocket.DefaultDialer.Dial(p.WSCfg.url, nil)
	if err != nil {
		log.Println("Error connecting to websocket:", err)
		return nil, err
	}

	msgsc := make(chan anyjson.AnyJson)
	done := make(chan string)

	go recieveMessage(ws, msgsc, done)

	for {
		select {
		case v := <-msgsc:
			utils.PrettyPrintJson(v)
			p.WSCfg.msgs = append(p.WSCfg.msgs, v)
		case m := <-done:
			if m == "resolved" {
				fmt.Println("Payload resolved")
				return p.GetPayloadByUUID(uuid)
			}
			if m == "expired" {
				fmt.Println("Payload expired")
				return nil, &PayloadExpiredError{UUID: uuid}
			}
			if m == "payloadUuidError" {
				fmt.Println("Payload does not exist")
				return nil, &PayloadUuidError{UUID: uuid}
			}
		}
	}
}

// Recieves messages from an open connection reads them and fires them into a channel
func recieveMessage(conn *websocket.Conn, msgs chan anyjson.AnyJson, done chan string) {
	defer conn.Close()

	for {
		var msg anyjson.AnyJson
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Could not read message!")
			log.Println(err)
		}
		msgs <- msg
		if checkMessage(msg, "payload_uuidv4") {
			done <- "resolved"
			return
		}
		if checkMessage(msg, "expired") {
			done <- "expired"
			return
		}
		if checkMessage(msg, "message") {
			if msg["message"] == "..." {
				done <- "payloadUuidError"
				return
			}
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
