package json

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestClassificationInteractor_GetJsonClassificationsMap(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer func() {
		repositories.Spree_db.Rollback()
		repositories.Spree_db.Close()
	}()

	repositories.Spree_db.Create(&domain.Product{Id: 1})
	repositories.Spree_db.Create(&domain.Taxon{Id: 1})
	repositories.Spree_db.Exec("INSERT INTO spree_products_taxons(taxon_id, product_id) VALUES(1, 1)")

	classificationInteractor := NewClassificationInteractor()

	classificationMap, err := classificationInteractor.GetJsonClassificationsMap([]int64{1, 2, 3})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	nclassifications := len(classificationMap)

	if nclassifications < 1 {
		t.Errorf("Wrong number of records %d", nclassifications)
	}

}
