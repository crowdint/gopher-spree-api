package repositories

import (
	"os"
	"reflect"
	"testing"
)

func TestVariantRepo(t *testing.T) {
	os.Setenv(dbUrlEnvName, "dbname=spree_dev sslmode=disable")
	os.Setenv(dbEngineEnvName, "postgres")

	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	variantRepo := NewVariantRepo()

	variants := variantRepo.FindByProductIds([]int64{1, 2, 3})

	nv := len(variants)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
	}

	temp := reflect.ValueOf(*variants[0]).Type().String()

	if temp != "models.Variant" {
		t.Error("Invalid type", t)
	}

}
