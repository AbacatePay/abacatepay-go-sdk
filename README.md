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

## Documentation

[https://abacatepay.readme.io](https://abacatepay.readme.io)

## License

MIT
