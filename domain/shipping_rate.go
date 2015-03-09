package domain

import (
	. "github.com/crowdint/gopher-spree-api/utils"
)

type ShippingRate struct {
	Id                 int64   `json:"id"`
	Cost               float64 `json:"cost,string"`
	DisplayCost        string  `json:"display_cost"`
	Name               string  `json:"name"`
	Selected           bool    `json:"selected"`
	ShippingMethodCode *string `json:"shipping_method_code"`
	ShippingMethodId   int64   `json:"shipping_method_id"`

	ShippingMethod ShippingMethod `json:"-"`
	Shipment       *Shipment      `json:"-" sql:"-"`
}

func (this *ShippingRate) SetComputedValues() {
	this.DisplayCost = Monetize(this.Cost, this.Currency())
}

func (this *ShippingRate) Currency() string {
	return this.Shipment.Currency()
}

func (this ShippingRate) TableName() string {
	return "spree_shipping_rates"
}
