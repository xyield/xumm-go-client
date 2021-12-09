package xumm

import (
	"errors"
	"log"
	"net/http"
	"os"
)

const (
	BASEURLV1 = "https://xumm.app/api/v1"
	// #nosec G101 -- This is a false positive
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

type CredentialOverrideError struct{}

func (c CredentialOverrideError) Error() string {
	return "Cannot override secret and key credentials - use WithAuth to manually set these."
}

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

func (cfg *Config) AddHeader(key, value string) error {

	if key == "X-API-Key" || key == "X-API-Secret" {
		log.Println("It is not possible to override X-API-key or X-API-Secret headers, please use WithAuth() to manually set these.")
		return CredentialOverrideError{}
	}
	cfg.headers[key] = []string{value}
	return nil
}

func (cfg *Config) GetHeaders() map[string][]string {
	return cfg.headers
}
