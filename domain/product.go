package domain

import "time"

type Product struct {
	Id                 int64             `json:"id"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	Price              string            `json:"price"`
	DisplayPrice       string            `json:"display_price"`
	AvailableOn        time.Time         `json:"available_on"`
	Slug               string            `json:"slug"`
	MetaDescription    string            `json:"meta_description"`
	MetaKeyWords       string            `json:"meta_keywords"`
	ShippingCategoryId int64             `json:"shipping_category_id"`
	TaxonIds           []int             `json:"taxon_ids"`
	TotalOnHand        int64             `json:"total_on_hand"`
	HasVariants        bool              `json:"has_variants"`
	Master             Variant           `json:"master"`
	Variants           []Variant         `json:"variants"`
	OptionTypes        []OptionType      `json:"option_types"`
	ProductProperties  []ProductProperty `json:"product_properties"`
	Classifications    []Classification  `json:"classifications"`

	DeletedAt     time.Time `json:"-"`
	TaxCategoryId int64     `json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	Promotionable bool      `json:"-"`
	MetaTitle     string    `json:"-"`
}

func (this Product) TableName() string {
	return "spree_products"
}
