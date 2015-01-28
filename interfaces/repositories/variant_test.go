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

	if spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer spree_db.Close()

	variantRepo := NewVariantRepo()

	variants, err := variantRepo.FindByProductId(1)
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	nv := len(variants)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
	}

	temp := reflect.ValueOf(variants[0]).Type().String()

	if temp != "models.Variant" {
		t.Error("Invalid type", t)
	}

}
