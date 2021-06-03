package xumm

import (
	"errors"
	"log"
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

// type httpClient struct {
// 	client    http.Client
// 	apiKey    string
// 	apiSecret string
// }

// func (c *httpClient) setApiKey(v string) {
// 	c.apiKey = v
// }

// func (c *httpClient) setApiSecret(v string) {
// 	c.apiSecret = v
// }

// func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Add("X-API-Key", c.apiKey)
// 	req.Header.Add("X-API-Secret", c.apiSecret)
// 	return c.client.Do(req)
// }

type SDK struct {
	BaseURL    string
	apiKey     string
	apiSecret  string
	HTTPClient HTTPClient
}

type SDKOpt func(c *SDK)

func NewClient(opts ...SDKOpt) (*SDK, error) {
	c := &SDK{
		BaseURL:    BASEURLV1,
		HTTPClient: &http.Client{},
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.apiKey == "" || c.apiSecret == "" {
		// Check Environment for variables if non set
		apiKey, err := getAuthEnv("XUMM_API_KEY")
		if err != nil {
			return c, err
		}
		c.apiKey = apiKey
		apiSecret, err := getAuthEnv("XUMM_API_SECRET")
		if err != nil {
			return c, err
		}
		c.apiSecret = apiSecret
	}
	return c, nil
}

func getAuthEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)

	if !ok {
		log.Printf("%v was not set", key)
		return "", errors.New(XUMMCREDENTIALSNOTSET)
	}
	return v, nil
}

func WithAuth(apiKey, apiSecret string) SDKOpt {
	return func(c *SDK) {
		c.apiKey = apiKey
		c.apiSecret = apiSecret
	}
}

func WithHttpClient(h HTTPClient) SDKOpt {
	return func(c *SDK) {
		c.HTTPClient = h
	}
}

func (s *SDK) SetXummHeaders(req *http.Request) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-Key", s.apiKey)
	req.Header.Add("X-API-Secret", s.apiSecret)
	return req
}
