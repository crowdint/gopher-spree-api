package repositories

import (
	"os"
	"testing"
)

func TestVariantRepo(t *testing.T) {
	os.Setenv(dbUrlEnvName, "dbname=spree_dev sslmode=disable")
	os.Setenv(dbmsEnvName, "postgres")

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

	t.Error(variants)

}
