package xumm

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
)

func TestXummConfigCreation(t *testing.T) {
	t.Run("Default config creation with no env", func(t *testing.T) {
		cfg, err := NewConfig()
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1}, cfg)
		assert.EqualError(t, err, XUMMCREDENTIALSNOTSET)
	})
	t.Run("Default config creation with env", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testapiKey")
		os.Setenv("XUMM_API_SECRET", "testapiSecret")

		cfg, err := NewConfig()
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, apiKey: "testapiKey", apiSecret: "testapiSecret", headers: map[string][]string{"X-API-Key": {"testapiKey"}, "X-API-Secret": {"testapiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.NoError(t, err)
	})
	t.Run("Custom http client with env", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testapiKey")
		os.Setenv("XUMM_API_SECRET", "testapiSecret")
		mockClient := &testutils.MockClient{}
		cfg, err := NewConfig(WithHttpClient(mockClient))

		assert.Equal(t, &Config{HTTPClient: mockClient, BaseURL: BASEURLV1, apiKey: "testapiKey", apiSecret: "testapiSecret", headers: map[string][]string{"X-API-Key": {"testapiKey"}, "X-API-Secret": {"testapiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.NoError(t, err)
	})
	t.Run("Manually set apikey and secret", func(t *testing.T) {
		cfg, err := NewConfig(WithAuth("manualapiKey", "manualapiSecret"))
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, apiKey: "manualapiKey", apiSecret: "manualapiSecret", headers: map[string][]string{"X-API-Key": {"manualapiKey"}, "X-API-Secret": {"manualapiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.NoError(t, err)
	})
	t.Run("Set extra headers leaving default intact", func(t *testing.T) {
		cfg, _ := NewConfig(WithAuth("manualapiKey", "manualapiSecret"))
		err := cfg.AddHeader("testKey", "testValue")
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, apiKey: "manualapiKey", apiSecret: "manualapiSecret", headers: map[string][]string{"X-API-Key": {"manualapiKey"}, "X-API-Secret": {"manualapiSecret"}, "Content-Type": {"application/json"}, "testKey": {"testValue"}}}, cfg)
		assert.NoError(t, err)
	})
	t.Run("Set an existing header with no overwrite", func(t *testing.T) {
		cfg, _ := NewConfig(WithAuth("manualapiKey", "manualapiSecret"))
		err := cfg.AddHeader("X-API-Key", "testKey")
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, apiKey: "manualapiKey", apiSecret: "manualapiSecret", headers: map[string][]string{"X-API-Key": {"manualapiKey"}, "X-API-Secret": {"manualapiSecret"}, "Content-Type": {"application/json"}}}, cfg)
		assert.Error(t, err)
		assert.EqualError(t, err, "Cannot override secret and key credentials - use WithAuth to manually set these.")
	})
}
