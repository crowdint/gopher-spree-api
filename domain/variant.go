package domain

import (
	"strconv"
	"time"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	. "github.com/crowdint/gopher-spree-api/utils"
)

var (
	variantErrors *ValidatorErrors
)

type Variant struct {
	Id           int64     `json:"id"`
	CostPrice    *string   `json:"cost_price"`
	Depth        float64   `json:"depth,string"`
	Height       float64   `json:"height,string"`
	IsMaster     bool      `json:"is_master"`
	OptionsText  string    `json:"options_text" sql:"-"`
	Price        *float64  `json:"price,string" sql:"-"`
	DefaultPrice Price     `json:"-"`
	ProductId    int64     `json:"product_id"`
	Product      *Product  `json:"-" sql:"-"`
	Sku          string    `json:"sku" sql:",unique"`
	Weight       float64   `json:"weight,string"`
	Width        float64   `json:"width,string"`
	DeletedAt    time.Time `json:"-"`

	Description     string `json:"description" sql:"-"`
	DisplayPrice    string `json:"display_price" sql:"-"`
	InStock         bool   `json:"in_stock" sql:"-"`
	IsBackorderable bool   `json:"is_backorderable" sql:"-"`
	IsDestroyed     bool   `json:"is_destroyed" sql:"-"`
	Name            string `json:"name" sql:"-"`
	Slug            string `json:"slug" sql:"-"`
	TotalOnHand     *int64 `json:"total_on_hand" sql:"-"`
	TrackInventory  bool   `json:"track_inventory" sql:"-"`

	Images       []*Asset      `json:"images" sql:"-"`
	OptionValues []OptionValue `json:"option_values" gorm:"many2many:spree_option_values_variants;"`
	StockItems   []*StockItem  `json:"-" sql:"-"`

	Position            *int64    `json:"-"`
	CostCurrency        string    `json:"-"`
	TaxCategoryId       int64     `json:"-"`
	UpdatedAt           time.Time `json:"-"`
	StockItemsCount     int64     `json:"-"`
	RealStockItemsCount int64     `json:"-" sql:"-"`
	Backorderable       bool      `json:"-" sql:"-"`
}

func (this *Variant) AfterFind() (err error) {
	this.IsDestroyed = !this.DeletedAt.IsZero()

	return
}

func (this *Variant) SetComputedValues() {
	this.setInventoryValues()

	this.DisplayPrice = Monetize(this.Price, this.CostCurrency)
}

func (this *Variant) setInventoryValues() {
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
