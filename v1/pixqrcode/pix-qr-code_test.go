package pixqrcode_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	pixqrcode "github.com/AbacatePay/abacatepay-go-sdk/v1/pixqrcode"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Create new client with valid params", func(t *testing.T) {
		// Create a mock HTTP client for testing
		mockClient, err := fetch.New("test-key", "https://api.test.com", "1.0.0", 30*time.Second)
		assert.NoError(t, err)

		client := pixqrcode.New(mockClient)
		assert.NotNil(t, client)
		assert.NotNil(t, client.GetHttpClient())
	})

	t.Run("Create new client with nil params", func(t *testing.T) {
		client := pixqrcode.New(nil)
		assert.NotNil(t, client)
		assert.Nil(t, client.GetHttpClient())
	})
}

func TestCreatePixQrCode(t *testing.T) {
	// Create a mock server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"id": "test-id",
				"amount": 100,
				"status": "pending",
				"devMode": true,
				"brCode": "test-br-code",
				"brCodeBase64": "test-base64",
				"platformFee": 0,
				"createdAt": "2023-01-01T00:00:00Z",
				"updatedAt": "2023-01-01T00:00:00Z",
				"expiresAt": "2023-01-01T01:00:00Z"
			}
		}`))
	}))
	defer server.Close()

	// Create a mock HTTP client with the test server URL
	mockClient, err := fetch.New("test-key", server.URL, "1.0.0", 30*time.Second)
	assert.NoError(t, err)

	var tt = []struct {
		name     string
		body     *pixqrcode.CreatePixQrCode
		hasError bool
	}{
		{
			name: "Valid body",
			body: &pixqrcode.CreatePixQrCode{
				Amount:      decimal.NewFromFloat(100.00), // Amount in cents (100.00 BRL)
				Description: "Test Payment",
			},
			hasError: false,
		},
		{
			name: "Invalid body - missing amount",
			body: &pixqrcode.CreatePixQrCode{
				Description: "Test Payment",
			},
			hasError: true,
		},
		{
			name: "Invalid body - zero amount",
			body: &pixqrcode.CreatePixQrCode{
				Amount:      decimal.Zero,
				Description: "Test Payment",
			},
			hasError: true,
		},
		{
			name: "Invalid body - negative amount",
			body: &pixqrcode.CreatePixQrCode{
				Amount:      decimal.NewFromFloat(-100.00),
				Description: "Test Payment",
			},
			hasError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client := pixqrcode.New(mockClient)
			resp, err := client.CreatePixQrCode(context.Background(), tc.body)
			if tc.hasError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
				assert.Equal(t, "test-id", resp.Data.Id)
				expectedAmount := decimal.NewFromFloat(100.00)
				assert.True(t, resp.Data.Amount.Equal(expectedAmount), "Expected amount %s, got %s", expectedAmount.String(), resp.Data.Amount.String())
			}
		})
	}
}

func TestSimulatePayment(t *testing.T) {
	// Create a mock server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"id": "test-id",
				"amount": 100,
				"status": "paid",
				"devMode": true,
				"brCode": "test-br-code",
				"brCodeBase64": "test-base64",
				"platformFee": 0,
				"createdAt": "2023-01-01T00:00:00Z",
				"updatedAt": "2023-01-01T00:00:00Z",
				"expiresAt": "2023-01-01T01:00:00Z"
			}
		}`))
	}))
	defer server.Close()

	// Create a mock HTTP client with the test server URL
	mockClient, err := fetch.New("test-key", server.URL, "1.0.0", 30*time.Second)
	assert.NoError(t, err)

	var tt = []struct {
		name     string
		body     *pixqrcode.SimulatePayment
		hasError bool
	}{
		{
			name: "Valid body",
			body: &pixqrcode.SimulatePayment{
				Id: "test-id",
			},
			hasError: false,
		},
		{
			name: "Invalid body - missing ID",
			body: &pixqrcode.SimulatePayment{
				Id: "",
			},
			hasError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client := pixqrcode.New(mockClient)
			resp, err := client.SimulatePayment(context.Background(), tc.body)
			if tc.hasError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
				assert.Equal(t, "test-id", resp.Data.Id)
				expectedAmount := decimal.NewFromFloat(100.00)
				assert.True(t, resp.Data.Amount.Equal(expectedAmount), "Expected amount %s, got %s", expectedAmount.String(), resp.Data.Amount.String())
			}
		})
	}
}

func TestCheckStatus(t *testing.T) {
	// Create a mock server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
				"data": {
					"status": "PENDING",
					"expiresAt": "2025-03-25T21:50:20.772Z"
				},
				"error": null
				}`))
	}))
	defer server.Close()

	// Create a mock HTTP client with the test server URL
	mockClient, err := fetch.New("test-key", server.URL, "1.0.0", 30*time.Second)
	assert.NoError(t, err)

	var tt = []struct {
		name     string
		id       string
		hasError bool
	}{
		{
			name:     "Valid ID",
			id:       "test-id",
			hasError: false,
		},
		{
			name:     "Invalid ID - empty",
			id:       "",
			hasError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client := pixqrcode.New(mockClient)
			resp, err := client.CheckStatus(context.Background(), tc.id)
			if tc.hasError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
				assert.Equal(t, "PENDING", resp.Data.Status)
				assert.Equal(t, "2025-03-25T21:50:20.772Z", resp.Data.ExpiresAt)
			}
		})
	}
}
