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
}

type WSCfg struct {
	BaseURL string
	msgs    []anyjson.AnyJson
}

type Payload struct {
	Cfg   *xumm.Config
	WSCfg WSCfg
}

type payloadOpt func(p *Payload)

func NewPayload(cfg *xumm.Config, opts ...payloadOpt) *Payload {
	p := &Payload{
		Cfg: cfg,
	}
	for _, opt := range opts {
		opt(p)
	}

	if p.WSCfg.BaseURL == "" {
		p.WSCfg.BaseURL = WEBSOCKETBASEURL
	}
	return p
}

func WithWSBaseUrl(url string) func(p *Payload) {
	return func(p *Payload) {
		p.WSCfg.BaseURL = url
	}
}
