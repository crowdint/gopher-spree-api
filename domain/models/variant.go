package models

import "time"

type VariantRepository interface {
	FindByProductId(id int64)
}

type Variant struct {
	Id                  int64
	Sku                 string
	Price               float64
	Weight              float64
	Height              float64
	Width               float64
	Depth               float64
	DeletedAt           time.Time
	IsMaster            bool
	ProductId           int64
	CostPrice           string
	Position            int64
	CostCurrency        string
	TrackInventory      bool
	TaxCategoryId       int64
	UpdatedAt           time.Time
	StockItemsCount     int64
	RealStockItemsCount int64
	Backorderable       bool
}

func (this Variant) TableName() string {
	return "spree_variants"
}
