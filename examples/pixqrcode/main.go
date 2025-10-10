package main

import (
	"context"
	"log"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/pixqrcode"
	"github.com/shopspring/decimal"
)

func main() {
	client, err := abacatepay.New(&abacatepay.ClientConfig{
		ApiKey:  "abc_dev_dC5zXgrEJEtnxEFeeu61xcus",
		Timeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	//create a new pix qr code
	body := &pixqrcode.CreatePixQrCode{
		Amount:      decimal.NewFromInt(100), // Amount in cents (e.g., 100 = R$1.00)
		ExpiresIn:   3600,                    // Expiration time in seconds (optional)
		Description: "Example Pix QR Code",   // Description (optional)
		Customer: pixqrcode.Customer{ // Customer information (optional)
			Name:  "John Doe",
			Email: "john.doe@example.com",
			TaxId: "000.000.000-00", // Document
			Phone: "+5511999999999",
		},
	}

	// Call the CreatePixQrCode method
	resp, err := client.PixQrCode.CreatePixQrCode(context.Background(), body)
	if err != nil {
		log.Fatalf("Failed to create Pix QR code: %v", err)
	}

	log.Printf("Created Pix QR code: %+v", resp)
}
