package domain

type ShippingCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (this ShippingCategory) TableName() string {
	return "spree_shipping_categories"
}
