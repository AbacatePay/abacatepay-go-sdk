package pixqrcode

import (
	"context"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
)

type PixQrCode struct {
	HttpClient *fetch.Fetch
}

func New(httpClient *fetch.Fetch) *PixQrCode {
	return &PixQrCode{
		HttpClient: httpClient,
	}
}

func (p *PixQrCode) GetHttpClient() *fetch.Fetch {
	return p.HttpClient
}

func (p *PixQrCode) CreatePixQrCode(ctx context.Context,
	body *CreatePixQrCode,
) (*CreatePixQrCodeResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	resp, err := p.HttpClient.Post(ctx, "/v1/pixQrCode/create", body)
	if err != nil {
		return nil, err
	}

	var response CreatePixQrCodeResponse
	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *PixQrCode) SimulatePayment(ctx context.Context,
	body *SimulatePayment,
) (*SimulatePaymentResponse, error) {

	if body.Id == "" {
		return nil, &InvalidPixQrCodeIdError{ID: body.Id}
	}

	resp, err := p.HttpClient.Post(ctx, "/v1/pixQrCode/simulate-payment", body)
	if err != nil {
		return nil, err
	}

	var response SimulatePaymentResponse
	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *PixQrCode) CheckStatus(ctx context.Context,
	id string,
) (*CheckStatusResponse, error) {

	if id == "" {
		return nil, &InvalidPixQrCodeIdError{ID: id}
	}

	resp, err := p.HttpClient.Get(ctx, "/v1/pixQrCode/check?id="+id)
	if err != nil {
		return nil, err
	}

	var response CheckStatusResponse
	err = fetch.ParseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
