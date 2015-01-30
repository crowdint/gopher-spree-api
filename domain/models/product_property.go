package models

import "time"

type ProductProperty struct {
	Id         int64
	Value      string
	ProductId  int64
	PropertyId int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Position   int64
}

func (this ProductProperty) TableName() string {
	return "spree_product_properties"
}
