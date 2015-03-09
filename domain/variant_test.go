package domain

import "testing"

func TestVariantStructure(t *testing.T) {
	expected := `{"id":1,"cost_price":"1.9","depth":"23.1","height":"20","is_master":true,` +
		`"options_text":"opt 1","price":"10.1","product_id":0,"sku":"A1234","weight":"11.1",` +
		`"width":"21","description":"variant 1","display_price":"$10.1","in_stock":true,` +
		`"is_backorderable":false,"is_destroyed":false,"name":"var1","slug":"var1","total_on_hand":90,` +
		`"track_inventory":true,"images":[],"option_values":[]}`

	var totalOnHand int64 = 90
	costPrice := "1.9"
	price := 10.1
	variant := Variant{
		Id:              1,
		Name:            "var1",
		Sku:             "A1234",
		Price:           &price,
		Weight:          11.1,
		Height:          20.0,
		Width:           21.0,
		Depth:           23.1,
		IsMaster:        true,
		Slug:            "var1",
		Description:     "variant 1",
		TrackInventory:  true,
		CostPrice:       &costPrice,
		DisplayPrice:    "$10.1",
		OptionsText:     "opt 1",
		InStock:         true,
		IsBackorderable: false,
		TotalOnHand:     &totalOnHand,
		IsDestroyed:     false,
		OptionValues:    []OptionValue{},
		Images:          []*Asset{},
	}

	AssertEqualJson(t, variant, expected)
}

func TestVariant_NewMasterVariant(t *testing.T) {
	productWithPrice := &Product{
		Name:        "Product Name",
		Description: "Product Description",
		Price:       "12.45",
	}

	master := NewMasterVariant(productWithPrice)

	if !master.IsMaster {
		t.Error("Master.IsMaster should be true")
	}

	if master.Product != productWithPrice {
		t.Error("Master.Product should be equals to Product with price")
	}

	if master.ProductId != productWithPrice.Id {
		t.Errorf("Master.ProductId should be equals to %d", productWithPrice.Id)
	}

	if master.DefaultPrice.Amount != 12.45 {
		t.Errorf("Master default price amount should be 12.45, but was %f", master.DefaultPrice.Amount)
	}

	if *master.Price != 12.45 {
		t.Errorf("Master price should be 12.45, but was %f", *master.Price)
	}

	if *master.Position != 1 {
		t.Error("Master positon should be 1, but was", master.Position)
	}

	productWithoutPrice := &Product{
		Name:        "Product Name",
		Description: "Product Description",
	}

	master = NewMasterVariant(productWithoutPrice)

	if !master.IsMaster {
		t.Error("Master.IsMaster should be true")
	}

	if master.Product != productWithoutPrice {
		t.Error("Master.Product should be equals to Product without price")
	}

	if master.ProductId != productWithoutPrice.Id {
		t.Errorf("Master.ProductId should be equals to %d", productWithoutPrice.Id)
	}

	if master.DefaultPrice.Amount != 0 {
		t.Errorf("Master default price amount should be 0, but was %d", master.DefaultPrice.Amount)
	}

	if master.Price != nil {
		t.Error("Master price should be nil, but was", master.Price)
	}

	if *master.Position != 1 {
		t.Error("Master positon should be 1, but was", master.Position)
	}
}
