package json

import (
	"encoding/json"
	"testing"

	jsn "github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductInteractor_GetMergedResponse(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	productInteractor := NewProductInteractor()

	jsonProductSlice, err := productInteractor.GetResponse(1, 10, "")
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if jsonProductSlice.(ContentResponse).GetCount() < 1 {
		t.Error("Error: Invalid number of rows")
		return
	}

	jsonBytes, err := json.Marshal(jsonProductSlice)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if string(jsonBytes) == "" {
		t.Error("Error: Json string is empty")
	}
}

func TestProductInteractor_getIdSlice(t *testing.T) {
	products := []*models.Product{
		&models.Product{
			Id: 1,
		},
		&models.Product{
			Id: 2,
		},
		&models.Product{
			Id: 3,
		},
	}

	productInteractor := NewProductInteractor()

	ids := productInteractor.getIdSlice(products)

	if len(ids) != 3 {
		t.Error("Incorrect number of ids")
	}

	if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Error("Incorrect id value")
	}
}

func TestProductInteractor_modelsToJsonProductsSlice(t *testing.T) {
	products := []*models.Product{
		&models.Product{
			Id:   1,
			Name: "name1",
		},
		&models.Product{
			Id:   2,
			Name: "name2",
		},
		&models.Product{
			Id:   3,
			Name: "name3",
		},
	}

	productInteractor := NewProductInteractor()

	jsonProducts := productInteractor.modelsToJsonProductsSlice(products)

	if len(jsonProducts) < 1 {
		t.Error("Incorrect product ids slice lenght")
	}

	p1 := jsonProducts[0]

	if p1.Id != 1 || p1.Name != "name1" {
		t.Error("Incorrect product values")
	}
}

func TestProductInteractor_mergeVariants(t *testing.T) {
	jsonProductSlice := []*jsn.Product{
		&jsn.Product{
			Id:   99991,
			Name: "Product1",
		},
		&jsn.Product{
			Id:   99992,
			Name: "Product2",
		},
	}

	jsonVariantsMap := JsonVariantsMap{
		99991: []*jsn.Variant{
			{
				Id: 99991,
			},
		},
		99992: []*jsn.Variant{
			{
				Id:       99992,
				IsMaster: true,
			},
		},
	}

	productInteractor := NewProductInteractor()

	productInteractor.mergeVariants(jsonProductSlice, jsonVariantsMap)

	p2 := jsonProductSlice[0]

	if p2.Variants == nil {
		t.Error("Product variants are nil")
		return
	}

	if len(p2.Variants) == 0 {
		t.Error("No product variants found")
		return
	}

	v1 := p2.Variants[0]

	if v1.Id != 99991 || v1.Name != "Product1" || v1.IsMaster {
		t.Errorf("Incorrect variant values %d %s %b", v1.Id, v1.Name, v1.IsMaster)
	}
}

func TestProductInteractor_mergeOptionTypes(t *testing.T) {
	jsonProductSlice := []*jsn.Product{
		&jsn.Product{
			Id: 3,
		},
	}

	jsonOptionTypesMap := JsonOptionTypesMap{
		3: []*jsn.OptionType{
			{
				Id:           1,
				Name:         "tshirt-size",
				Presentation: "Size",
			},
			{
				Id:           2,
				Name:         "tshirt-color",
				Presentation: "Color",
			},
		},
	}

	productInteractor := NewProductInteractor()

	productInteractor.mergeOptionTypes(jsonProductSlice, jsonOptionTypesMap)

	product := jsonProductSlice[0]

	if product.OptionTypes == nil {
		t.Error("Product OptionTypes are nil")
		return
	}

	if len(product.OptionTypes) == 0 {
		t.Error("No product optionTypes found")
		return
	}

	optionType1 := product.OptionTypes[0]

	if optionType1.Id != 1 || optionType1.Name != "tshirt-size" || optionType1.Presentation != "Size" {
		t.Errorf("Incorrect optionType values: \n Id -> %d, Name -> %s, Presentation -> %d", optionType1.Id, optionType1.Name, optionType1.Presentation)
	}
}

func TestProductInteractor_mergeClassifications(t *testing.T) {
	jsonProductSlice := []*jsn.Product{
		&jsn.Product{
			Id: 3,
		},
		&jsn.Product{
			Id: 5,
		},
	}

	jsonOptionTypesMap := JsonClassificationsMap{
		3: []*jsn.Classification{
			{
				TaxonId:  1,
				Position: 5,
				Taxon: jsn.Taxon{
					Id:   1,
					Name: "taxonName",
				},
			},
		},
	}

	productInteractor := NewProductInteractor()

	productInteractor.mergeClassifications(jsonProductSlice, jsonOptionTypesMap)

	product1 := jsonProductSlice[0]
	product2 := jsonProductSlice[1]

	if product1.Classifications == nil || product2.Classifications == nil {
		t.Error("Product.Classifications should be and empty slice [] at least")
	}

	classification := product1.Classifications[0]

	if classification.TaxonId != 1 || classification.Taxon.Id != 1 {
		t.Error("Wrong assignment of classifications")
	}

	if len(product2.Classifications) > 0 {
		t.Error("Wrong assignment of classficiations")
	}

	if product1.TaxonIds[0] != 1 {
		t.Error("Wrong assignment of taxon ids")
	}

	if len(product2.TaxonIds) != 0 {
		t.Error("Wrong assignment of taxon ids")
	}

}
