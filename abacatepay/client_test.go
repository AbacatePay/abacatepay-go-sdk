package abacatepay_test

import (
	"testing"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Create new client with valid params", func(t *testing.T) {
		cl, err := abacatepay.New(&abacatepay.ClientConfig{
			APIKey:  "test-key",
			Timeout: 10 * time.Second,
		})
		assert.NoError(t, err)
		assert.NotNil(t, cl)
	})

	t.Run("Error if API key is empty", func(t *testing.T) {
		cl, err := abacatepay.New(&abacatepay.ClientConfig{
			APIKey:  "",
			Timeout: 10 * time.Second,
		})
		assert.Error(t, err)
		assert.ErrorIs(t, abacatepay.ErrInvalidAPIKey, err)
		assert.Nil(t, cl)
	})
}
