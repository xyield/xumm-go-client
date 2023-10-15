package jwt

import (
	"github.com/xyield/xumm-go-client/xumm/models"
	"github.com/xyield/xumm-go-client/xumm/payload"
)

const (
	JWTGETPAYLOADBYCUSTOMIDENDPOINT = "/jwt/payload/ci/"
)

// GetPayloadByCustomId returns the payload details or payload resolve status and result data by custom identifier.
// Takes a single argument of the payload custom id string.
func (j *Jwt) GetPayloadByCustomId(customId string, jwt ...string) (*models.XummPayload, error) {
	if customId == "" {
		return nil, &payload.EmptyIdError{}
	}
	return GetPayload(j, JWTGETPAYLOADBYCUSTOMIDENDPOINT+customId, jwt[0])
}
