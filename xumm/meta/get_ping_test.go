package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestPingEndpoint(t *testing.T) {
	expected := &models.Pong{
		Pong: true,
		Auth: models.ApplicationDetails{
			Quota: map[string]interface{}{},
			Application: models.Application{
				UUIDV4:     "8525e32b-1bd0-4839-af2f-f794874a80b0",
				Name:       "test-application",
				WebhookUrl: "https://test-webhook",
				Disabled:   0,
			},
			Call: models.Call{
				UUIDV4: "4b97cf7a-1837-471f-baed-2ebebcf5adb4",
			},
		},
	}
	json := `{
		"pong": true,
		"auth": {
		  "quota": {},
		  "application": {
			"uuidv4": "8525e32b-1bd0-4839-af2f-f794874a80b0",
			"name": "test-application",
			"webhookurl": "https://test-webhook",
			"disabled": 0
		  },
		  "call": {
			"uuidv4": "4b97cf7a-1837-471f-baed-2ebebcf5adb4"
		  }
		}
	  }`
	m := &testutils.MockClient{}
	m.DoFunc = testutils.MockResponse(json, 200, m)
	cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
	assert.NoError(t, err)
	meta := &Meta{
		Cfg: cfg,
	}
	pong, err := meta.GetPing()
	assert.NoError(t, err)
	assert.Equal(t, http.Header{
		"X-API-Key":    {"testApiKey"},
		"X-API-Secret": {"testApiSecret"},
		"Content-Type": {"application/json"},
	}, m.Spy.Header)
	assert.Equal(t, expected, pong)
}

func TestGetPingEndpointErrorResponse(t *testing.T) {
	json := `{
		"error": {
			"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
			"code": 812
		}
	}`
	m := &testutils.MockClient{}
	m.DoFunc = testutils.MockResponse(json, 403, m)
	cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
	assert.NoError(t, err)
	meta := &Meta{
		Cfg: cfg,
	}
	p, err := meta.GetPing()
	assert.Nil(t, p)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error returned with reference 3a04c7d3-94aa-4d8d-9559-62bb5e8a653c and code 812")
}

func TestGetPingEndpointErrorUnauthorisedResponse(t *testing.T) {
	json := `{
		"error": true,
		"message": "Endpoint unknown or method invalid for given endpoint",
		"reference": "",
		"code": 404,
		"req": "/v1/platform/payload/payload_uuid",
		"method": "GET"
	  }`

	m := &testutils.MockClient{}
	m.DoFunc = testutils.MockResponse(json, 404, m)
	cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
	assert.NoError(t, err)
	meta := &Meta{
		Cfg: cfg,
	}
	p, err := meta.GetPing()
	assert.Nil(t, p)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error returned with code 404, reference '' and message 'Endpoint unknown or method invalid for given endpoint'")
}
