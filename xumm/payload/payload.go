package payload

import (
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type PayloadInterface interface {
	PostPayload(models.XummPostPayload) (*models.CreatedPayload, error)
	GetPayloadByUUID(uuid string) (*models.XummPayload, error)
	GetPayloadByCustomId(customId string) (*models.XummPayload, error)
	CancelPayloadByUUID(uuid string) (*models.XummDeletePayloadResponse, error)
	Subscribe(uuid string) (*models.XummPayload, error)
	CreateAndSubscribe(payloadBody models.XummPostPayload) (*models.XummPayload, error)
}

type WSCfg struct {
	baseUrl string
	msgs    []anyjson.AnyJson
}

type Payload struct {
	Cfg   *xumm.Config
	WSCfg WSCfg
}
