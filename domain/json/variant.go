package json

import (
	"time"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type Variant struct {
	Id          int64     `json:"id"`
	CostPrice   string    `json:"cost_price"`
	Depth       float64   `json:"depth,string"`
	Height      float64   `json:"height,string"`
	IsMaster    bool      `json:"is_master"`
	OptionsText string    `json:"options_text"`
	Price       float64   `json:"price,string"`
	ProductId   int64     `json:"product_id"`
	Sku         string    `json:"sku"`
	Weight      float64   `json:"weight,string"`
	Width       float64   `json:"width,string"`
	DeletedAt   time.Time `json:"-"`

	// Computed
	Description     string `json:"description"`
	DisplayPrice    string `json:"display_price"`
	InStock         bool   `json:"in_stock"`
	IsBackorderable bool   `json:"is_backorderable"`
	IsDestroyed     bool   `json:"is_destroyed"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	TotalOnHand     *int64 `json:"total_on_hand"`
	TrackInventory  bool   `json:"track_inventory"`

	// Associations
	Images       []*Asset             `json:"images"`
	OptionValues []models.OptionValue `json:"option_values" gorm:"many2many:spree_option_values_variants;"`
	StockItems   []*StockItem         `json:"-"`
}

func (this *Variant) AfterFind() (err error) {
	this.IsDestroyed = !this.DeletedAt.IsZero()
	return
}

func (this *Variant) SetInventoryValues() {
	if this.ShouldTrackInventory() {
		for _, stockItem := range this.StockItems {
			var totalOnHand int64

			if this.TotalOnHand != nil {
				totalOnHand = (*this.TotalOnHand + stockItem.CountOnHand)
			} else {
				totalOnHand = stockItem.CountOnHand
			}

			this.TotalOnHand = &totalOnHand

			if stockItem.Backorderable {
				this.IsBackorderable = true
			}
		}
		this.InStock = *this.TotalOnHand > 0
	} else {
		this.IsBackorderable = true
		this.InStock = true
	}
}

func (this *Variant) ShouldTrackInventory() bool {
	return this.TrackInventory && spree.IsInventoryTrackingEnabled()
}

func (this Variant) TableName() string {
	return "spree_variants"
}
