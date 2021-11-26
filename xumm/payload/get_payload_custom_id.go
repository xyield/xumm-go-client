package payload

import (
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	GETPAYLOADBYCUSTOMIDENDPOINT = "/platform/payload/ci/"
)

type EmptyIdError struct {
}

func (e *EmptyIdError) Error() string {
	return "Empty custom ID provided."
}

func (p *Payload) GetPayloadByCustomId(customId string) (*models.PayloadUuidResponse, error) {

	if customId == "" {
		return nil, &EmptyIdError{}
	}

	return GetPayload(p, customId)
}
