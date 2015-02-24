package json

type ShippingMethod struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	// Associations
	ShippingCategories []ShippingCategory `json:"shipping_categories"`
	Zones              []ShippingZone     `json:"zones" gorm:"many2many:spree_shipping_methods_zones;"`
}

func (this ShippingMethod) TableName() string {
	return "spree_shipping_methods"
}
