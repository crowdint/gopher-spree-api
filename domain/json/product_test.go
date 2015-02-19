package json

import (
	"testing"
	"time"
)

func TestProductStructure(t *testing.T) {
	expected := `{"id":1,"name":"Product1","description":"A prodcut","price":"10.0",` +
		`"display_price":"$10.0","available_on":"2013-10-10T00:00:00Z","slug":"product1",` +
		`"meta_description":"...","meta_keywords":"...","shipping_category_id":1,` +
		`"taxon_ids":[1,10],"total_on_hand":10,"has_variants":true,"master":{"id":0,` +
		`"cost_price":"","depth":"0","height":"0","is_master":false,"options_text":"",` +
		`"price":"0","product_id":0,"sku":"","weight":"0","width":"0","description":"",` +
		`"display_price":"","in_stock":false,"is_backorderable":false,"is_destroyed":false,` +
		`"name":"","slug":"","total_on_hand":null,"track_inventory":false,"images":null,` +
		`"option_values":null},"variants":[],"option_types":[],"product_properties":[],"classifications":[]}`

	someTime := time.Date(2013, 10, 10, 0, 0, 0, 0, time.UTC)

	product := Product{
		Name:               "Product1",
		Id:                 1,
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
