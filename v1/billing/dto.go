package billing

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type CreateBillingBody struct {
	Frequency     Frequency         `json:"frequency"     validate:"required"`
	Methods       []Method          `json:"methods"       validate:"required,dive"`
	ReturnUrl     string            `json:"returnUrl"     validate:"required,url"`
	CompletionUrl string            `json:"completionUrl" validate:"required,url"`
	Products      []*BillingProduct `json:"products"      validate:"required,dive"`
	CustomerId    string            `json:"customerId"`
	Customer      *BillingCustomer  `json:"customer"`
}

type BillingCustomer struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
	Email     string `json:"email" validate:"required"`
	TaxID     string `json:"taxId"`
}

type BillingProduct struct {
	ExternalId  string `json:"externalId"  validate:"required"`
	Name        string `json:"name"        validate:"required"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"    validate:"required,gte=1"`
	Price       int    `json:"price"       validate:"required,gte=100"`
}

type ProductItem struct {
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
	Quantity   int    `json:"quantity"`
}

type CreateBillingResponseItem struct {
	PublicID  string        `json:"publicId"`
	Products  []ProductItem `json:"products"`
	Amount    int64         `json:"amount"`
	Status    string        `json:"status"`
	DevMode   bool          `json:"devMode"`
	Methods   []string      `json:"methods"`
	Frequency string        `json:"frequency"`
	Metadata  struct {
		Fee           int64  `json:"fee"`
		ReturnURL     string `json:"returnUrl"`
		CompletionURL string `json:"completionUrl"`
	} `json:"metadata"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ID        string `json:"_id"`
	Version   int    `json:"__v"`
	URL       string `json:"url"`
	BillingID string `json:"id"`
}

type CreateBillingResponse struct {
	Data  CreateBillingResponseItem `json:"data"`
	Error string                    `json:"error"`
}

type Metadata struct {
	Fee           int    `json:"fee"`
	ReturnURL     string `json:"returnUrl"`
	CompletionURL string `json:"completionUrl"`
}

type CustomerMetadata struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
	TaxID     string `json:"taxId"`
	Email     string `json:"email"`
}

// Customer representa o cliente associado a uma cobrança.
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

type BillingListItem struct {
	ID       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
	Customer struct {
		ID       string           `json:"id"`
		Metadata CustomerMetadata `json:"metadata"`
	} `json:"customer"`
	CustomerId struct {
		Metadata  CustomerMetadata `json:"metadata"`
		ID        string           `json:"id"`
		PublicID  string           `json:"publicId"`
		AccountID string           `json:"accountId"`
		StoreID   string           `json:"storeId"`
		DevMode   bool             `json:"devMode"`
		CreatedAt time.Time        `json:"createdAt"`
		UpdatedAt time.Time        `json:"updatedAt"`
		Version   int              `json:"__v"`
	} `json:"customerId"`
	PublicID  string        `json:"publicId"`
	Amount    int64         `json:"amount"`
	Status    string        `json:"status"`
	DevMode   bool          `json:"devMode"`
	Methods   []Method      `json:"methods"`
	Frequency Frequency     `json:"frequency"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Version   int           `json:"__v"`
	URL       string        `json:"url"`
	Products  []ProductItem `json:"products"`
}

type ListBillingResponse struct {
	Data  []BillingListItem `json:"data"`
	Error string            `json:"error"`
}

func init() {
	validate = validator.New()
}

func (p *CreateBillingBody) Validate() error {
	return validate.Struct(p)
}
