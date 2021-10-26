package xumm

import (
	"errors"
	"net/http"
	"os"
)

const (
	BASEURLV1             = "https://xumm.app/api/v1"
	XUMMCREDENTIALSNOTSET = "API Key or Secret not set"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	HTTPClient HTTPClient
	BaseURL    string
	apiKey     string
	apiSecret  string
}

func NewConfig() (*Config, error) {

	cfg := &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1}

	if cfg.apiKey == "" || cfg.apiSecret == "" {

		apiKey, err := getAuthEnv("XUMM_API_KEY")

		if err != nil {
			return cfg, err
		}

		cfg.apiKey = apiKey

		apiSecret, err := getAuthEnv("XUMM_API_SECRET")

		if err != nil {
			return cfg, err
		}
		cfg.apiSecret = apiSecret
	}

	return cfg, nil

}

func getAuthEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(XUMMCREDENTIALSNOTSET)
	}
	return v, nil
}
