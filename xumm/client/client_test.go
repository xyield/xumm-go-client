// +build unit

package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/meta"
	"github.com/xyield/xumm-go-client/xumm/payload"
	"github.com/xyield/xumm-go-client/xumm/storage"
	"github.com/xyield/xumm-go-client/xumm/xapp"
)

func TestClientCreation(t *testing.T) {
	cfg, _ := xumm.NewConfig()
	tt := []struct {
		description string
		inputOpts   []clientOpt
		expected    *Client
	}{
		{
			description: "No input opts, base configuration",
			inputOpts:   nil,
			expected:    &Client{Config: cfg, Storage: &storage.Storage{Cfg: cfg}, Meta: &meta.Meta{Cfg: cfg}, Payload: &payload.Payload{Cfg: cfg, WSCfg: payload.WSCfg{BaseURL: payload.WEBSOCKETBASEURL}}, Xapp: &xapp.Xapp{Cfg: cfg}},
		},
		{
			description: "Payload input opt",
			inputOpts:   []clientOpt{WithPayload(payload.NewPayload(cfg, payload.WithWSBaseUrl("testUrl")))},
			expected:    &Client{Config: cfg, Storage: &storage.Storage{Cfg: cfg}, Meta: &meta.Meta{Cfg: cfg}, Payload: &payload.Payload{Cfg: cfg, WSCfg: payload.WSCfg{BaseURL: "testUrl"}}, Xapp: &xapp.Xapp{Cfg: cfg}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			s := New(cfg, tc.inputOpts...)

			assert.Equal(t, tc.expected, s)
		})
	}
}
