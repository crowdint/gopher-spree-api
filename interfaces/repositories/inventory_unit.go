package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/utils"
)

type InventoryUnitRepository struct {
	DbRepository
}

func NewInventoryUnitRepository() *InventoryUnitRepository {
	return &InventoryUnitRepository{
		DbRepository{Spree_db},
	}
}

func (this *InventoryUnitRepository) AllByShipmentAssociation(shipment *json.Shipment) {
	this.Association(shipment, &shipment.Manifest, "ShipmentId")

	manifestToMap := utils.ToMap(shipment.Manifest, "LineItemId", true)

	for i := 0; i < len(shipment.Manifest); i++ {
		states := make(map[string]int64)

		for _, v := range manifestToMap {
			inventoryUnits := v.([]interface{})
			for j := 0; j < len(inventoryUnits); j++ {
				states[inventoryUnits[j].(json.InventoryUnit).State]++
			}
		}

		shipment.Manifest[i].States = states
	}
}
