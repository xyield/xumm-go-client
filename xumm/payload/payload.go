package payload

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type PayloadInterface interface {
	PostPayload(models.XummPostPayload) (*models.XummPostPayloadResponse, error)
	GetPayloadByUuid(uuid string) (*models.PayloadUuidResponse, error)
	GetPayloadByCustomId(customId string) (*models.PayloadUuidResponse, error)
}

type Payload struct {
	Cfg *xumm.Config
}
