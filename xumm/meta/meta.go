package meta

import "github.com/xyield/xumm-go-client/xumm"

type MetaInterface interface {
	Ping()
	CuratedAssets()
	KycAccountStatus()
	KycStatusState()
	RatesCurrency()
	XrplTransaction()
}

type Meta struct {
	Cfg *xumm.Config
}
