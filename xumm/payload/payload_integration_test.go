// +build integration

package payload

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
	xummdevice "github.com/xyield/xumm-user-device"
)

func TestPostPayloadIntegration(t *testing.T) {
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
			// Sequence: account_data.Sequence
		},
	})
	if err != nil {
		log.Println("Error creating payload", err)
	}
	// utils.PrettyPrintJson(cp)
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
}
