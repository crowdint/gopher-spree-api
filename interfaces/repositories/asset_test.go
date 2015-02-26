package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/models"
)

func TestAssetRepo(t *testing.T) {
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	asset := &models.Asset{
		ViewableId: 11,
	}

	Spree_db.Save(asset)

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	assetRepo := NewAssetRepo()

	assets, err := assetRepo.FindByViewableIds([]int64{11})
	if err != nil {
		t.Error("There was an error", err)
		return
	}

	nv := len(assets)

	if nv < 1 {
		t.Error("Invalid number of assets: %d", nv)
		return
	}

	temp := reflect.ValueOf(*assets[0]).Type().String()

	if temp != "domain.AssetModel" {
		t.Error("Invalid type", t)
	}

}

func TestAssetRepo_AllImagesByVariantId(t *testing.T) {
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	asset := &models.Asset{
		ViewableId:   11,
		ViewableType: "Spree::Variant",
		Type:         "Spree::Image",
	}

	Spree_db.Save(asset)

	if err != nil {
		t.Error("An error has ocurred", err)
		return
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
		return
	}

	assetRepo := NewAssetRepo()

	images, err := assetRepo.AllImagesByVariantId(11)
	if err != nil {
		t.Error("There was an error", err)
		return
	}

	if len(images) < 1 {
		t.Errorf("There aren't images for this variant: %d", 11)
	}

}
