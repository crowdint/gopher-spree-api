package models

import "time"

type ProductRepository interface {
	FindById(id int64)
}

type Product struct {
	Id                 int64
	Name               string
	Description        string
	AvailableOn        time.Time
	DeletedAt          time.Time
	Slug               string
	MetaDescription    string
	MetaKeyWords       string
	TaxCategoryId      int64
	ShippingCategoryId int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Promotionable      bool
	MetaTitle          string
}

func (this Product) TableName() string {
	return "spree_products"
}
