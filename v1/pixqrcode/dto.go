package pixqrcode

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

var validate *validator.Validate

// CreatePixQrCode is used to create a Pix QR Code
// It requires the amount to be charged and optionally accepts an expiration time, description, customer information, and metadata
type CreatePixQrCode struct {
	Amount      decimal.Decimal `json:"amount" validate:"required,decimal_gt_zero"` // Amount to be charged in cents (must be greater than 0)
	ExpiresIn   int             `json:"expiresIn"`                                  // Expiration time in seconds
	Description string          `json:"description"`                                // Description of the Pix QR Code
	Customer    Customer        `json:"customer" validate:"required"`               // Customer information
	Metadata    Metadata        `json:"metadata"`                                   // Metadata to be attached to the Pix QR Code object
}

type Customer struct {
	Name  string `json:"name" validate:"required"`        // Customer's name
	Email string `json:"email" validate:"required,email"` // Customer's email
	Phone string `json:"cellphone" validate:"required"`   // Customer's phone number
	TaxId string `json:"taxId" validate:"required"`       // Customer's tax identification number (CPF/CNPJ)
}

// Metadata is a map of key-value pairs that can be attached to a Pix QR Code object
type Metadata struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

// SimulatePayment is used to simulate a payment on development environment for a Pix QR Code
// It requires the ID of the Pix QR Code to be paid
type SimulatePayment struct {
	Id string `json:"id" validate:"required"` // ID of the Pix QR Code to be paid
}

// CheckStatus is used to check the status of a Pix QR Code
// It requires the ID of the Pix QR Code to check
type CheckStatus struct {
	Id string `json:"id" validate:"required"` // ID of the Pix QR Code to check status
}

// CreatePixQrCodeResponse represents the response from creating a Pix QR Code
type CreatePixQrCodeResponse struct {
	Data  CreatePixQrCodeResponseData `json:"data"`
	Error *error                      `json:"error"`
}

// CreatePixQrCodeResponseData contains the details of the created Pix QR Code
type CreatePixQrCodeResponseData struct {
	Id           string          `json:"id"`
	Amount       decimal.Decimal `json:"amount"`
	Status       string          `json:"status"`
	DevMode      bool            `json:"devMode"`
	BrCode       string          `json:"brCode"`
	BrCodeBase64 string          `json:"brCodeBase64"`
	PlatformFee  int64           `json:"platformFee"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedAt    string          `json:"updatedAt"`
	ExpiresAt    string          `json:"expiresAt"`
}

// SimulatePaymentResponse represents the response from simulating a payment on a Pix QR Code
type SimulatePaymentResponse struct {
	Data  SimulatePaymentResponseData `json:"data"`
	Error *error                      `json:"error"`
}

// CheckStatusResponse represents the response from checking the status of a Pix QR Code
type SimulatePaymentResponseData struct {
	Id           string          `json:"id"`
	Amount       decimal.Decimal `json:"amount"`
	Status       string          `json:"status"`
	DevMode      bool            `json:"devMode"`
	BrCode       string          `json:"brCode"`
	BrCodeBase64 string          `json:"brCodeBase64"`
	PlatformFee  int64           `json:"platformFee"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedAt    string          `json:"updatedAt"`
	ExpiresAt    string          `json:"expiresAt"`
}

// CheckStatusResponse represents the response from checking the status of a Pix QR Code
type CheckStatusResponse struct {
	Data  CheckStatusResponseData `json:"data"`
	Error *error                  `json:"error"`
}

// CheckStatusResponseData contains the status and expiration time of the Pix QR Code
type CheckStatusResponseData struct {
	Status    string `json:"status"`
	ExpiresAt string `json:"expiresAt"`
}

func init() {
	validate = validator.New()

	validate.RegisterValidation("decimal_gt_zero", validateDecimalGtZero)
}

func validateDecimalGtZero(fl validator.FieldLevel) bool {
	field := fl.Field()
	if !field.IsValid() {
		return false
	}

	if dec, ok := field.Interface().(decimal.Decimal); ok {
		return dec.GreaterThan(decimal.Zero)
	}

	return false
}

func (p *CreatePixQrCode) Validate() error {
	return validate.Struct(p)
}
