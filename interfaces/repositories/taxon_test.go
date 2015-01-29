package repositories

import (
	"os"
	"reflect"
	"testing"
)

func TestTaxonRepo(t *testing.T) {
	os.Setenv(DbUrlEnvName, "dbname=spree_dev sslmode=disable")
	os.Setenv(DbEngineEnvName, "postgres")

	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	taxonRepo := NewTaxonRepo()

	taxons := taxonRepo.List()

	nv := len(taxons)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*taxons[0]).Type().String()

	if temp != "models.Taxon" {
		t.Error("Invalid type", t)
	}

}
