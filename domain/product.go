package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	productErrors = &ValidatorErrors{}
)

type Product struct {
	Id                 int64             `json:"id"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	Price              string            `json:"price" sql:"-"`
	DisplayPrice       string            `json:"display_price" sql:"-"`
	AvailableOn        time.Time         `json:"available_on"`
	Slug               string            `json:"slug"`
	MetaDescription    string            `json:"meta_description"`
	MetaKeyWords       string            `json:"meta_keywords" sql:"-"`
	ShippingCategoryId int64             `json:"shipping_category_id"`
	TaxonIds           []int             `json:"taxon_ids" sql:"-"`
	TotalOnHand        int64             `json:"total_on_hand" sql:"-"`
	HasVariants        bool              `json:"has_variants" sql:"-"`
	Master             Variant           `json:"master"`
	Variants           []Variant         `json:"variants" sql:"-"`
	OptionTypes        []OptionType      `json:"option_types" sql:"-"`
	ProductProperties  []ProductProperty `json:"product_properties" sql:"-"`
	Classifications    []Classification  `json:"classifications" sql:"-"`

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

func (this Product) SpreeClass() string {
	return "Spree::Product"
}

func (this *Product) Key() string {
	return fmt.Sprintf("%s/%d/%d", this.SpreeClass(), this.Id, this.UpdatedAt.Unix())
}

func (this *Product) KeyWithPrefix(prefix string) string {
	return fmt.Sprintf("%s/%s/%d/%d", this.SpreeClass(), prefix, this.Id, this.UpdatedAt.Unix())
}

func (this *Product) Marshal() ([]byte, error) {
	return json.Marshal(this)
}

func (this *Product) Unmarshal(data []byte) error {
	return json.Unmarshal(data, this)
}

func (this *Product) IsValid() bool {
	productErrors = &ValidatorErrors{}

	if this.Name == "" {
		productErrors.Add("name", ErrNotBlank.Error())
	}

	if this.Price == "" {
		productErrors.Add("price", ErrNotBlank.Error())
	}

	if this.ShippingCategoryId == 0 {
		productErrors.Add("shipping_category_id", ErrNotBlank.Error())
	}

	if len(this.Slug) < 3 {
		productErrors.Add("slug", ErrTooShort(3).Error())
	}

	return productErrors.IsEmpty()
}

func (this *Product) SlugCandidates()[]interface{} {
	return []interface{}{
		this.Name,
		[]interface{}{this.Name, this.Master.Sku},
	}
}

func (this *Product) SetSlug(slug string) {
	this.Slug = slug
}

func (this *Product) GetErrors() *ValidatorErrors {
	if productErrors.IsEmpty() {
		return nil
	}

	return productErrors
}
