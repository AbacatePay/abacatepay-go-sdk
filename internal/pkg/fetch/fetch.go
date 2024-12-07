package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ErrInvalidAPIKey = errors.New("invalid API key")
	ErrInvalidAPIUrl = errors.New("invalid API url")
)

type Fetch struct {
	apiKey  string
	apiUrl  string
	version string
	timeout time.Duration
}

type RequestOptions struct {
	Timeout time.Duration
	Headers map[string]string
}

func New(apiKey, apiUrl, version string, timeout time.Duration) (*Fetch, error) {
	if apiKey == "" {
		return nil, ErrInvalidAPIKey
	}

	if apiUrl == "" {
		return nil, ErrInvalidAPIUrl
	}

	return &Fetch{
		apiKey:  apiKey,
		apiUrl:  apiUrl,
		version: version,
		timeout: timeout,
	}, nil
}

func (f *Fetch) Request(ctx context.Context, method, endpoint string, body interface{}, opts ...RequestOptions) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", f.apiUrl, endpoint)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error on serializing request body: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error on creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", f.apiKey))
	req.Header.Set("User-Agent", fmt.Sprintf("AbacatePay-Go-SDK/%s", f.version))

	var timeout time.Duration = f.timeout

	if len(opts) > 0 {
		opt := opts[0]

		if opt.Timeout > 0 {
			timeout = opt.Timeout
		}

		for k, v := range opt.Headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{
		Timeout: timeout,
	}

	return client.Do(req)
}

func (f *Fetch) Get(ctx context.Context, endpoint string, opts ...RequestOptions) (*http.Response, error) {
	return f.Request(ctx, http.MethodGet, endpoint, nil, opts...)
}

func (f *Fetch) Post(ctx context.Context, endpoint string, body interface{}, opts ...RequestOptions) (*http.Response, error) {
	return f.Request(ctx, http.MethodPost, endpoint, body, opts...)
}

func (f *Fetch) Put(ctx context.Context, endpoint string, body interface{}, opts ...RequestOptions) (*http.Response, error) {
	return f.Request(ctx, http.MethodPut, endpoint, body, opts...)
}

func (f *Fetch) Delete(ctx context.Context, endpoint string, opts ...RequestOptions) (*http.Response, error) {
	return f.Request(ctx, http.MethodDelete, endpoint, nil, opts...)
}

func ParseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error on reading response body: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error on request: status %d, body: %s", resp.StatusCode, string(body))
	}

	if target != nil {
		if err := json.Unmarshal(body, target); err != nil {
			return fmt.Errorf("error on deserializing response: %v", err)
		}
	}

	return nil
}
