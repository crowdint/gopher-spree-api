package models

import "time"

type Product struct {
	ID                 int64
	Name               string
	Description        string
	AvailableOn        time.Time
	DeletedAt          time.Time
	Slug               string
	MetaDescription    string
	MetaKeyWords       string
	TaxCategoryID      int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Promotionable      bool
	MetaTitle          string
	ShippingCategoryId int64
}

func (this Product) TableName() string {
	return "spree_products"
}
