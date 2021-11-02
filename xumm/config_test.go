package xumm

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
)

func TestXummConfigCreation(t *testing.T) {
	t.Run("Default config creation with no env", func(t *testing.T) {
		cfg, err := NewConfig()
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1}, cfg)
		assert.EqualError(t, err, XUMMCREDENTIALSNOTSET)
	})
	t.Run("Default config creation with env", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testApiKey")
		os.Setenv("XUMM_API_SECRET", "testApiSecret")

		cfg, err := NewConfig()
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, ApiKey: "testApiKey", ApiSecret: "testApiSecret", Headers: map[string][]string{"XUMM_API_KEY": {"testApiKey"}, "XUMM_API_SECRET": {"testApiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.NoError(t, err)
	})
	t.Run("Custom http client with env", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testApiKey")
		os.Setenv("XUMM_API_SECRET", "testApiSecret")
		mockClient := &testutils.MockClient{}
		cfg, err := NewConfig(WithHttpClient(mockClient))

		assert.Equal(t, &Config{HTTPClient: mockClient, BaseURL: BASEURLV1, ApiKey: "testApiKey", ApiSecret: "testApiSecret", Headers: map[string][]string{"XUMM_API_KEY": {"testApiKey"}, "XUMM_API_SECRET": {"testApiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.NoError(t, err)
	})
	t.Run("Manually set apikey and secret", func(t *testing.T) {
		cfg, err := NewConfig(WithAuth("manualApiKey", "manualApiSecret"))
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, ApiKey: "manualApiKey", ApiSecret: "manualApiSecret", Headers: map[string][]string{"XUMM_API_KEY": {"manualApiKey"}, "XUMM_API_SECRET": {"manualApiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.NoError(t, err)
	})
}
