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

	jsonProductSlice, err := productInteractor.GetResponse(1, 10)
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

	if p1.ID != 1 || p1.Name != "name1" {
		t.Error("Incorrect product values")
	}
}

func TestProductInteractor_toJson(t *testing.T) {
	product := &models.Product{
		Id:          1,
		Name:        "name1",
		Description: "desc1",
	}

	productInteractor := NewProductInteractor()

	jsonProduct := productInteractor.toJson(product)

	if jsonProduct.ID != 1 || jsonProduct.Name != "name1" || jsonProduct.Description != "desc1" {
		t.Error("incorrect json.Product values")
	}
}

func TestProductInteractor_mergeVariants(t *testing.T) {
	jsonProductSlice := []*jsn.Product{
		&jsn.Product{
			ID: 1,
		},
		&jsn.Product{
			ID: 2,
		},
	}

	jsonVariantsMap := JsonVariantsMap{
		1: []*jsn.Variant{
			{
				ID:   1,
				Name: "variant1",
			},
		},
		2: []*jsn.Variant{
			{
				ID:       2,
				Name:     "variant2",
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

	if v1.ID != 1 || v1.Name != "variant1" || v1.IsMaster {
		t.Errorf("Incorrect variant values %d %s %b", v1.ID, v1.Name, v1.IsMaster)
	}
}
