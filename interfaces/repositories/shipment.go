package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/utils"
)

type ShipmentRepository struct {
	DbRepository
}

func NewShipmentRepository() *ShipmentRepository {
	return &ShipmentRepository{
		DbRepository{Spree_db},
	}
}

func (this *ShipmentRepository) AllByOrder(order *json.Order) []json.Shipment {
	shipments := []json.Shipment{}
	this.All(&shipments, nil, "order_id = ?", order.Id)
	lineItemsMap := utils.ToMap(*(order.LineItems), "Id", false)

	for i, _ := range shipments {
		shipments[i].OrderId = order.Number

		stockLocationRepository := NewStockLocationRepository()
		stockLocationRepository.FindByShipmentAssociation(&shipments[i])
		shipments[i].StockLocationName = shipments[i].StockLocation.Name

		shippingRateRepository := NewShippingRateRepository()
		shipments[i].ShippingRates = shippingRateRepository.AllByShipment(&shipments[i])

		inventoryUnitRepository := NewInventoryUnitRepository()
		inventoryUnitRepository.AllByShipmentAssociation(&shipments[i])
		for j := 0; j < len(shipments[i].Manifest); j++ {
			shipments[i].Manifest[j].Quantity = lineItemsMap[shipments[i].Manifest[j].LineItemId].(json.LineItem).Quantity
		}

		adjustmentRepository := NewAdjustmentRepository()
		shipments[i].Adjustments = adjustmentRepository.AllByAdjustable(shipments[i].Id, "Spree::Shipment")
	}
	return shipments
}
