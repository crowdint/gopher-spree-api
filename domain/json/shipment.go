package json

import (
	"time"
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

	// Associations
	StockLocation   StockLocation    `json:"-"`
	Adjustments     []Adjustment     `json:"adjustments"`
	ShippingMethods []ShippingMethod `json:"shipping_methods"`
	ShippingRates   []ShippingRate   `json:"shipping_rates"`
	Manifest        []InventoryUnit  `json:"shipment_manifest"`
}

func (this Shipment) TableName() string {
	return "spree_shipments"
}
