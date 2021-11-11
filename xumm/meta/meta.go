package meta

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type MetaInterface interface {
	Ping() (*models.Pong, error)
	CuratedAssets() (*models.CuratedAssetsResponse, error)
	KycAccountStatus(a string) (*models.KycAccountStatusResponse, error)
	KycStatusState(body models.KycStatusStateRequest) (*models.KycStatusStateResponse, error)
	RatesCurrency(cur string) (*models.RatesCurrencyResponse, error)
	XrplTransaction(txid string) (*models.XrpTxResponse, error)
}

type Meta struct {
	Cfg *xumm.Config
}
