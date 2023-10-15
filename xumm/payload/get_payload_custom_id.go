package payload

import (
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	GETPAYLOADBYCUSTOMIDENDPOINT = "/platform/payload/ci/"
)

// EmptyIdError is returned when an empty string is provided for the payload custom id.
type EmptyIdError struct {
}

func (e *EmptyIdError) Error() string {
	return "Empty custom ID provided."
}

// GetPayloadByCustomId returns the payload details or payload resolve status and result data by custom identifier.
// Takes a single argument of the payload custom id string.
func (p *Payload) GetPayloadByCustomId(customId string) (*models.XummPayload, error) {

	if customId == "" {
		return nil, &EmptyIdError{}
	}

	return GetPayload(p, GETPAYLOADBYCUSTOMIDENDPOINT+customId)
}
