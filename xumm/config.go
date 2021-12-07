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
	ApiKey     string
	ApiSecret  string
	headers    map[string][]string
}

type ConfigOpt func(cfg *Config)

func NewConfig(opts ...ConfigOpt) (*Config, error) {

	cfg := &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1}

	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.ApiKey == "" || cfg.ApiSecret == "" {

		apiKey, err := getAuthEnv("XUMM_API_KEY")

		if err != nil {
			return cfg, err
		}

		cfg.ApiKey = apiKey

		apiSecret, err := getAuthEnv("XUMM_API_SECRET")

		if err != nil {
			return cfg, err
		}
		cfg.ApiSecret = apiSecret

		cfg.headers = map[string][]string{
			"X-API-Key":    {apiKey},
			"X-API-Secret": {apiSecret},
			"Content-Type": {"application/json"},
		}
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

func WithHttpClient(c HTTPClient) ConfigOpt {
	return func(cfg *Config) {
		cfg.HTTPClient = c
	}
}

func WithAuth(key, secret string) ConfigOpt {
	return func(cfg *Config) {
		cfg.ApiKey = key
		cfg.ApiSecret = secret
		cfg.headers = map[string][]string{
			"X-API-Key":    {key},
			"X-API-Secret": {secret},
			"Content-Type": {"application/json"},
		}
	}
}

func (cfg *Config) AddHeader(key, value string) {

	newHeaders := make(map[string][]string)
	h := cfg.headers
	for k, v := range h {
		newHeaders[k] = v
	}
	newHeaders[key] = []string{value}

	cfg.headers = newHeaders
}

func (cfg *Config) GetHeaders() map[string][]string {
	return cfg.headers
}
