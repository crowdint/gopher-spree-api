package json

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/models"
)

func TestVariantStructure(t *testing.T) {
	expected := `{"id":1,"cost_price":"1.9","depth":"23.1","height":"20","is_master":true,` +
		`"options_text":"opt 1","price":"10.1","product_id":0,"sku":"A1234","weight":"11.1",` +
		`"width":"21","description":"variant 1","display_price":"$10.1","in_stock":true,` +
		`"is_backorderable":false,"is_destroyed":false,"name":"var1","slug":"var1","total_on_hand":90,` +
		`"track_inventory":true,"images":[],"option_values":[]}`

	var totalOnHand int64 = 90
	variant := Variant{
		Id:              1,
		Name:            "var1",
		Sku:             "A1234",
		Price:           10.1,
		Weight:          11.1,
		Height:          20.0,
		Width:           21.0,
		Depth:           23.1,
		IsMaster:        true,
		Slug:            "var1",
		Description:     "variant 1",
		TrackInventory:  true,
		CostPrice:       "1.9",
		DisplayPrice:    "$10.1",
		OptionsText:     "opt 1",
		InStock:         true,
		IsBackorderable: false,
		TotalOnHand:     &totalOnHand,
		IsDestroyed:     false,
		OptionValues:    []models.OptionValue{},
		Images:          []*Asset{},
	}

	AssertEqualJson(t, variant, expected)
}
