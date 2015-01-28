package repositories

import (
	"os"
	"reflect"
	"testing"
)

func TestTaxonRepo(t *testing.T) {
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

	taxonRepo := NewTaxonRepo()

	taxons := taxonRepo.List()

	nv := len(taxons)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
	}

	temp := reflect.ValueOf(*taxons[0]).Type().String()

	if temp != "models.Taxon" {
		t.Error("Invalid type", t)
	}

}
