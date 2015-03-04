package json

import (
	"encoding/json"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestTaxonomyInteractor_GetResponse(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Taxonomy{Id: 1})
	repositories.Spree_db.Create(&domain.Taxonomy{Id: 2})
	repositories.Spree_db.Create(&domain.Taxon{Id: 1, TaxonomyId: 1, Name: "Brands"})
	repositories.Spree_db.Create(&domain.Taxon{Id: 2, TaxonomyId: 2, Name: "Categories"})
	repositories.Spree_db.Create(&domain.Taxon{Id: 3, TaxonomyId: 2, ParentId: 1, Name: "Bags"})

	taxonInteractor := NewTaxonomyInteractor()

	jsonTaxonomySlice, err := taxonInteractor.GetResponse(1, 10, &FakeResponseParameters{})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if jsonTaxonomySlice.(ContentResponse).GetCount() < 1 {
		t.Error("Error: Invalid number of rows")
		return
	}

	jsonBytes, err := json.Marshal(jsonTaxonomySlice)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if string(jsonBytes) == "" {
		t.Error("Error: Json string is empty")
	}
}
