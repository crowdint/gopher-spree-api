package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
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

func (this *ShipmentRepository) AllByOrder(order *domain.Order) []domain.Shipment {
	shipments := []domain.Shipment{}
	this.All(&shipments, nil, "order_id = ?", order.Id)
	lineItemsMap := utils.ToMap(*(order.LineItems), "Id", false)

	stockLocationRepository := NewStockLocationRepository()
	shippingRateRepository := NewShippingRateRepository()
	inventoryUnitRepository := NewInventoryUnitRepository()
	adjustmentRepository := NewAdjustmentRepository()

	for i, _ := range shipments {
		shipments[i].OrderId = order.Number

		stockLocationRepository.FindByShipmentAssociation(&shipments[i])
		shipments[i].StockLocationName = shipments[i].StockLocation.Name

		shipments[i].ShippingRates = shippingRateRepository.AllByShipment(&shipments[i])

		inventoryUnitRepository.AllByShipmentAssociation(&shipments[i])
		for j := 0; j < len(shipments[i].Manifest); j++ {
			shipments[i].Manifest[j].Quantity = lineItemsMap[shipments[i].Manifest[j].LineItemId].(domain.LineItem).Quantity
		}

		shipments[i].Adjustments = adjustmentRepository.AllByAdjustable(shipments[i].Id, shipments[i].SpreeClass())
	}
	return shipments
}
