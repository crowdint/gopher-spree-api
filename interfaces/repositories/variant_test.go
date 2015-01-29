package repositories

import (
	"reflect"
	"testing"
)

func TestVariantRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	variantRepo := NewVariantRepo()

	variants, err := variantRepo.FindByProductIds([]int64{1, 2, 3})
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	nv := len(variants)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*variants[0]).Type().String()

	if temp != "models.Variant" {
		t.Error("Invalid type", t)
	}

}
