package meta

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type MetaInterface interface {
	Ping() (*models.Pong, error)
	GetCuratedAssets() (*models.CuratedAssetsResponse, error)
	GetKycStatusByAccount(a string) (*models.KycStatusByAccountResponse, error)
	GetKycStatusByUserToken(body models.KycStatusByUserTokenRequest) (*models.KycStatusByUserTokenResponse, error)
	GetRatesForCurrency(cur string) (*models.RatesCurrencyResponse, error)
	GetXrplTransaction(txid string) (*models.XrpTxResponse, error)
}

type Meta struct {
	Cfg *xumm.Config
}
