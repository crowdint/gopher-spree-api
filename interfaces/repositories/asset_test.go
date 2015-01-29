package repositories

import (
	"os"
	"reflect"
	"testing"
)

func TestAssetRepo(t *testing.T) {
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

	assetRepo := NewAssetRepo()

	assets := assetRepo.FindByViewableIds([]int64{11})

	nv := len(assets)

	if nv < 1 {
		t.Error("Invalid number of assets: %d", nv)
		return
	}

	temp := reflect.ValueOf(*assets[0]).Type().String()

	if temp != "models.Asset" {
		t.Error("Invalid type", t)
	}
}
