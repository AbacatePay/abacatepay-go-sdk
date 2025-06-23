package abacatepay

import (
	"errors"
	"os"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
)

const Version = "v0.1.0"

const DefaultTimeout = 500 * time.Millisecond

type Client struct {
	httpClient *fetch.Fetch
	Billing    *billing.Billing
}

type ClientConfig struct {
	URL     string
	APIKey  string
	Timeout time.Duration
}

type RequestOptions struct {
	Timeout time.Duration
	Headers map[string]string
}

var (
	ErrInvalidAPIKey = errors.New("invalid API key")
)

func New(config *ClientConfig) (*Client, error) {
	if config == nil || config.APIKey == "" {
		return nil, ErrInvalidAPIKey
	}

	APIUrl := os.Getenv("ABACATEPAY_API_URL")
	if config.URL != "" {
		APIUrl = config.URL
	}
	if APIUrl == "" {
		APIUrl = "https://api.abacatepay.com"
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	httpClient, err := fetch.New(config.APIKey, APIUrl, Version, timeout)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: httpClient,
		Billing:    billing.New(httpClient),
	}, nil
}
