package billing

import (
	"context"

	"github.com/antunesgabriel/abacatepay-go-sdk/internal/pkg/fetch"
)

type Billing struct {
	HttpClient *fetch.Fetch
}

func New(httpClient *fetch.Fetch) *Billing {
	return &Billing{
		HttpClient: httpClient,
	}
}

func (b *Billing) Create(
	ctx context.Context,
	body *CreateBillingBody,
) (*CreateBillingResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	var response CreateBillingResponse

	resp, err := b.HttpClient.Post(ctx, "/v1/billing/create", body)
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

func (b *Billing) ListAll(
	ctx context.Context,
	page int,
	limit int,
) (*ListBillingResponse, error) {
	var response ListBillingResponse

	resp, err := b.HttpClient.Get(ctx, "/v1/billing/list")
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
