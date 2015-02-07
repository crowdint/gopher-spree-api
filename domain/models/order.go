package models

import (
	"time"
)

type Order struct {
	Id        int64
	Number    string
	ItemTotal float64
	Total     float64

	State               string
	AdjustmentTotal     float64
	UserId              *int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CompletedAt         time.Time
	PaymentTotal        float64
	ShipmentState       string
	Email               string
	SpecialInstructions string
	Channel             string
	IncludedTaxTotal    float64
	AdditionalTaxTotal  float64

	Currency   string
	GuestToken string

	ApprovedAt            time.Time
	ApproverId            int64
	BillAddressId         int64
	CanceledAt            time.Time
	CancelerId            int64
	ConfirmationDelivered bool
	ConsideredRisky       bool
	CreatedBy             int64
	ItemCount             int64
	LastIpAddress         string
	PaymentState          string
	PromoTotal            float64
	ShipAddressId         int64
	ShipmentTotal         float64
	ShippingMethodId      int64
	StateLockVersion      int64
	StoreId               int64

	Quantity int64
}

func (this Order) TableName() string {
	return "spree_orders"
}

func (this Order) TaxTotal() float64 {
	return this.IncludedTaxTotal + this.AdditionalTaxTotal
}
