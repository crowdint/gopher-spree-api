package domain

import (
	"time"
)

type Product struct {
	Name               string    `json:"name"`
	ID                 int64     `json:"id"`
	Description        string    `json:"description"`
	Slug               string    `json:"slug"`
	MetaDescription    string    `json:"meta_description"`
	MetaKeyWords       string    `json:"meta_keywords"`
	AvailableOn        time.Time `json:"available_on"`
	ShippingCategoryId int64     `json:"shipping_category_id"`
	MetaTitle          string    `json:"-"`
	Promotionable      bool      `json:"-"`
	TaxCathegoryID     int64     `json:"-"`
	DeletedAt          time.Time `json:"-"`
	UpdatedAt          time.Time `json:"-"`
}
