package json

import (
	"encoding/json"
	"testing"

	jsn "github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestTaxonInteractor_GetResponse(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	taxonInteractor := NewTaxonInteractor()

	jsonTaxonSlice, err := taxonInteractor.GetResponse(1, 10, &FakeResponseParameters{})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if jsonTaxonSlice.(ContentResponse).GetCount() < 1 {
		t.Error("Error: Invalid number of rows")
		return
	}

	jsonBytes, err := domain.Marshal(jsonTaxonSlice)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if string(jsonBytes) == "" {
		t.Error("Error: Json string is empty")
	}
}

func TestTaxonInteractor_modelsToJsonTaxonsSlice(t *testing.T) {
	taxons := []*jsn.Taxon{
		&jsn.Taxon{
			Id:   1,
			Name: "Categories",
		},
		&jsn.Taxon{
			Id:   2,
			Name: "Bags",
		},
		&jsn.Taxon{
			Id:   3,
			Name: "Mugs",
		},
	}

	taxonInteractor := NewTaxonInteractor()

	jsonTaxons := taxonInteractor.modelsToJsonTaxonsSlice(taxons)

	if len(jsonTaxons) < 1 {
		t.Error("Incorrect taxon ids slice length")
	}

	p1 := jsonTaxons[0]

	if p1.Id != 1 || p1.Name != "Categories" {
		t.Error("Incorrect taxon values")
	}
}
