package domain

import "time"

type Price struct {
  Id        int64 `json:"-"`
	VariantId int64 `json:"-"`
	Amount    float64 `json:"-"`
	Currency  string `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (this Price) TableName() string {
	return "spree_prices"
}
