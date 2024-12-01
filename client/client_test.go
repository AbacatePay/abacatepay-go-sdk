package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antunesgabriel/abacatepay-go-sdk/client"
)

func TestNew(t *testing.T) {
	t.Run("Create new client with valid params", func(t *testing.T) {
		client := client.New(&client.ClientConfig{
			ApiKey:  "test-key",
			Timeout: 10,
		})
		assert.NotNil(t, client)
	})

	t.Run("Panic if API key is empty", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			assert.Contains(t, r, "API key is required")
		}()
		client.New(&client.ClientConfig{
			ApiKey:  "",
			Timeout: 10,
		})
	})
}
