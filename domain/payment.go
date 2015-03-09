package domain

import (
	"time"

	. "github.com/crowdint/gopher-spree-api/utils"
)

type Payment struct {
	Id              int64         `json:"id"`
	Amount          float64       `json:"amount,string"`
	AVSResponse     string        `json:"avs_response"`
	DisplayAmount   string        `json:"display_amount"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	PaymentMethodId int64         `json:"payment_method_id"`
	ResponseCode    string        `json:"response_code"`
	Source          PaymentSource `json:"source"`
	SourceId        int64         `json:"source_id"`
	SourceType      string        `json:"source_type"`
	State           string        `json:"state"`
	OrderId         int64         `json:"-"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`

	Order *Order `json:"-" sql:"-"`
}

func (this *Payment) AfterFind() (err error) {
	this.DisplayAmount = Monetize(this.Amount, this.Currency())

	return
}

func (this *Payment) Currency() string {
	return this.Order.Currency
}

func (this Payment) TableName() string {
	return "spree_payments"
}
