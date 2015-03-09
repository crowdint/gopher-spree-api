package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestVariantRepo(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	costPrice := "10"
	Spree_db.Create(&domain.Variant{Id: 1, ProductId: 1, CostPrice: &costPrice, DefaultPrice: domain.Price{}})
	Spree_db.Exec("INSERT INTO spree_stock_items(variant_id) values(1)")
	Spree_db.Exec("INSERT INTO spree_prices(variant_id, currency) values(1, 'USD')")

	variantRepo := NewVariantRepository()

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

	if temp != "domain.Variant" {
		t.Error("Invalid type", t)
	}
}

func TestVariantRepository_Create(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	costPrice := "10"
	variant := &domain.Variant{Id: 1, ProductId: 1, CostPrice: &costPrice, DefaultPrice: domain.Price{}}

	variantRepository := NewVariantRepository()
	if err = variantRepository.Create(variant); err != nil {
		t.Error("An error occured while creating variant")
	}

	if len(variant.StockItems) == 0 {
		t.Error("Variant should have stock items.")
	}
}

func TestVariantRepository_CreateStockItems(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	costPrice := "10"
	variant := &domain.Variant{Id: 1, ProductId: 1, CostPrice: &costPrice, DefaultPrice: domain.Price{}}

	Spree_db.Create(variant)
	Spree_db.Exec("INSERT INTO spree_prices(variant_id, currency) values(1, 'USD')")

	variantRepository := NewVariantRepository()
	if err = variantRepository.createStockItems(variant); err != nil {
		t.Error("An error occured while creating stock items")
	}

	if len(variant.StockItems) == 0 {
		t.Error("Variant should have stock items.")
	}
}
