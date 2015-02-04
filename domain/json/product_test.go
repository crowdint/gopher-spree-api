package json

import (
	"testing"
	"time"
)

func TestProductStructure(t *testing.T) {
	expected := `{"id":1,"name":"Product1","description":"A prodcut","price":"10.0",` +
		`"display_price":"$10.0","available_on":"2013-10-10T00:00:00Z","slug":"product1",` +
		`"meta_description":"...","meta_keywords":"...","shipping_category_id":1,` +
		`"taxon_ids":[1,10],"total_on_hand":10,"has_variants":true,"master":{"id":0,"name":"",` +
		`"sku":"","price":"","weight":"","height":"","width":"","depth":"","is_master":false,"slug":"",` +
		`"description":"","track_inventory":false,"cost_price":"","display_price":"","options_text":"",` +
		`"in_stock":false,"is_backorderable":false,"total_on_hand":0,"is_destroyed":false,` +
		`"option_values":null,"images":null},"variants":[],"option_types":[],` +
		`"product_properties":[],"classifications":[]}`

	someTime := time.Date(2013, 10, 10, 0, 0, 0, 0, time.UTC)

	product := Product{
		Name:               "Product1",
		ID:                 1,
		Description:        "A prodcut",
		Slug:               "product1",
		MetaDescription:    "...",
		MetaKeyWords:       "...",
		AvailableOn:        someTime,
		ShippingCategoryId: 1,
		Price:              "10.0",
		DisplayPrice:       "$10.0",
		TaxonIds:           []int{1, 10},
		TotalOnHand:        10,
		HasVariants:        true,
		Master:             Variant{},
		Variants:           []Variant{},
		OptionTypes:        []OptionType{},
		ProductProperties:  []ProductProperty{},
		Classifications:    []Classification{},
	}

	AssertEqualJson(t, product, expected)
}
