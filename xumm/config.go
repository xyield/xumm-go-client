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

// HTTPClient interface allows users to use a custom http client. The net/*http.Client satisfies this interface.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Config to interact with the XUMM API is stored in this object.
type Config struct {
	HTTPClient HTTPClient
	BaseURL    string
	apiKey     string
	apiSecret  string
	headers    map[string][]string
}

// ConfigOpt functions are used to modify the config object.
type ConfigOpt func(cfg *Config)

// Error returned when user attempts to override required headers.
type CredentialOverrideError struct{}

func (c CredentialOverrideError) Error() string {
	return "Cannot override secret and key credentials - use WithAuth to manually set these."
}

// Function used to create a config object.
// If no arguments are provided reasonable defaults are used.
// These defaults can be overridden by adding arguments of type ConfigOpt.
func NewConfig(opts ...ConfigOpt) (*Config, error) {

	cfg := &Config{HTTPClient: &http.Client{}, BaseURL: BASEURLV1}

	for _, opt := range opts {
		opt(cfg)
	}
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

// This ConfigOpt function allows the setting of a custom http client.
func WithHttpClient(c HTTPClient) ConfigOpt {
	return func(cfg *Config) {
		cfg.HTTPClient = c
	}
}

// This ConfigOpt function allows user to manually set the XUMM api key and secret.
// Otherwise they will be set from environment variables.
func WithAuth(key, secret string) ConfigOpt {
	return func(cfg *Config) {
		cfg.apiKey = key
		cfg.apiSecret = secret
		cfg.headers = map[string][]string{
			"X-API-Key":    {key},
			"X-API-Secret": {secret},
			"Content-Type": {"application/json"},
		}
	}
}

// AddHeader allows the user to add a custom header to all requests to the XUMM api.
func (cfg *Config) AddHeader(key, value string) error {

	if key == "X-API-Key" || key == "X-API-Secret" {
		log.Println("It is not possible to override X-API-key or X-API-Secret headers, please use WithAuth() to manually set these.")
		return CredentialOverrideError{}
	}
	cfg.headers[key] = []string{value}
	return nil
}

// GetHeaders returns all current headers in the config object.
func (cfg *Config) GetHeaders() map[string][]string {
	return cfg.headers
}
