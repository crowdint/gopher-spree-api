package domain

import (
	"time"
)

type Payment struct {
	Id              int64         `json:"id"`
	Amount          string        `json:"amount"`
	AVSResponse     string        `json:"avs_response"`
	DisplayAmount   string        `json:"display_amount"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	PaymentMethodId int64         `json:"payment_method_id"`
	ResponseCode    string        `json:"response_code"`
	Source          PaymentSource `json:"source"`
	SourceId        int64         `json:"source_id"`
	SourceType      string        `json:"source_type"`
	State           string        `json:"state"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

func (this Payment) TableName() string {
	return "spree_payments"
}
