package domain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	. "github.com/crowdint/gopher-spree-api/utils"
)

var (
	productErrors *ValidatorErrors
)

type Product struct {
	Id                 int64             `json:"id"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	Price              *float64          `json:"price,string" sql:"-"`
	AvailableOn        time.Time         `json:"available_on"`
	Slug               string            `json:"slug"`
	MetaDescription    string            `json:"meta_description"`
	MetaKeyWords       string            `json:"meta_keywords" sql:"-"`
	ShippingCategoryId int64             `json:"shipping_category_id"`
	TaxonIds           []int             `json:"taxon_ids" sql:"-"`
	TotalOnHand        int64             `json:"total_on_hand" sql:"-"`
	HasVariants        bool              `json:"has_variants" sql:"-"`
	Master             *Variant          `json:"master" sql:"-"`
	Variants           []*Variant        `json:"variants" sql:"-"`
	OptionTypes        []OptionType      `json:"option_types" sql:"-"`
	ProductProperties  []ProductProperty `json:"product_properties" sql:"-"`
	Classifications    []Classification  `json:"classifications" sql:"-"`

	DeletedAt     time.Time `json:"-"`
	TaxCategoryId int64     `json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	Promotionable bool      `json:"-"`
	MetaTitle     string    `json:"-"`

	DisplayPrice string `json:"display_price" sql:"-"`
}

func (this *Product) Currency() string {
	return this.Master.Currency()
}

type PermittedProductParams struct {
	Id                 int64     `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Price              float64   `json:"price,string"`
	AvailableOn        time.Time `json:"available_on"`
	MetaDescription    string    `json:"meta_description"`
	MetaKeyWords       string    `json:"meta_keywords"`
	ShippingCategoryId int64     `json:"shipping_category_id"`
	ShippingCategory   string    `json:"shipping_category"`
}

type ProductParams struct {
	PermittedProductParams *PermittedProductParams `json:"product"`
}

func NewProductFromPermittedParams(productParams *ProductParams) *Product {
	permittedProductParams := productParams.PermittedProductParams
	if permittedProductParams == nil {
		return &Product{}
	}

	product := &Product{
		Id:                 permittedProductParams.Id,
		Name:               permittedProductParams.Name,
		Description:        permittedProductParams.Description,
		Price:              &permittedProductParams.Price,
		AvailableOn:        permittedProductParams.GetAvailableOn(),
		Slug:               strings.Trim(permittedProductParams.Name, " "),
		MetaDescription:    permittedProductParams.MetaDescription,
		MetaKeyWords:       permittedProductParams.MetaKeyWords,
		ShippingCategoryId: permittedProductParams.ShippingCategoryId,
		Promotionable:      true,
	}
	product.Master = NewMasterVariant(product)
	return product
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

func (this *Product) SetComputedValues() {
	this.Price = this.Master.Price
	this.DisplayPrice = Monetize(*this.Price, this.Currency())
}

func (this *Product) Unmarshal(data []byte) error {
	return json.Unmarshal(data, this)
}

func (this *Product) IsValid() bool {
	productErrors = &ValidatorErrors{}

	if this.Name == "" {
		productErrors.Add("name", ErrNotBlank.Error())
	}

	if this.Price == nil && spree.Get(spree.MASTER_PRICE) == "true" {
		productErrors.Add("price", ErrNotBlank.Error())
	}

	if this.ShippingCategoryId == 0 {
		productErrors.Add("shipping_category_id", ErrNotBlank.Error())
	}

	if len(this.Slug) < 3 {
		productErrors.Add("slug", ErrTooShort(3).Error())
	}

	if len(this.MetaKeyWords) > 255 {
		productErrors.Add("meta_keywords", ErrMaxLen(255).Error())
	}

	if len(this.MetaTitle) > 255 {
		productErrors.Add("meta_title", ErrMaxLen(255).Error())
	}

	return productErrors.IsEmpty()
}

func (this *Product) SlugCandidates() []interface{} {
	return []interface{}{
		this.Name,
		[]interface{}{this.Name, this.Master.Sku},
	}
}

func (this *Product) SetSlug(slug string) {
	this.Slug = slug
}

func (this *Product) Errors() *ValidatorErrors {
	if productErrors.IsEmpty() {
		return nil
	}

	return productErrors
}

func (this *Product) BeforeCreate() (err error) {
	if !this.IsValid() {
		err = ErrNotValid
	}
	return
}

func (this *PermittedProductParams) GetAvailableOn() time.Time {
	if this.AvailableOn.IsZero() {
		this.AvailableOn = time.Now()
	}
	return this.AvailableOn
}
