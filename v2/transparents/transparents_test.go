package transparents_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"
	abacatepay "github.com/AbacatePay/abacatepay-go-sdk/v2"
	"github.com/AbacatePay/abacatepay-go-sdk/v2/transparents"
)

func newClient(t *testing.T, handler http.Handler) (*abacatepay.Client, func()) {
	t.Helper()
	server := httptest.NewServer(handler)
	client, err := abacatepay.New(abacatepay.ClientConfig{BaseURL: server.URL, APIKey: "test-key", Timeout: time.Second})
	if err != nil {
		server.Close()
		t.Fatal(err)
	}
	return client, server.Close
}

func TestCreateUsesTransparentPIXEnvelope(t *testing.T) {
	client, cleanup := newClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v2/transparents/create" {
			t.Fatalf("request = %s %s", r.Method, r.URL.String())
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Fatalf("authorization = %q", got)
		}
		var body struct {
			Method string `json:"method"`
			Data   struct {
				Amount      uint64         `json:"amount"`
				ExpiresIn   int            `json:"expiresIn"`
				Description string         `json:"description"`
				ExternalID  string         `json:"externalId"`
				Metadata    map[string]any `json:"metadata"`
			} `json:"data"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body.Method != "PIX" || body.Data.Amount != 1234 || body.Data.ExpiresIn != 900 || body.Data.ExternalID != "order-1" || body.Data.Metadata["reference"] != "order-1" {
			t.Fatalf("body = %+v", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"id":"pix-1","amount":1234,"status":"PENDING","brCode":"pix-code","metadata":{"reference":"order-1"}},"success":true}`))
	}))
	defer cleanup()

	expires := uint64(900)
	externalID := "order-1"
	response, err := client.Transparents.Create(context.Background(), transparents.CreateQRCodePIXBody{
		Method: "PIX",
		Data:   transparents.CreateQRCodePIXDataBody{Amount: 1234, ExpiresIn: &expires, ExternalID: &externalID, Metadata: map[string]any{"reference": externalID}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !response.Success.Bool() || response.Data == nil || response.Data.ID != "pix-1" || response.Data.Amount != 1234 {
		t.Fatalf("response = %+v", response)
	}
}

func TestListBuildsFiltersAndReturnsFullTransparent(t *testing.T) {
	client, cleanup := newClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/transparents/list" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		want := url.Values{"externalId": {"order-1"}, "id": {"pix-1"}, "limit": {"1"}}
		if r.URL.Query().Get("externalId") != want.Get("externalId") || r.URL.Query().Get("id") != want.Get("id") || r.URL.Query().Get("limit") != want.Get("limit") {
			t.Fatalf("query = %v", r.URL.Query())
		}
		_, _ = w.Write([]byte(`{"data":[{"id":"pix-1","amount":1234,"status":"PAID","brCode":"pix-code","externalId":"order-1","metadata":{"reference":"order-1"}}],"success":true,"pagination":{"hasMore":false,"next":null,"before":null}}`))
	}))
	defer cleanup()

	response, err := client.Transparents.List(context.Background(), transparents.ListQRCodePIXQuery{ID: "pix-1", ExternalID: "order-1", Limit: 1})
	if err != nil {
		t.Fatal(err)
	}
	if !response.Success.Bool() || len(response.Data) != 1 || response.Data[0].Status != transparents.PaymentStatusPaid || response.Data[0].Amount != 1234 {
		t.Fatalf("response = %+v", response)
	}
}

func TestCheckStatusReturnsPaidState(t *testing.T) {
	client, cleanup := newClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/transparents/check" || r.URL.Query().Get("id") != "pix-1" {
			t.Fatalf("request = %s", r.URL.String())
		}
		_, _ = w.Write([]byte(`{"data":{"id":"pix-1","status":"PAID","expiresAt":"2026-09-25T12:00:00Z"},"success":true}`))
	}))
	defer cleanup()

	response, err := client.Transparents.CheckStatus(context.Background(), "pix-1")
	if err != nil {
		t.Fatal(err)
	}
	if !response.Success.Bool() || response.Data == nil || response.Data.Status != transparents.PaymentStatusPaid {
		t.Fatalf("response = %+v", response)
	}
}

func TestTransparentHTTPErrorIsReturned(t *testing.T) {
	client, cleanup := newClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
	}))
	defer cleanup()

	_, err := client.Transparents.Get(context.Background(), "pix-1")
	var httpErr *fetch.HTTPError
	if !errors.As(err, &httpErr) || httpErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("error = %v", err)
	}
}
