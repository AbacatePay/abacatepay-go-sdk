# abacatepay-go-sdk

## Installation

```bash
go get github.com/antunesgabriel/abacatepay-go-sdk
```

## Usage

```go
package main

import (
	"fmt"
    "context"

    "https://github.com/antunesgabriel/abacatepay-go-sdk"
	"https://github.com/antunesgabriel/abacatepay-go-sdk/v1/client"
)

func main() {
	c := client.New(&client.ClientConfig{
		ApiKey: "your-api-key",
	})

    // List all billings
    ctx := context.Background()

    listResponse, err := c.Billing.ListAll(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println(listResponse)

    // Create a new Billing
    body := &billing.CreateBillingBody{
        Frequency:     abacatepay.OneTime,
        Methods:       []abacatepay.Method{abacatepay.PIX},
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
    }

    createResponse, err := c.Billing.Create(ctx, body)
    if err != nil {
        panic(err)
    }

    fmt.Println(createResponse)
}
```

## Documentation

[https://abacatepay.readme.io](https://abacatepay.readme.io)

## License

MIT
