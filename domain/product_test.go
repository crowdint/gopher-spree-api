package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestProductStructure(t *testing.T) {
	expected := `{"name":"Product1","id":1,"description":` +
		`"A prodcut","slug":"product1","meta_description":"..."` +
		`,"meta_keywords":"...","available_on":"2013-10-10T00:00:00Z"` +
		`,"shipping_category_id":1}`

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

	jsonData, err := json.Marshal(&product)
	if err != nil {
		t.Errorf("An error has ocurred", err.Error())
	}

	jsonString := string(jsonData)

	if jsonString != expected {
		t.Errorf("Mismacth: ", jsonString)
	}
}
