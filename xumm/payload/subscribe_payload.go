package payload

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	WEBSOCKETBASEURL = "wss://xumm.app/sign/"
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

type ConnectionError struct {
	UUID string
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("Connection dropped for payload websocket with uuid %v", e.UUID)
}

// Subscribes to payload websocket to receive messages and returns payload if it is resolved
func (p *Payload) Subscribe(uuid string) (*models.XummPayload, error) {

	ws, _, err := websocket.DefaultDialer.Dial(p.WSCfg.BaseURL+uuid, nil)
	if err != nil {
		log.Println("Error connecting to websocket:", err)
		return nil, err
	}

	msgsc := make(chan anyjson.AnyJson)
	done := make(chan string)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go recieveMessage(ws, msgsc, done)

	for {
		select {
		case v := <-msgsc:
			p.WSCfg.msgs = append(p.WSCfg.msgs, v)
		case m := <-done:
			if m == "resolved" {
				return p.GetPayloadByUUID(uuid)
			}
			if m == "expired" {
				return nil, &PayloadExpiredError{UUID: uuid}
			}
			if m == "payloadUuidError" {
				return nil, &PayloadUuidError{UUID: uuid}
			}
		case <-interrupt:
			return nil, &ConnectionError{UUID: uuid}
		}
	}
}

func (p *Payload) CreateAndSubscribe(payloadBody models.XummPostPayload) (*models.XummPayload, error) {

	pb, err := p.PostPayload(payloadBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return p.Subscribe(pb.UUID)
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
