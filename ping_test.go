package xumm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
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
			b, _ := jsoniter.Marshal(pong)
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

func TestPingEndpointErrorResponse(t *testing.T) {
	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")
	json := `{
		"error": {
			"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
			"code": 812
		}
	}`
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
			return &http.Response{StatusCode: 403, Body: b}, nil
		},
	}
	c, _ := NewClient(WithHttpClient(mockClient))
	p, err := c.Ping()
	assert.Nil(t, p)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error returned with reference 3a04c7d3-94aa-4d8d-9559-62bb5e8a653c and code 812")
}
