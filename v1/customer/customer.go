package customer

import (
	"context"

	"github.com/antunesgabriel/abacatepay-go-sdk/internal/pkg/fetch"
)


type CustomerClient struct {
	HttpClient *fetch.Fetch
}

func New(httpClient *fetch.Fetch) *CustomerClient {
	return &CustomerClient{
		HttpClient: httpClient,
	}
}

func (c *CustomerClient) Create(
	ctx context.Context,
	body *CreateCustomerBody,
) (*CreateCustomerResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	var response CreateCustomerResponse

	resp, err := c.HttpClient.Post(ctx, "/v1/customer/create", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *CustomerClient) ListAll(ctx context.Context) (*ListCustomerResponse, error) {
	var response ListCustomerResponse

	resp, err := c.HttpClient.Get(ctx, "/v1/customer/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
