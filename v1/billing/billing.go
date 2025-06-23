package billing

import (
	"context"
	"fmt"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
)

type Billing struct {
	HTTPClient *fetch.Fetch
}

func New(httpClient *fetch.Fetch) *Billing {
	return &Billing{
		HTTPClient: httpClient,
	}
}

func (b *Billing) Create(
	ctx context.Context,
	body *CreateBillingBody,
) (*CreateBillingResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	if body.CustomerID == "" && (body.Customer == nil || body.Customer.Email == "") {
		return nil, fmt.Errorf("customerId or customer.email is required")
	}

	var response CreateBillingResponse

	resp, err := b.HTTPClient.Post(ctx, "/v1/billing/create", body)
	if err != nil {
		return nil, err
	}

	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (b *Billing) ListAll(ctx context.Context) (*ListBillingResponse, error) {
	var response ListBillingResponse

	resp, err := b.HTTPClient.Get(ctx, "/v1/billing/list")
	if err != nil {
		return nil, err
	}

	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
