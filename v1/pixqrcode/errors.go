package pixqrcode

import "fmt"

type InvalidPixQrCodeIdError struct {
	ID string
}

func (e *InvalidPixQrCodeIdError) Error() string {
	return fmt.Sprintf("invalid Pix QR Code ID: %q", e.ID)
}
