package json

import (
	"encoding/json"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestTaxonInteractor_GetResponse(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Taxon{Id: 1})
	repositories.Spree_db.Exec("INSERT INTO spree_products_taxons(taxon_id, product_id) values(1, 1)")

	taxonInteractor := NewTaxonInteractor()

	jsonTaxonSlice, err := taxonInteractor.GetResponse(1, 10, &FakeResponseParameters{})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if jsonTaxonSlice.(ContentResponse).GetCount() < 1 {
		t.Error("Error: Invalid number of rows")
		return
	}

	jsonBytes, err := json.Marshal(jsonTaxonSlice)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if string(jsonBytes) == "" {
		t.Error("Error: Json string is empty")
	}
}
