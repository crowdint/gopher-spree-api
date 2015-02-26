package domain

import (
	"time"
)

type StockItem struct {
	Id              int64
	Backorderable   bool
	CountOnHand     int64
	StockLocationId int64
	VariantId       int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (this StockItem) TableName() string {
	return "spree_stock_items"
}
