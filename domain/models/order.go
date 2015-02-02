package models

import (
	"time"
)

type Order struct {
	Id int64

	AdditionalTaxTotal    float64
	AdjustmentTotal       float64
	ApprovedAt            time.Time
	Approver              int64
	BillAddress           int64
	CanceledAt            time.Time
	Canceler              int64
	Channel               string
	CompletedAt           time.Time
	ConfirmationDelivered bool
	ConsideredRisky       bool
	CreatedAt             time.Time
	CreatedBy             int64
	Currency              string
	Email                 string
	GuestToken            string
	IncludedTaxTotal      float64
	ItemCount             int64
	ItemTotal             float64
	LastIpAddress         string
	Number                string
	PaymentState          string
	PaymentTotal          float64
	PromoTotal            float64
	ShipAddress           int64
	ShipmentState         string
	ShipmentTotal         float64
	ShippingMethod        int64
	SpecialInstructions   string
	State                 string
	StateLockVersion      int64
	Store                 int64
	Total                 float64
	UpdatedAt             time.Time
	UserId                int64
}

func (o Order) TableName() string {
  return "spree_orders"
}
