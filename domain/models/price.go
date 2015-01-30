package models

import "time"

type Price struct {
	Id        int64
	VariantId int64
	Amount    float64
	Currency  string
	DeletedAt time.Time
}

func (this Price) TableName() string {
	return "spree_prices"
}
