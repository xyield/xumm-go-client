package client

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/storage"
)

type Client struct {
	Config  *xumm.Config
	Storage storage.StorageInterface
}

func New(cfg *xumm.Config) *Client {
	return &Client{
		Config: cfg,
		Storage: &storage.Storage{
			Cfg: cfg,
		},
	}
}
