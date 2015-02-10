package json

import "testing"

func TestVariantStructure(t *testing.T) {
	expected := `{"id":1,"name":"var1","sku":"A1234","price":"10.0","weight":"11.1",` +
		`"height":"20.0","width":"21.0","depth":"23.1","is_master":true,"slug":"var1",` +
		`"description":"variant 1","track_inventory":true,"cost_price":"1.9",` +
		`"display_price":"$10.0","options_text":"opt 1","in_stock":true,"is_backorderable":false,` +
		`"total_on_hand":90,"is_destroyed":false,"option_values":[],"images":[]}`

	variant := Variant{
		ID:              1,
		Name:            "var1",
		Sku:             "A1234",
		Price:           "10.0",
		Weight:          "11.1",
		Height:          "20.0",
		Width:           "21.0",
		Depth:           "23.1",
		IsMaster:        true,
		Slug:            "var1",
		Description:     "variant 1",
		TrackInventory:  true,
		CostPrice:       "1.9",
		DisplayPrice:    "$10.0",
		OptionsText:     "opt 1",
		InStock:         true,
		IsBackorderable: false,
		TotalOnHand:     90,
		IsDestroyed:     false,
		OptionValues:    []*OptionValue{},
		Images:          []*Asset{},
	}

	AssertEqualJson(t, variant, expected)
}
