t-billing:
	go test -v ./v1/billing

t-fetch:
	go test -v ./internal/pkg/fetch

t-client:
	go test -v ./abacatepay

.PHONY: t-billing t-fetch t-client