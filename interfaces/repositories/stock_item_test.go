package repositories

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestStockItemRepository_Create(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	stockItem := &domain.StockItem{VariantId: 1, StockLocationId: 1, Backorderable: true}
	stockItemRepository := NewStockItemRepository()

	if err = stockItemRepository.Create(stockItem); err != nil {
		t.Error("An error occured while creating the stock item")
	}
}
