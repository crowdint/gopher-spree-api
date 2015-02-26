package domain

type ShippingRate struct {
	Id                 int64   `json:"id"`
	Cost               string  `json:"cost"`
	DisplayCost        string  `json:"display_cost"`
	Name               string  `json:"name"`
	Selected           bool    `json:"selected"`
	ShippingMethodCode *string `json:"shipping_method_code"`
	ShippingMethodId   int64   `json:"shipping_method_id"`

	ShippingMethod ShippingMethod `json:"-"`
}

func (this ShippingRate) TableName() string {
	return "spree_shipping_rates"
}
