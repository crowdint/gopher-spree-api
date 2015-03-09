package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
)

type ShippingRateRepository struct {
	DbRepository
}

func NewShippingRateRepository() *ShippingRateRepository {
	return &ShippingRateRepository{
		DbRepository{Spree_db},
	}
}

func (this *ShippingRateRepository) AllByShipment(shipment *domain.Shipment) []*domain.ShippingRate {
	shippingRates := []*domain.ShippingRate{}

	this.All(&shippingRates, map[string]interface{}{
		"order": "cost ASC",
	}, "shipment_id = ?", shipment.Id)

	shippingMethodRepository := NewShippingMethodRepository()

	for _, shippingRate := range shippingRates {
		shippingMethodRepository.FindByShippingRateAssociation(shippingRate)
		shippingRate.Name = shippingRate.ShippingMethod.Name
		shippingRate.ShippingMethodCode = &shippingRate.ShippingMethod.Code
		shippingRate.Shipment = shipment
		shippingRate.SetComputedValues()

		shipment.ShippingMethods = append(shipment.ShippingMethods, shippingRate.ShippingMethod)

		if shippingRate.Selected {
			shipment.SelectedShippingRate = *shippingRate
		}
	}

	return shippingRates
}
