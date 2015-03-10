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

func (this *StockLocationRepository) AllBy(query interface{}, values ...interface{}) ([]*domain.StockLocation, error) {
	stockLocations := []*domain.StockLocation{}

	err := this.All(&stockLocations, nil, query, values...)
	if err != nil {
		return stockLocations, err
	}

	return stockLocations, nil
}
