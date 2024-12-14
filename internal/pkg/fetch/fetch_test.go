package fetch_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
)

type TestResponse struct {
	Message string `json:"message"`
}

func TestNewFetchClient(t *testing.T) {
	t.Run("Create new client with valid params", func(t *testing.T) {
		client, err := fetch.New("test-key", "https://api.test.com", "1.0.0", 10)
		assert.NotNil(t, client)
		assert.NoError(t, err)
	})

	t.Run("Return error if API key is empty", func(t *testing.T) {
		_, err := fetch.New("", "https://api.test.com", "1.0.0", 10)
		assert.Error(t, err)
		assert.ErrorIs(t, fetch.ErrInvalidAPIKey, err)
	})

	t.Run("Panic if API url is empty", func(t *testing.T) {
		_, err := fetch.New("test-key", "", "1.0.0", 10)
		assert.Error(t, err)
		assert.ErrorIs(t, fetch.ErrInvalidAPIUrl, err)
	})
}

func TestFetchMethods(t *testing.T) {
	t.Run("Have GET method", func(t *testing.T) {
		ctx := context.Background()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			resp := TestResponse{Message: "Success"}
			json.NewEncoder(w).Encode(resp)
		}))

		defer server.Close()

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)

		var response *http.Response
		response, err = client.Get(ctx, "/test")

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Have DELETE method", func(t *testing.T) {
		ctx := context.Background()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			resp := TestResponse{Message: "Success"}
			json.NewEncoder(w).Encode(resp)
		}))

		defer server.Close()

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)

		var response *http.Response
		response, err = client.Delete(ctx, "/test/xpto")

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Have POST method", func(t *testing.T) {
		ctx := context.Background()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			resp := TestResponse{Message: "Success"}
			json.NewEncoder(w).Encode(resp)
		}))

		defer server.Close()

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)

		var response *http.Response

		response, err = client.Post(ctx, "/test/xpto", TestResponse{Message: "Success"})

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Have PUT method", func(t *testing.T) {
		ctx := context.Background()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			resp := TestResponse{Message: "Success"}
			json.NewEncoder(w).Encode(resp)
		}))

		defer server.Close()

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)

		var response *http.Response

		response, err = client.Put(ctx, "/test/xpto", TestResponse{Message: "Success"})

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}

func TestParseResponse(t *testing.T) {
	t.Run("Parse response with success", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"message": "Success"}`)),
		}

		var result TestResponse
		err := fetch.ParseResponse(resp, &result)

		assert.NoError(t, err)
		assert.Equal(t, "Success", result.Message)
	})

	t.Run("Error on invalid status code", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 400,
			Body:       io.NopCloser(bytes.NewBufferString(`{"error": "Bad Request"}`)),
		}

		var result TestResponse
		err := fetch.ParseResponse(resp, &result)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error on request: status 400")
	})

	t.Run("Error with invalid JSON", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"invalid": json}`)),
		}

		var result TestResponse
		err := fetch.ParseResponse(resp, &result)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error on deserializing response")
	})
}

func TestRequestOptions(t *testing.T) {
	t.Run("Configure custom timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(TestResponse{Message: "Success"})
		}))

		defer server.Close()

		client, err := fetch.New("test-key", server.URL, "1.0.0", 10*time.Second)
		assert.NoError(t, err)
		opts := fetch.RequestOptions{
			Timeout: 5 * time.Second,
			Headers: map[string]string{
				"X-Custom-Header": "test-value",
			},
		}

		response, err := client.Get(context.Background(), "/test", opts)
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}
