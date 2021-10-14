package xumm

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
)

func TestXummClientCreation(t *testing.T) {
	t.Run("Default client creation with no env", func(t *testing.T) {
		c, err := NewClient()
		assert.Equal(t, &SDK{BaseURL: BASEURLV1, HTTPClient: &http.Client{}}, c)
		assert.EqualError(t, err, XUMMCREDENTIALSNOTSET)
	})
	t.Run("Default client creation with only api key set", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testApiKey")
		c, err := NewClient()
		assert.Equal(t, &SDK{BaseURL: BASEURLV1, HTTPClient: &http.Client{}, apiKey: "testApiKey"}, c)
		assert.EqualError(t, err, XUMMCREDENTIALSNOTSET)
	})
	t.Run("Default client creation with Env set", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testApiKey")
		os.Setenv("XUMM_API_SECRET", "testApiSecret")
		c, err := NewClient()
		assert.NoError(t, err)
		assert.Equal(t, &SDK{BaseURL: BASEURLV1, HTTPClient: &http.Client{}, apiKey: "testApiKey", apiSecret: "testApiSecret"}, c)
	})
	t.Run("Client Creation With Auth Option", func(t *testing.T) {
		c, err := NewClient(WithAuth("testApiKey", "testApiSecret"))
		assert.NoError(t, err)
		assert.Equal(t, &SDK{BaseURL: BASEURLV1, HTTPClient: &http.Client{}, apiKey: "testApiKey", apiSecret: "testApiSecret"}, c)
	})
	t.Run("Client Creation with http client option", func(t *testing.T) {
		os.Setenv("XUMM_API_KEY", "testApiKey")
		os.Setenv("XUMM_API_SECRET", "testApiSecret")
		c, err := NewClient(WithHttpClient(&testutils.MockClient{}))
		assert.NoError(t, err)
		assert.Equal(t, &SDK{BaseURL: BASEURLV1, HTTPClient: &testutils.MockClient{}, apiKey: "testApiKey", apiSecret: "testApiSecret"}, c)
	})
}
