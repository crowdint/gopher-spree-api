package domain

import (
	"time"

	"github.com/crowdint/gopher-spree-api/configs/spree"
)

type Shipment struct {
	Id                   int64        `json:"id"`
	Cost                 string       `json:"cost"`
	Number               string       `json:"number"`
	OrderId              string       `json:"order_id"`
	SelectedShippingRate ShippingRate `json:"selected_shipping_rate"`
	State                string       `json:"state"`
	StockLocationName    string       `json:"stock_location_name"`
	StockLocationId      int64        `json:"-"`
	Tracking             *string      `json:"tracking"`
	ShippedAt            *time.Time   `json:"shipped_at"`

	StockLocation   StockLocation    `json:"-"`
	Adjustments     []Adjustment     `json:"adjustments"`
	ShippingMethods []ShippingMethod `json:"shipping_methods"`
	ShippingRates   []ShippingRate   `json:"shipping_rates"`
	Manifest        []InventoryUnit  `json:"shipment_manifest"`

	Order Adjustable `json:"-" sql:"-"`
}

func (this Shipment) AdjustableId() int64 {
	return this.Id
}

func (this Shipment) AdjustableCurrency() string {
	if this.Order != nil {
		return this.Order.AdjustableCurrency()
	}

	return spree.Get(spree.CURRENCY)
}

func (this Shipment) TableName() string {
	return "spree_shipments"
}

func (this Shipment) SpreeClass() string {
	return "Spree::Shipment"
}
