package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/meta"
	"github.com/xyield/xumm-go-client/xumm/payload"
	"github.com/xyield/xumm-go-client/xumm/storage"
)

func TestClientCreation(t *testing.T) {
	cfg, _ := xumm.NewConfig()
	t.Run("Default SDK creation", func(t *testing.T) {
		s := New(cfg)

		assert.Equal(t, &Client{Config: cfg, Storage: &storage.Storage{Cfg: cfg}, Meta: &meta.Meta{Cfg: cfg}, Payload: &payload.Payload{Cfg: cfg}}, s)
	})
}
