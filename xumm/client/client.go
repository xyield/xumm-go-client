package client

import (
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/meta"
	"github.com/xyield/xumm-go-client/xumm/payload"
	"github.com/xyield/xumm-go-client/xumm/storage"
	"github.com/xyield/xumm-go-client/xumm/xapp"
)

// Client struct is used to interact with the XUMM api.
type Client struct {
	Config  *xumm.Config
	Storage storage.StorageInterface
	Meta    meta.MetaInterface
	Payload payload.PayloadInterface
	Xapp    xapp.XappInterface
}

type clientOpt func(c *Client)

// New creates a new Client object for interacting with the XUMM api.
func New(cfg *xumm.Config, opts ...clientOpt) *Client {
	c := &Client{
		Config: cfg,
		Storage: &storage.Storage{
			Cfg: cfg,
		},
		Meta: &meta.Meta{
			Cfg: cfg,
		},
		Payload: payload.NewPayload(cfg),
		Xapp: &xapp.Xapp{
			Cfg: cfg,
		},
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithPayload(p payload.PayloadInterface) func(c *Client) {
	return func(c *Client) {
		c.Payload = p
	}
}
