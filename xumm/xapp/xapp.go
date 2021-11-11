package xapp

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type XappInterface interface {
	GetXappOtt(t string) (*models.XappResponse, error)
}

type Xapp struct {
	Cfg *xumm.Config
}
