package repositories

import (
	"reflect"
	"testing"
)

func TestAssetRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	assetRepo := NewAssetRepo()

	assets, err := assetRepo.FindByViewableIds([]int64{11})
	if err != nil {
		t.Error("There was an error", err)
	}

	nv := len(assets)

	if nv < 1 {
		t.Error("Invalid number of assets: %d", nv)
		return
	}

	temp := reflect.ValueOf(*assets[0]).Type().String()

	if temp != "json.AssetModel" {
		t.Error("Invalid type", t)
	}
}

func TestAssetRepo_AllImagesByVariantId(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	assetRepo := NewAssetRepo()

	images, err := assetRepo.AllImagesByVariantId(11)
	if err != nil {
		t.Error("There was an error", err)
	}

	if len(images) < 1 {
		t.Errorf("There aren't images for this variant: %d", 11)
	}
}
