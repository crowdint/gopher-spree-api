package domain

import (
	"time"
)

type Variant struct {
	ID              int64     `json:"id"`
	Sku             string    `json:"sku"`
	Weight          float64   `json:"weight"`
	Height          float64   `json:"height"`
	Width           float64   `json:"width"`
	Depth           float64   `json:"depth"`
	IsMaster        bool      `json:"is_master"`
	ProductId       int64     `json:"product_id"`
	CostPrice       float64   `json:"cost_price"`
	CostCurrency    string    `json:"cost_currency"`
	TrackInventory  bool      `json:"track_inventory"`
	TaxCategoryId   int64     `json:"tax_category_id"`
	StockItemsCount int64     `json:"stock_items_count"`
	DeletedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}
