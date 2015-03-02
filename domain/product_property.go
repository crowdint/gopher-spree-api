package domain

import "time"

type ProductProperty struct {
	Id           int64  `json:"id"`
	ProductId    int64  `json:"product_id"`
	PropertyId   int64  `json:"property_id"`
	Value        string `json:"value"`
	PropertyName string `json:"property_name" sql:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Position  int64     `json:"-"`
}

func (this ProductProperty) TableName() string {
	return "spree_product_properties"
}
