package client

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/meta"
	"github.com/xyield/xumm-go-client/xumm/payload"
	"github.com/xyield/xumm-go-client/xumm/storage"
	"github.com/xyield/xumm-go-client/xumm/xapp"
)

type Client struct {
	Config  *xumm.Config
	Storage storage.StorageInterface
	Meta    meta.MetaInterface
	Payload payload.PayloadInterface
	Xapp    xapp.XappInterface
}

func New(cfg *xumm.Config) *Client {
	return &Client{
		Config: cfg,
		Storage: &storage.Storage{
			Cfg: cfg,
		},
		Meta: &meta.Meta{
			Cfg: cfg,
		},
		Payload: &payload.Payload{
			Cfg: cfg,
		},
		Xapp: &xapp.Xapp{
			Cfg: cfg,
		},
	}
}
