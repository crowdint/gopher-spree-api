package json

type ShippingMethod struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	// Associations
	ShippingCategories []ShippingCategory `json:"shipping_categories"`
	Zones              []ShippingZone     `json:"zones"`
}
