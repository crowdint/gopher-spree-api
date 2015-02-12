package json

import (
	"time"

	"github.com/crowdint/gopher-spree-api/configs/spree"
)

type Variant struct {
	Id              int64          `json:"id"`
	Name            string         `json:"name"`
	Sku             string         `json:"sku"`
	Price           float64        `json:"price,string"`
	Weight          float64        `json:"weight,string"`
	Height          float64        `json:"height,string"`
	Width           float64        `json:"width,string"`
	Depth           float64        `json:"depth,string"`
	IsMaster        bool           `json:"is_master"`
	Slug            string         `json:"slug"`
	Description     string         `json:"description"`
	TrackInventory  bool           `json:"track_inventory"`
	CostPrice       string         `json:"cost_price"`
	DisplayPrice    string         `json:"display_price"`
	OptionsText     string         `json:"options_text"`
	InStock         bool           `json:"in_stock"`
	IsBackorderable bool           `json:"is_backorderable"`
	TotalOnHand     int64          `json:"total_on_hand"`
	IsDestroyed     bool           `json:"is_destroyed"`
	OptionValues    []*OptionValue `json:"option_values"`
	Images          []*Asset       `json:"images"`
	ProductId       int64          `json:"product_id"`
	StockItems      []*StockItem   `json:"-"`
	DeletedAt       time.Time      `json:"-"`
}

func (this *Variant) AfterFind() (err error) {
	this.IsDestroyed = !this.DeletedAt.IsZero()
	return
}

func (this *Variant) ShouldTrackInventory() bool {
	return this.TrackInventory && spree.IsInventoryTrackingEnabled()
}

func (this Variant) TableName() string {
	return "spree_variants"
}
