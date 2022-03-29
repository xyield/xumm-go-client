package payload

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/xumm"
)

func TestNewPayload(t *testing.T) {
	cfg, _ := xumm.NewConfig(xumm.WithAuth("testApiKey", "testApiSecret"))
	tt := []struct {
		description string
		inputOpts   []payloadOpt
		expected    *Payload
	}{
		{
			description: "test payload creation with no opts",
			inputOpts:   nil,
			expected: &Payload{
				Cfg: cfg,
				WSCfg: WSCfg{
					BaseURL: "wss://xumm.app/sign/",
				},
			},
		},
		{
			description: "test payload creation with different base ws url",
			inputOpts:   []payloadOpt{WithWSBaseUrl("testUrl")},
			expected: &Payload{
				Cfg: cfg,
				WSCfg: WSCfg{
					BaseURL: "testUrl",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			actual := NewPayload(cfg, tc.inputOpts...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
