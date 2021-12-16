package meta

import (
	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/models"
)

type MetaInterface interface {
	Ping() (*models.Pong, error)
	GetCuratedAssets() (*models.CuratedAssetsResponse, error)
	KycAccountStatus(a string) (*models.KycAccountStatusResponse, error)
	KycStatusState(body models.KycStatusStateRequest) (*models.KycStatusStateResponse, error)
	GetRatesForCurrency(cur string) (*models.RatesCurrencyResponse, error)
	GetXrplTransaction(txid string) (*models.XrpTxResponse, error)
}

type Meta struct {
	Cfg *xumm.Config
}
