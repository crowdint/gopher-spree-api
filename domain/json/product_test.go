package json

import (
	"testing"
	"time"
)

func TestProductStructure(t *testing.T) {
	expected := `{"id":1,"name":"Product1","description":` +
		`"A prodcut","price":"10.0","display_price":"$10.0"` +
		`,"available_on":"2013-10-10T00:00:00Z","slug":` +
		`"product1","meta_description":"...","meta_keywords"` +
		`:"...","shipping_category_id":1,"taxon_ids":[1,10]` +
		`,"total_on_hand":10,"has_variants":true,"master":` +
		`{"id":1,"name":"prod1","price":"10,0","sku":"A1233"` +
		`,"weight":0,"height":0,"width":0,"depth":0,"is_master"` +
		`:false,"slug":"","description":"","cost_price":""` +
		`,"display_price":"","options_text":"","in_stock"` +
		`:false,"is_backorderable":false,"total_on_hand":0` +
		`,"is_destroyed":false,"option_values":null,"images":` +
		`null,"track_inventory":false},"variants":[{"id":2` +
		`,"name":"prod1","price":"10,0","sku":"B1233","weight"` +
		`:0,"height":0,"width":0,"depth":0,"is_master":false,` +
		`"slug":"","description":"","cost_price":"","display_price"` +
		`:"","options_text":"","in_stock":false,"is_backorderable"` +
		`:false,"total_on_hand":0,"is_destroyed":false,` +
		`"option_values":null,"images":null,"track_inventory"` +
		`:false},{"id":3,"name":"prod1","price":"10,0","sku":` +
		`"A1234","weight":0,"height":0,"width":0,"depth":0,` +
		`"is_master":false,"slug":"","description":"","cost_price"` +
		`:"","display_price":"","options_text":"","in_stock"` +
		`:false,"is_backorderable":false,"total_on_hand":0,` +
		`"is_destroyed":false,"option_values":null,"images"` +
		`:null,"track_inventory":false}],"option_types":[{"id"` +
		`:90,"name":"option 90","presentation":"pres1","position"` +
		`:99},{"id":91,"name":"option 91","presentation":` +
		`"pres2","position":100}],"product_properties"` +
		`:[{"id":1,"product_id":1,"property_id":91,"value"` +
		`:"prop1","property_name":"name_prop1"}],"classifications"` +
		`:[{"taxon_id":1,"position":20,"taxon":{"id":1,"parent_id"` +
		`:1,"taxonomy_id":0,"position":0,"name":"taxon1",` +
		`"pretty_name":"","permalink":"","taxons":null}}]}`

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
		Master: Variant{
			ID:    1,
			Name:  "prod1",
			Price: "10,0",
			Sku:   "A1233",
		},
		Variants: []Variant{
			{
				ID:    2,
				Name:  "prod1",
				Price: "10,0",
				Sku:   "B1233",
			},
			{
				ID:    3,
				Name:  "prod1",
				Price: "10,0",
				Sku:   "A1234",
			},
		},
		OptionTypes: []OptionType{
			{
				ID:           90,
				Name:         "option 90",
				Presentation: "pres1",
				Position:     99,
			},
			{
				ID:           91,
				Name:         "option 91",
				Presentation: "pres2",
				Position:     100,
			},
		},
		ProductProperties: []ProductProperty{
			{
				ID:           1,
				ProductID:    1,
				PropertyID:   91,
				Value:        "prop1",
				PropertyName: "name_prop1",
			},
		},
		Classifications: []Classification{
			{
				TaxonId:  1,
				Position: 20,
				Taxon: Taxon{
					ID:       1,
					ParentID: 1,
					Name:     "taxon1",
				},
			},
		},
	}

	AssertEqualJson(t, product, expected)
}
