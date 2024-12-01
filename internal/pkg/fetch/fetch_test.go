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

	"github.com/antunesgabriel/abacatepay-go-sdk/internal/pkg/fetch"
)

type TestResponse struct {
	Message string `json:"message"`
}

func TestNewFetchClient(t *testing.T) {
	t.Run("Create new client with valid params", func(t *testing.T) {
		client := fetch.New("test-key", "https://api.test.com", "1.0.0", 10)
		assert.NotNil(t, client)
	})

	t.Run("Panic if API key is empty", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			assert.Contains(t, r, "API key is required")
		}()
		fetch.New("", "https://api.test.com", "1.0.0", 10)
	})

	t.Run("Panic if API url is empty", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			assert.Contains(t, r, "API url is required")
		}()

		fetch.New("test-key", "", "1.0.0", 10)
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

		client := fetch.New("test-key", server.URL, "1.0.0", 10)

		var response *http.Response
		var err error

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

		client := fetch.New("test-key", server.URL, "1.0.0", 10)

		var response *http.Response
		var err error

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

		client := fetch.New("test-key", server.URL, "1.0.0", 10)

		var response *http.Response
		var err error

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

		client := fetch.New("test-key", server.URL, "1.0.0", 10)

		var response *http.Response
		var err error

		response, err = client.Put(ctx, "/test/xpto", TestResponse{Message: "Success"})

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}

func TestParseResponse(t *testing.T) {
	t.Run("Parse resposta com sucesso", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"message": "Success"}`)),
		}

		var result TestResponse
		err := fetch.ParseResponse(resp, &result)

		assert.NoError(t, err)
		assert.Equal(t, "Success", result.Message)
	})

	t.Run("Erro em status code inválido", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 400,
			Body:       io.NopCloser(bytes.NewBufferString(`{"error": "Bad Request"}`)),
		}

		var result TestResponse
		err := fetch.ParseResponse(resp, &result)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error on request: status 400")
	})

	t.Run("Erro em JSON inválido", func(t *testing.T) {
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
	t.Run("Configurar timeout personalizado", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(TestResponse{Message: "Success"})
		}))

		defer server.Close()

		client := fetch.New("test-key", server.URL, "1.0.0", 10)
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
