package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
)

type ShippingRateRepository struct {
	DbRepository
}

func NewShippingRateRepository() *ShippingRateRepository {
	return &ShippingRateRepository{
		DbRepository{Spree_db},
	}
}

func (this *ShippingRateRepository) AllByShipment(shipment *json.Shipment) []json.ShippingRate {
	shippingRates := []json.ShippingRate{}
	this.All(&shippingRates, map[string]interface{}{
		"order": "cost ASC",
	}, "shipment_id = ?", shipment.Id)

	for i, _ := range shippingRates {
		shippingMethodRepository := NewShippingMethodRepository()
		shippingMethodRepository.FindByShippingRateAssociation(&shippingRates[i])
		shippingRates[i].Name = shippingRates[i].ShippingMethod.Name
		shippingRates[i].ShippingMethodCode = &shippingRates[i].ShippingMethod.Code
		shipment.ShippingMethods = append(shipment.ShippingMethods, shippingRates[i].ShippingMethod)

		if shippingRates[i].Selected {
			shipment.SelectedShippingRate = shippingRates[i]
		}
	}

	return shippingRates
}
