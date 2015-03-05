package domain

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

func TestProductValidator(t *testing.T) {
	p := &Product{}

	if p.IsValid() {
		t.Error("Product should be invalid")
	}

	if p.GetErrors().Size() != 4 {
		t.Errorf("Product should have 4 errors, but has %d", p.GetErrors().Size())
	}

	p.Name = "Test Product"

	if p.IsValid() {
		t.Error("Product should be invalid")
	}

	if p.GetErrors().Size() != 3 {
		t.Errorf("Product should have 3 errors, but has %d", p.GetErrors().Size())
	}

	p.Price = "3"

	if p.IsValid() {
		t.Error("Product should be invalid")
	}

	if p.GetErrors().Size() != 2 {
		t.Errorf("Product should have 2 errors, but has %d", p.GetErrors().Size())
	}

	p.Slug = "test-product"

	if p.IsValid() {
		t.Error("Product should be invalid")
	}

	if p.GetErrors().Size() != 1 {
		t.Errorf("Product should have 1 error, but has %d", p.GetErrors().Size())
	}

	p.ShippingCategoryId = 3

	if !p.IsValid() {
		t.Error("Product should be valid")
	}

	if productErrors.Size() != 0 {
		t.Errorf("Product should not have errors, but has %d", productErrors.Size())
	}
}

func TestNewProductFromPermittedParams(t *testing.T) {
	permittedProductParams := &PermittedProductParams{
		Name:               "Test Product",
		Description:        "Test Description",
		Price:              "12.40",
		ShippingCategoryId: 3,
	}

	product := NewProductFromPermittedParams(&ProductParams{permittedProductParams})

	if product.Name != permittedProductParams.Name {
		t.Errorf("Product Name should be %s, but was %s", permittedProductParams.Name, product.Name)
	}

	if product.Description != permittedProductParams.Description {
		t.Errorf("Product Description should be %s, but was %s", permittedProductParams.Description, product.Description)
	}

	if product.Price != permittedProductParams.Price {
		t.Errorf("Product Price should be %s, but was %s", permittedProductParams.Price, product.Price)
	}
}

func TestPermittedParams_GetAvailableOn(t *testing.T) {
	permittedProductParams := &PermittedProductParams{
		Name:               "Test Product",
		Description:        "Test Description",
		Price:              "12.40",
		ShippingCategoryId: 3,
	}

	permittedProductParams.GetAvailableOn()
	if permittedProductParams.AvailableOn.IsZero() {
		t.Error("Available On time should not be zero")
	}
}
