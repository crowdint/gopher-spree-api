package domain

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/configs/spree"
)

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

func TestVariantValidator(t *testing.T) {
	costPrice := "-1.0"
	price := -1.0

	variant := &Variant{CostPrice: &costPrice, Price: &price}

	if variant.IsValid() {
		t.Error("Variant should be invalid")
	}

	if variant.Errors().Size() != 2 {
		t.Errorf("Variant should have 2 errors, but has %d", variant.Errors().Size())
	}

	costPrice = "2.0"
	variant.CostPrice = &costPrice

	if variant.IsValid() {
		t.Error("Variant should be invalid")
	}

	if variant.Errors().Size() != 1 {
		t.Errorf("Variant should have 1 error, but has %d", variant.Errors().Size())
	}

	price = 2.59
	variant.Price = &price

	if !variant.IsValid() {
		t.Error("Variant should be valid")
	}

	if variantErrors.Size() != 0 {
		t.Errorf("Variant should not have errors, but has %d", variantErrors.Size())
	}

}

func TestVariant_CheckPrice(t *testing.T) {
	product := &Product{
		Name:        "Product Test",
		Description: "Product Description",
		Price:       "12.79",
	}

	product.Master = *NewMasterVariant(product)
	variant := Variant{}
	spree.Set(spree.MASTER_PRICE, "true")

	if err := variant.checkPrice(); err == nil {
		t.Error("Should have 'No master variant found to infer price' error")
	}

	variant.Product = product
	variant.Product.Master.IsMaster = false
	if err := variant.checkPrice(); err == nil {
		t.Error("Should have 'No master variant found to infer price' error")
	}

	variant.Product.Master.IsMaster = true

	tempMaster := variant.Product.Master
	price := tempMaster.Price

	tempMaster.Price = nil
	variant.Product.Master.Price = nil
	if err := tempMaster.checkPrice(); err == nil {
		t.Error("Should have 'Must supply price for variant or master.price for product.' error")
	}

	variant.Product.Master.Price = price
	if err := tempMaster.checkPrice(); err != nil {
		t.Error("Should not have errors, but", err.Error())
	}

	if *tempMaster.Price != *price {
		t.Error("Price should be %f, but was %f", *price, *tempMaster.Price)
	}

	if tempMaster.DefaultPrice.Currency == "" {
		t.Error("Currency was not set")
	}
}
