package domain

import (
	"time"
)

type Product struct {
	ID                 int64             `json:"id" sql:"id"`
	Name               string            `json:"name" sql:"name"`
	Description        string            `json:"description" sql:"description"`
	Price              string            `json:"price" sql:"-"`
	DisplayPrice       string            `json:"display_price" sql:"-"`
	AvailableOn        time.Time         `json:"available_on" sql:"available_on"`
	Slug               string            `json:"slug" sql:"slug"`
	MetaDescription    string            `json:"meta_description" sql:"meta_description"`
	MetaKeyWords       string            `json:"meta_keywords" sql:"meta_keywords"`
	ShippingCategoryId int64             `json:"shipping_category_id" sql:"shipping_category_id"`
	TaxonIds           []int             `json:"taxon_ids" sql:"-"`
	TotalOnHand        int64             `json:"total_on_hand" sql:"-"`
	HasVariants        bool              `json:"has_variants" sql:"-"`
	Master             Variant           `json:"master" sql:"-"`
	Variants           []Variant         `json:"variants" sql:"-"`
	OptionTypes        []OptionType      `json:"option_types" sql:"-"`
	ProductProperties  []ProductProperty `json:"product_properties" sql:"-"`
	Classifications    []Classification  `json:"classifications" sql:"-"`
}
