package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestTaxonRepo_FindByProductIds(t *testing.T) {
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	Spree_db.Create(&domain.Taxon{Id: 1})
	Spree_db.Exec("INSERT INTO spree_products_taxons(taxon_id, product_id) values(1, 1)")

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
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	Spree_db.Create(&domain.Taxon{Id: 1, TaxonomyId: 1})
	Spree_db.Exec("INSERT INTO spree_taxonomies(id, name) values(1, 'foo')")

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
