package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
)

type StockLocationRepository struct {
	DbRepository
}

func NewStockLocationRepository() *StockLocationRepository {
	return &StockLocationRepository{
		DbRepository{Spree_db},
	}
}

func (this *StockLocationRepository) FindByShipmentAssociation(shipment *domain.Shipment) {
	this.Association(shipment, &shipment.StockLocation, "StockLocationId")
}
