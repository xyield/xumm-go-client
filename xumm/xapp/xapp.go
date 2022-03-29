package xapp

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type XappInterface interface {
	GetXappOtt(t string) (*models.XappOttResponse, error)
	PostXappEvent(b models.XappRequest) (*models.XappResponse, error)
	PostXappPush(b models.XappRequest) (*models.XappResponse, error)
}

type Xapp struct {
	Cfg *xumm.Config
}
