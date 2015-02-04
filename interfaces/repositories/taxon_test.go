package repositories

import (
	"reflect"
	"testing"
)

func TestTaxonRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonRepo := NewTaxonRepo()

	taxons, err := taxonRepo.FindByProductIds([]int64{1, 2})

	if err != nil {
		t.Error("An error has ocurred", err)
		return
	}

	temp := reflect.ValueOf(*taxons[0]).Type().String()

	if temp != "models.Taxon" {
		t.Error("Invalid type", t)
	}

}
