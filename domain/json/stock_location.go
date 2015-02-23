package json

import "time"

type StockLocation struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
	Active  bool   `json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (this StockLocation) TableName() string {
	return "spree_stock_locations"
}
