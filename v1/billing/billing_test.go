package billing_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
)

func TestNew(t *testing.T) {
	t.Run("Create new client with valid params", func(t *testing.T) {
		client := billing.New(nil)
		assert.NotNil(t, client)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Should validate body", func(t *testing.T) {
		client := billing.New(nil)

		body := &billing.CreateBillingBody{
			Frequency:     billing.OneTime,
			Methods:       []billing.Method{billing.PIX},
			CompletionUrl: "https://example.com/completion",
		}

		ctx := context.Background()

		response, err := client.Create(ctx, body)

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("Should create new billing", func(t *testing.T) {
		body := &billing.CreateBillingBody{
			Frequency:     billing.OneTime,
			Methods:       []billing.Method{billing.PIX},
			CompletionUrl: "https://example.com/completion",
			ReturnUrl:     "https://example.com/return",
			Products: []*billing.BillingProduct{
				{
					ExternalId:  "pix-1234",
					Name:        "PIX",
					Description: "PIX",
					Quantity:    1,
					Price:       100,
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var bodyRef billing.CreateBillingBody

			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "/v1/billing/create", r.URL.Path)

			defer r.Body.Close()

			json.NewDecoder(r.Body).Decode(&bodyRef)

			assert.Equal(t, *body, bodyRef)

			resp := billing.CreateBillingResponse{
				Billing: billing.CreateBillingResponseItem{
					PublicID: "pix-1234",
					Products: []billing.ProductItem{},
				},
			}

			json.NewEncoder(w).Encode(resp)
		}))

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)

		b := billing.New(client)

		ctx := context.Background()

		response, err := b.Create(ctx, body)

		assert.NoError(t, err)
		assert.NotNil(t, response.Billing)
	})
}

func TestListAll(t *testing.T) {
	t.Run("Should list all billings", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "/v1/billing/list", r.URL.Path)

			resp := billing.ListBillingResponse{
				Billings: []billing.BillingListItem{
					{
						ID:        "pix-1234",
						Metadata:  billing.Metadata{},
						PublicID:  "pix-1234",
						Amount:    0,
						Status:    "",
						DevMode:   false,
						Methods:   []billing.Method{},
						Frequency: "",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Version:   0,
						URL:       "",
						BillingID: "",
						Products:  []billing.ProductItem{},
					},
				},
			}

			json.NewEncoder(w).Encode(resp)
		}))

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)

		b := billing.New(client)

		ctx := context.Background()

		response, err := b.ListAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response.Billings)
	})
}
