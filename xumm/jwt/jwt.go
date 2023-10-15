package jwt

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type JwtInterface interface {
	JWTGetAuthorization() (*models.XappJWTAuthorizeResponse, error)
	CancelPayloadByUUID(uuid string, jwt ...string) (*models.XummDeletePayloadResponse, error)
	GetCuratedAssets(jwt ...string) (*models.CuratedAssetsResponse, error)
	GetPayloadByUUID(uuid string, jwt ...string) (*models.XummPayload, error)
	GetPayloadByCustomId(customId string, jwt ...string) (*models.XummPayload, error)
	GetRatesForCurrency(cur string, jwt ...string) (*models.RatesCurrencyResponse, error)
	PostPayload(body models.XummPostPayload, jwt ...string) (*models.CreatedPayload, error)
	Ping(Jwt ...string) (*models.Pong, error)
}

type Jwt struct {
	Cfg *xumm.Config
}
