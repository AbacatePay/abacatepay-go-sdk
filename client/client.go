package client

import (
	"os"
	"time"

	"github.com/antunesgabriel/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/antunesgabriel/abacatepay-go-sdk/v1/billing"
	"github.com/antunesgabriel/abacatepay-go-sdk/v1/customer"
)

const Version = "v0.1.0"

const DefaultTimeout = 500 * time.Millisecond

type Client struct {
	httpClient *fetch.Fetch
	Billing    *billing.Billing
	Customer   *customer.CustomerClient
}

type ClientConfig struct {
	ApiKey  string
	Timeout time.Duration
}

type RequestOptions struct {
	Timeout time.Duration
	Headers map[string]string
}

func New(config *ClientConfig) *Client {
	if config == nil || config.ApiKey == "" {
		panic("API key is required")
	}

	apiUrl := os.Getenv("ABACATEPAY_API_URL")
	if apiUrl == "" {
		apiUrl = "https://api.abacatepay.com"
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	httpClient := fetch.New(config.ApiKey, apiUrl, Version, timeout)

	return &Client{
		httpClient: httpClient,
		Billing:    billing.New(httpClient),
		Customer:   customer.New(httpClient),
	}
}
