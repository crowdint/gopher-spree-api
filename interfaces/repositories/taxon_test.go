package repositories

import (
	"reflect"
	"testing"

	. "github.com/crowdint/gransak"
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

	taxonSlice, err := taxonRepo.List(1, 10, "", []interface{}{})
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	nv := len(taxonSlice)

	if nv > 10 || nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*taxonSlice[0]).Type().String()

	if temp != "models.Taxon" {
		t.Error("Invalid type", t)
	}

	count, err := taxonRepo.CountAll("", []interface{}{})
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	if count < 1 {
		t.Error("dbHandler error:", err)
	}
}

func TestTaxonRepo_WithGransakQuery(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonRepo := NewTaxonRepo()

	whereStr, params := Gransak.ToSql("name_cont", "Rails")

	taxonSlice, err := taxonRepo.List(1, 10, whereStr, params)
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	nv := len(taxonSlice)

	if nv > 10 || nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*taxonSlice[0]).Type().String()

	if temp != "models.Taxon" {
		t.Error("Invalid type", t)
	}

	count, err := taxonRepo.CountAll("", []interface{}{})
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	if count < 1 {
		t.Error("dbHandler error:", err)
	}
}

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

	if temp != "models.Taxon" {
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

	if temp != "models.Taxon" {
		t.Error("Invalid type", t)
	}
}
