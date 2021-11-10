package xapp

import (
	"github.com/xyield/xumm-go-client/models"
	"github.com/xyield/xumm-go-client/xumm"
)

type XappInterface interface {
	GetXappOtt(t string) (*models.XappResponse, error)
}

type Xapp struct {
	Cfg *xumm.Config
}
