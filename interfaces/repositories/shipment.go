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

func (this *ShipmentRepository) AllByOrder(order *domain.Order) []*domain.Shipment {
	shipments := []*domain.Shipment{}
	this.All(&shipments, nil, "order_id = ?", order.Id)
	lineItemsMap := utils.ToMap(*(order.LineItems), "Id", false)

	stockLocationRepository := NewStockLocationRepository()
	shippingRateRepository := NewShippingRateRepository()
	inventoryUnitRepository := NewInventoryUnitRepository()
	adjustmentRepository := NewAdjustmentRepository()

	for _, shipment := range shipments {
		shipment.OrderId = order.Number

		stockLocationRepository.FindByShipmentAssociation(shipment)
		shipment.StockLocationName = shipment.StockLocation.Name

		shipment.ShippingRates = shippingRateRepository.AllByShipment(shipment)

		inventoryUnitRepository.AllByShipmentAssociation(shipment)
		for j := 0; j < len(shipment.Manifest); j++ {
			shipment.Manifest[j].Quantity = lineItemsMap[shipment.Manifest[j].LineItemId].(domain.LineItem).Quantity
		}

		shipment.Adjustments = adjustmentRepository.AllByAdjustable(shipment)
	}
	return shipments
}
