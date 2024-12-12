package customer

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type CreateCustomerBody struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
	TaxID     string `json:"taxId"`
	Email     string `json:"email" validate:"required,email"`
}

type CreateCustomerResponse struct {
	Customer CreateCustomerResponseItem `json:"data"`
}

type CreateCustomerResponseItem struct {
	AccountID string           `json:"accountId"`
	StoreID   string           `json:"storeId"`
	DevMode   bool           `json:"devMode"`
	Metadata  CustomerMetadata `json:"metadata"`
	CreatedAt string           `json:"createdAt"`
	UpdatedAt string           `json:"updatedAt"`
	ID        string           `json:"id"`
}

type ListCustomerResponse struct {
	Customers []CustomerResponseItem `json:"data"`
}

type CustomerResponseItem struct {
	Metadata CustomerMetadata `json:"metadata"`
	ID       string           `json:"id"`
}
type Customer struct {
	Metadata  CustomerMetadata `json:"metadata"`
	ID        string           `json:"_id"`
	PublicID  string           `json:"publicId"`
	AccountID string           `json:"accountId"`
	StoreID   string           `json:"storeId"`
	DevMode   bool             `json:"devMode"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Version   int              `json:"__v"`
}

type CustomerMetadata struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
	TaxID     string `json:"taxId"`
	Email     string `json:"email"`
}

func init() {
	validate = validator.New()
}

func (p *CreateCustomerBody) Validate() error {
	return validate.Struct(p)
}
