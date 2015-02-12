package json

import (
	"time"
)

type StockItem struct {
	Id              int64
	StockLocationId int64
	VariantId       int64
	CountOnHand     int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Backorderable   bool
}

func (this StockItem) TableName() string {
	return "spree_stock_items"
}
