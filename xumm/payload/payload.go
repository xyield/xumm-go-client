package payload

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type PayloadInterface interface {
	PostPayload(models.XummPostPayload) (*models.XummPostPayloadResponse, error)
}

type Payload struct {
	Cfg *xumm.Config
}
