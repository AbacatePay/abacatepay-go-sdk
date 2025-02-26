package billing

type CouponStatus string

const (
	Active   CouponStatus = "ACTIVE"
	Deleted  CouponStatus = "DELETED"
	Disabled CouponStatus = "DISABLED"
)

// Cupom que compoe a billingresponse
type CouponUsed struct {
	CouponName string `json:"coupon"`
	Discount   int    `json:"discount"`
}

type Coupon struct {
	ID           string       `json:"id"`
	DiscountKind string       `json:"discountKind"`
	Discount     int          `json:"discount"`
	MaxRedeems   int          `json:"maxRedeems"`
	RedeemsCount int          `json:"redeemsCount"`
	Status       CouponStatus `json:"status"`
	DevMode      bool         `json:"devMode"`
	Notes        string       `json:"notes"`
	CreatedAt    string       `json:"createdAt"`
	UpdatedAt    string       `json:"updatedAt"`
}
