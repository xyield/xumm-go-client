//go:build integration
// +build integration

package payload

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
	xummdevice "github.com/xyield/xumm-user-device"
)

func TestPostPayloadIntegration_SignRequest(t *testing.T) {
	xd := xummdevice.NewUserDevice(os.Getenv("XUMM_USER_DEVICE_ACCESS_TOKEN"), os.Getenv("XUMM_USER_DEVICE_UID"))
	cfg, _ := xumm.NewConfig()

	p := &Payload{
		Cfg: cfg,
	}

	cp, err := p.PostPayload(models.XummPostPayload{
		TxJson: anyjson.AnyJson{
			"TransactionType": "Payment",
			"Account":         "rQNrSWi3t6ojNFof8gE3Wq8Pwz88QUr6Hx",
			"Amount":          "1",
			"Destination":     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
			"Fee":             "12",
		},
	})
	if err != nil {
		log.Println("Error creating payload", err)
	}
	err = xd.OpenPayload(cp.UUID)
	if err != nil {
		log.Println("Error opening payload:", err)
	}
	payload, err := p.GetPayloadByUUID(cp.UUID)
	if err != nil {
		log.Println("Error fetching payload", err)
	}
	assert.Equal(t, true, payload.Meta.AppOpened)

	err = xd.SignRequest(cp.UUID, "Payment")
	assert.NoError(t, err)

	payload, err = p.GetPayloadByUUID(cp.UUID)
	assert.NoError(t, err)
	assert.Equal(t, true, payload.Meta.Signed)
	assert.Equal(t, true, payload.Meta.Resolved)
	assert.Equal(t, true, payload.Meta.OpenedByDeeplink)
}

func TestPostPayloadIntegration_RejectRequest(t *testing.T) {
	xd := xummdevice.NewUserDevice(os.Getenv("XUMM_USER_DEVICE_ACCESS_TOKEN"), os.Getenv("XUMM_USER_DEVICE_UID"))
	cfg, _ := xumm.NewConfig()

	p := &Payload{
		Cfg: cfg,
	}

	cp, err := p.PostPayload(models.XummPostPayload{
		TxJson: anyjson.AnyJson{
			"TransactionType": "Payment",
			"Account":         "rQNrSWi3t6ojNFof8gE3Wq8Pwz88QUr6Hx",
			"Amount":          "1",
			"Destination":     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
			"Fee":             "12",
		},
	})
	if err != nil {
		log.Println("Error creating payload", err)
	}

	err = xd.OpenPayload(cp.UUID)
	if err != nil {
		log.Println("Error opening payload:", err)
	}
	payload, err := p.GetPayloadByUUID(cp.UUID)
	if err != nil {
		log.Println("Error fetching payload", err)
	}
	assert.Equal(t, true, payload.Meta.AppOpened)

	err = xd.RejectRequest(cp.UUID)
	assert.NoError(t, err)
	payload, err = p.GetPayloadByUUID(cp.UUID)
	assert.NoError(t, err)
	assert.Equal(t, false, payload.Meta.Signed)
	assert.Equal(t, true, payload.Meta.Resolved)
}

func TestSubscribeSignRequestIntegration(t *testing.T) {

	xd := xummdevice.NewUserDevice(os.Getenv("XUMM_USER_DEVICE_ACCESS_TOKEN"), os.Getenv("XUMM_USER_DEVICE_UID"))
	cfg, _ := xumm.NewConfig()

	p := &Payload{
		Cfg: cfg,
		WSCfg: WSCfg{
			baseUrl: WEBSOCKETBASEURL,
		},
	}

	cp, err := p.PostPayload(models.XummPostPayload{
		TxJson: anyjson.AnyJson{
			"TransactionType": "Payment",
			"Account":         "rQNrSWi3t6ojNFof8gE3Wq8Pwz88QUr6Hx",
			"Amount":          "1",
			"Destination":     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
			"Fee":             "12",
		},
	})
	assert.NoError(t, err)

	go p.Subscribe(cp.UUID)
	time.Sleep(2 * time.Second)

	msg := fmt.Sprint(p.WSCfg.msgs[0])
	assert.Contains(t, msg, "message:Welcome")

	xd.OpenPayload(cp.UUID)

	msg = fmt.Sprint(p.WSCfg.msgs[2])
	assert.Contains(t, msg, "opened:true")

	xd.SignRequest(cp.UUID, transactionTypeToString[Payment])
	time.Sleep(2 * time.Second)

	msg = fmt.Sprint(p.WSCfg.msgs[3])
	assert.Contains(t, msg, "opened_by_deeplink:true", "signed:true", "user_token:true")

	time.Sleep(2 * time.Second)
	payload, err := p.GetPayloadByUUID(cp.UUID)
	if err != nil {
		log.Println("Error fetching payload", err)
	}

	assert.Equal(t, true, payload.Meta.AppOpened)
	assert.Equal(t, true, payload.Meta.OpenedByDeeplink)
	assert.Equal(t, true, payload.Meta.Signed)
	assert.Equal(t, true, payload.Meta.Resolved)

}

func TestSubscribeRejectRequestIntegration(t *testing.T) {
	xd := xummdevice.NewUserDevice(os.Getenv("XUMM_USER_DEVICE_ACCESS_TOKEN"), os.Getenv("XUMM_USER_DEVICE_UID"))
	cfg, _ := xumm.NewConfig()

	p := &Payload{
		Cfg: cfg,
		WSCfg: WSCfg{
			baseUrl: WEBSOCKETBASEURL,
		},
	}

	cp, err := p.PostPayload(models.XummPostPayload{
		TxJson: anyjson.AnyJson{
			"TransactionType": "Payment",
			"Account":         "rQNrSWi3t6ojNFof8gE3Wq8Pwz88QUr6Hx",
			"Amount":          "1",
			"Destination":     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
			"Fee":             "12",
		},
	})
	assert.NoError(t, err)

	go p.Subscribe(cp.UUID)
	time.Sleep(2 * time.Second)

	msg := fmt.Sprint(p.WSCfg.msgs[0])
	assert.Contains(t, msg, "message:Welcome")

	xd.OpenPayload(cp.UUID)

	msg = fmt.Sprint(p.WSCfg.msgs[2])
	assert.Contains(t, msg, "opened:true")

	xd.RejectRequest(cp.UUID)
	time.Sleep(2 * time.Second)

	msg = fmt.Sprint(p.WSCfg.msgs[3])
	assert.Contains(t, msg, "opened_by_deeplink:true", "signed:false", "user_token:true")

	payload, err := p.GetPayloadByUUID(cp.UUID)
	if err != nil {
		log.Println("Error fetching payload", err)
	}

	assert.Equal(t, true, payload.Meta.AppOpened)
	assert.Equal(t, true, payload.Meta.OpenedByDeeplink)
	assert.Equal(t, false, payload.Meta.Signed)
	assert.Equal(t, true, payload.Meta.Resolved)

}

// func TestCreateAndSubscribeIntegration(t *testing.T) {
// xd := xummdevice.NewUserDevice(os.Getenv("XUMM_USER_DEVICE_ACCESS_TOKEN"), os.Getenv("XUMM_USER_DEVICE_UID"))
// 	cfg, _ := xumm.NewConfig()

// 	p := &Payload{
// 		Cfg: cfg,
// 		WSCfg: WSCfg{
// 			baseUrl: WEBSOCKETBASEURL,
// 		},
// 	}

// 	p.CreateAndSubscribe(models.XummPostPayload{
// 		TxJson: anyjson.AnyJson{
// 			"TransactionType": "Payment",
// 			"Account":         "rQNrSWi3t6ojNFof8gE3Wq8Pwz88QUr6Hx",
// 			"Amount":          "1",
// 			"Destination":     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
// 			"Fee":             "12",
// 		},
// 	})

// }
