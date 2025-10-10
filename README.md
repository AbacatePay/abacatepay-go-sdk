# abacatepay-go-sdk

## Installation

```bash
go get github.com/AbacatePay/abacatepay-go-sdk
```

## Usage

```go
package main

import (
	"context"
	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
	"log"
	"time"
)

func main() {
	client, err := abacatepay.New(&abacatepay.ClientConfig{
		ApiKey:  "abc_dev",
		Timeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	//create a new billing
	body := &billing.CreateBillingBody{
		Frequency:     billing.OneTime,
		Methods:       []billing.Method{billing.PIX},
		CompletionUrl: "https://example.com/completion",
		ReturnUrl:     "https://example.com/return",
		Products: []*billing.BillingProduct{
			{
				ExternalId:  "pix-1234",
				Name:        "Example Product",
				Description: "Example product description",
				Quantity:    1,
				Price:       100,
			},
		},
		Customer: &billing.BillingCustomer{
			Email: "test@example.com",
		},
	}

	ctx := context.Background()
	createResponse, err := client.Billing.Create(ctx, body)
	if err != nil {
		panic(err)
	}

	log.Println(createResponse)

	// list all billings
	billings, err := client.Billing.ListAll(ctx)
	if err != nil {
		panic(err)
	}

	log.Println(billings.Data)
}
```

### PIX QR Code

```go
package main

import (
	"context"
	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/pixqrcode"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

func main() {
	client, err := abacatepay.New(&abacatepay.ClientConfig{
		ApiKey:  "abc_dev",
		Timeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	// create a new PIX QR Code
	body := &pixqrcode.CreatePixQrCode{
		Amount:      decimal.NewFromInt(100), // Amount in cents (e.g., 100 = R$1.00)
		ExpiresIn:   3600,                    // Expiration time in seconds
		Description: "Example Pix QR Code",
		Customer: pixqrcode.Customer{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			TaxId: "502.842.288-50",
			Phone: "+5511999999999",
		},
	}

	ctx := context.Background()
	createResponse, err := client.PixQrCode.CreatePixQrCode(ctx, body)
	if err != nil {
		panic(err)
	}

	log.Println(createResponse)

	// check PIX QR Code status
	statusResponse, err := client.PixQrCode.CheckStatus(ctx, &pixqrcode.CheckStatus{
		Id: createResponse.Data.Id,
	})
	if err != nil {
		panic(err)
	}

	log.Println(statusResponse)
}
```

## Documentation

[https://docs.abacatepay.com/](https://docs.abacatepay.com/)

## License

MIT
