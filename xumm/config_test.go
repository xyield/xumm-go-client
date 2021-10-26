package xumm

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1, apiKey: "testApiKey", apiSecret: "testApiSecret"}, cfg)
		assert.NoError(t, err)
	})
}
