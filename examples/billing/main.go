package main

import (
	"context"
	"log"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
)

func main() {
	client, err := abacatepay.New(&abacatepay.ClientConfig{
		APIKey:  "abc_dev",
		Timeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	//create a new billing
	body := &billing.CreateBillingBody{
		Frequency:     billing.OneTime,
		Methods:       []billing.Method{billing.PIX},
		CompletionURL: "https://example.com/completion",
		ReturnURL:     "https://example.com/return",
		Products: []*billing.BillingProduct{
			{
				ExternalID:  "pix-1234",
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
