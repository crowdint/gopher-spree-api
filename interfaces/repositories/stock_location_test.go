package repositories

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestStockLocationRepository_AllBy(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	if err = createDefaultStockLocation(); err != nil {
		t.Error("An error occurred while creating default stock location:", err.Error())
	}

	stockLocationRepository := NewStockLocationRepository()
	stockLocations, err := stockLocationRepository.AllBy("propagate_all_variants = ?", true)
	if err != nil {
		t.Error("An error occured while getting all stock locations:", err.Error())
	}

	if len(stockLocations) == 0 {
		t.Error("There should be stock locations in the DB")
	}
}

func createDefaultStockLocation() error {
	stockLocation := &domain.StockLocation{
		Name:                 "defualt",
		Default:              false,
		Active:               true,
		BackorderableDefault: true,
		PropagateAllVariants: true,
	}

	return Spree_db.Create(stockLocation).Error
}
