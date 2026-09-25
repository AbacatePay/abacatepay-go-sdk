package transparents

import (
	"encoding/json"
	"fmt"
	"time"

	v2 "github.com/AbacatePay/go-types/v2"
)

// CreateQRCodePIXBody is the request envelope expected by the transparent
// checkout endpoint.
type CreateQRCodePIXBody struct {
	Method string                  `json:"method"`
	Data   CreateQRCodePIXDataBody `json:"data"`
}

type CreateQRCodePIXDataBody struct {
	Amount      uint64              `json:"amount"`
	ExpiresIn   *uint64             `json:"expiresIn,omitempty"`
	Description *string             `json:"description,omitempty"`
	Customer    *v2.APIBaseCustomer `json:"customer,omitempty"`
	ExternalID  *string             `json:"externalId,omitempty"`
	Metadata    map[string]any      `json:"metadata,omitempty"`
}

type PaymentStatus string

const (
	PaymentStatusPending      PaymentStatus = "PENDING"
	PaymentStatusExpired      PaymentStatus = "EXPIRED"
	PaymentStatusCancelled    PaymentStatus = "CANCELLED"
	PaymentStatusPaid         PaymentStatus = "PAID"
	PaymentStatusUnderDispute PaymentStatus = "UNDER_DISPUTE"
	PaymentStatusRefunded     PaymentStatus = "REFUNDED"
	PaymentStatusRedeemed     PaymentStatus = "REDEEMED"
	PaymentStatusApproved     PaymentStatus = "APPROVED"
	PaymentStatusFailed       PaymentStatus = "FAILED"
	PaymentStatusLost         PaymentStatus = "LOST"
)

type QRCodePIX struct {
	ID           string         `json:"id"`
	Amount       uint64         `json:"amount"`
	Status       PaymentStatus  `json:"status"`
	DevMode      bool           `json:"devMode"`
	BrCode       string         `json:"brCode"`
	BrCodeBase64 string         `json:"brCodeBase64"`
	PlatformFee  uint64         `json:"platformFee"`
	ReceiptURL   *string        `json:"receiptUrl,omitempty"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	ExpiresAt    time.Time      `json:"expiresAt"`
	ExternalID   string         `json:"externalId,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

type APIResponse[T any] struct {
	Data    *T      `json:"data,omitempty"`
	Error   *string `json:"error,omitempty"`
	Success Success `json:"success"`
}

type Success bool

func (s *Success) UnmarshalJSON(data []byte) error {
	var flag bool
	if err := json.Unmarshal(data, &flag); err == nil {
		*s = Success(flag)
		return nil
	}

	var object map[string]any
	if err := json.Unmarshal(data, &object); err == nil && object != nil {
		*s = true
		return nil
	}

	return fmt.Errorf("transparents: invalid success value")
}

func (s Success) Bool() bool {
	return bool(s)
}

type CreateQRCodePIXData = APIResponse[QRCodePIX]

type ListQRCodePIXQuery struct {
	ID         string
	ExternalID string
	Status     PaymentStatus
	Method     string
	After      string
	Before     string
	StartDate  string
	EndDate    string
	Limit      int
}

type ListQRCodePIXData struct {
	Data       []QRCodePIX          `json:"data"`
	Error      *string              `json:"error,omitempty"`
	Success    Success              `json:"success"`
	Pagination CursorPaginationData `json:"pagination"`
}

type CursorPaginationData struct {
	HasMore bool    `json:"hasMore"`
	Next    *string `json:"next"`
	Before  *string `json:"before"`
}

type CheckStatusData = APIResponse[CheckStatus]

type CheckStatus struct {
	ID        string        `json:"id"`
	Status    PaymentStatus `json:"status"`
	ExpiresAt time.Time     `json:"expiresAt"`
}

// SimulatePaymentData remains compatible with the existing v2 response type.
type SimulatePaymentData = v2.RESTPostSimulateQRCodePixPaymentData
