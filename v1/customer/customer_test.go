package customer_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antunesgabriel/abacatepay-go-sdk/internal/pkg/fetch"
	"github.com/antunesgabriel/abacatepay-go-sdk/v1/customer"
	"github.com/stretchr/testify/assert"
)

func mockCustomer() *customer.CustomerMetadata {
	return &customer.CustomerMetadata{
		Name: "Seiya de Pegasus",
		Cellphone: "11 4002-8922",
		TaxID: "42066612369",
		Email: "seiya@pegasus.com",
	}
}
func TestCreateCustomer(t *testing.T) {
	t.Run("Should validate body", func(t *testing.T) {
		client := customer.New(nil)

		body := &customer.CreateCustomerBody{

		}

		ctx := context.Background()

		response, err := client.Create(ctx, body)

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("Should create new customer", func(t *testing.T) {
		mockMetadata := mockCustomer()

		body := &customer.CreateCustomerBody{
			Name: mockMetadata.Name,
			Cellphone: mockMetadata.Name,
			TaxID: mockMetadata.TaxID,
			Email: mockMetadata.Email,
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var bodyRef customer.CreateCustomerBody

			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "/v1/customer/create", r.URL.Path)

			defer r.Body.Close()

			json.NewDecoder(r.Body).Decode(&bodyRef)

			assert.Equal(t, *body, bodyRef)

			resp := customer.CreateCustomerResponse{
				Customer: customer.CreateCustomerResponseItem{
					AccountID: "accountID",
					StoreID: "storeID",
					DevMode: true,
					Metadata: *mockMetadata,
				},
			}

			json.NewEncoder(w).Encode(resp)
		}))

		client := fetch.New("test-key", server.URL, "1.0.0", 10)

		c := customer.New(client)

		ctx := context.Background()

		response, err := c.Create(ctx, body)

		assert.NoError(t, err)
		assert.NotNil(t, response.Customer)
	})
}

func TestListAllCustomers(t *testing.T) {
	mockCustomer := mockCustomer()
	t.Run("Should list all customers", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "/v1/customer/list", r.URL.Path)

			resp := customer.ListCustomerResponse{
				Customers: []customer.CustomerResponseItem{
					{
						Metadata: *mockCustomer,
						ID: "id",
					},
				},
			}

			json.NewEncoder(w).Encode(resp)
		}))

		client := fetch.New("test-key", server.URL, "1.0.0", 10)

		c := customer.New(client)

		ctx := context.Background()

		response, err := c.ListAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response.Customers)
	})
}
