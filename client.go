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

type Client struct {
	BaseURL    string
	apiKey     string
	apiSecret  string
	HTTPClient HTTPClient
}

type ClientOpt func(c *Client)

func NewClient(opts ...ClientOpt) (*Client, error) {
	c := &Client{
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

func WithAuth(apiKey, apiSecret string) ClientOpt {
	return func(c *Client) {
		c.apiKey = apiKey
		c.apiSecret = apiSecret
	}
}

func WithHttpClient(h HTTPClient) ClientOpt {
	return func(c *Client) {
		c.HTTPClient = h
	}
}
