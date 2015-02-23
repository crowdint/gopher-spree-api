package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
)

type StockLocationRepository struct {
	DbRepository
}

func NewStockLocationRepository() *StockLocationRepository {
	return &StockLocationRepository{
		DbRepository{Spree_db},
	}
}

func (this *StockLocationRepository) FindByShipmentAssociation(shipment *json.Shipment) {
	this.Association(shipment, &shipment.StockLocation, "StockLocationId")
}
