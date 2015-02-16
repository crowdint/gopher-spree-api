package repositories

import (
	"reflect"
	"testing"

	. "github.com/crowdint/gransak"
)

func TestTaxonomyRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonomyRepo := NewTaxonomyRepo()

	taxonomieslice, err := taxonomyRepo.List(1, 10, "", []interface{}{})
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	nv := len(taxonomieslice)

	if nv > 10 || nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*taxonomieslice[0]).Type().String()

	if temp != "models.Taxonomy" {
		t.Error("Invalid type", t)
	}

	count, err := taxonomyRepo.CountAll("", []interface{}{})
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	if count < 1 {
		t.Error("dbHandler error:", err)
	}
}

func TestTaxonomyRepo_WithGransakQuery(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonomyRepo := NewTaxonomyRepo()

	whereStr, params := Gransak.ToSql("name_cont", "Categories")

	taxonomieslice, err := taxonomyRepo.List(1, 10, whereStr, params)
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	nv := len(taxonomieslice)

	if nv > 10 || nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*taxonomieslice[0]).Type().String()

	if temp != "models.Taxonomy" {
		t.Error("Invalid type", t)
	}

	count, err := taxonomyRepo.CountAll("", []interface{}{})
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	if count < 1 {
		t.Error("dbHandler error:", err)
	}
}
