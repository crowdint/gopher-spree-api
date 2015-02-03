package json

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"

	"testing"
)

func TestTaxonInteractor_GetJsonTaxonsMap(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	taxonInteractor := NewTaxonInteractor()

	taxonMap, err := taxonInteractor.GetJsonTaxonsMap([]int64{1, 2, 3})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	ntaxons := len(taxonMap)

	if ntaxons < 1 {
		t.Errorf("Wrong number of records %d", ntaxons)
	}

	tarray1 := taxonMap[1]
	if len(tarray1) < 1 {
		t.Error("No taxons found")
	}

	t1 := tarray1[0]

	if t1.PrettyName == "" {
		t.Error("Taxon has no pretty name")
	}
}

func TestTaxonInteractor_modelsToJsonTaxonsMap(t *testing.T) {
	taxonSlice := []*models.Taxon{
		&models.Taxon{
			Id:         3,
			Name:       "Bags",
			PrettyName: "Category -> Bags",
			Permalink:  "category/bags",
			ParentId:   1,
			TaxonomyId: 1,
		},
	}

	taxonInteractor := NewTaxonInteractor()

	jsonTaxonMap := taxonInteractor.modelsToJsonTaxonsMap(taxonSlice)

	t1 := jsonTaxonMap[3][0]

	if t1 == nil {
		t.Error("Error: nil value on map")
	}

	if t1.ID != 3 || t1.Name != "Bags" || t1.PrettyName != "Category -\u003e Bags" || t1.Permalink != "category/bags" || t1.ParentID != 1 || t1.TaxonomyID != 1 {
		t.Error("Invalid values for taxon")
	}
}

func TestTaxonInteractor_toJson(t *testing.T) {
	taxon := &models.Taxon{
		Id:         11,
		Name:       "Rails",
		PrettyName: "Brand -> Rails",
		Permalink:  "brand/rails",
		ParentId:   2,
		TaxonomyId: 2,
	}

	taxonInteractor := NewTaxonInteractor()

	jsonTaxon := taxonInteractor.toJson(taxon)

	if jsonTaxon.ID != 1 {
		t.Error("Invalid values for json.Taxon")
	}
}
