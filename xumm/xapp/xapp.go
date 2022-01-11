package xapp

import (
	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type XappInterface interface {
	GetXappOtt(t string) (*models.XappOttResponse, error)
}

type Xapp struct {
	Cfg *xumm.Config
}
