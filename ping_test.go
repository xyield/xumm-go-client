package xumm

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/xyield/xumm-go-client/models"
)

func TestPingEndpoint(t *testing.T) {
	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")
	pong := &models.Pong{
		Pong: true,
		Auth: models.ApplicationDetails{
			Quota: map[string]interface{}{},
			Application: models.Application{
				UUIDV4:     uuid.New().String(),
				Name:       "test-application",
				WebhookUrl: "https://test-webhook",
				Disabled:   0,
			},
			Call: models.Call{
				UUIDV4: uuid.New().String(),
			},
		},
	}
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			b, _ := json.Marshal(pong)
			r := ioutil.NopCloser(bytes.NewReader(b))
			return &http.Response{StatusCode: 200, Body: r}, nil
		},
	}
	c, err := NewClient(WithHttpClient(mockClient))
	assert.NoError(t, err)
	res, err := c.Ping()
	assert.NoError(t, err)
	assert.Equal(t, pong, res)
}
