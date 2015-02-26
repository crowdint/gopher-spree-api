package repositories

import (
	"reflect"
	"testing"
)

func TestTaxonRepo_FindByProductIds(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonRepo := NewTaxonRepo()

	taxonsByProduct, err := taxonRepo.FindByProductIds([]int64{1, 2})

	if err != nil {
		t.Error("An error has ocurred", err)
		return
	}

	temp := reflect.ValueOf(*taxonsByProduct[0]).Type().String()

	if temp != "domain.Taxon" {
		t.Error("Invalid type", t)
	}

}

func TestTaxonRepo_FindByTaxonomyIds(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonRepo := NewTaxonRepo()

	taxonsByTaxonomy, err := taxonRepo.FindByTaxonomyIds([]int64{1, 2})

	if err != nil {
		t.Error("An error has ocurred", err)
		return
	}

	temp := reflect.ValueOf(*taxonsByTaxonomy[0]).Type().String()

	if temp != "domain.Taxon" {
		t.Error("Invalid type", t)
	}
}
