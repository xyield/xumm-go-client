package payload

import (
	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/models"
)

type PayloadInterface interface {
	PostPayload(models.XummPostPayload) (*models.CreatedPayload, error)
	GetPayloadByUUID(uuid string) (*models.XummPayload, error)
	GetPayloadByCustomId(customId string) (*models.XummPayload, error)
	CancelPayloadByUUID(uuid string) (*models.XummDeletePayloadResponse, error)
}

type Payload struct {
	Cfg *xumm.Config
}
