package domain

type InventoryUnit struct {
	Id         int64            `json:"-"`
	LineItemId int64            `json:"-"`
	Pending    bool             `json:"-"`
	Quantity   int64            `json:"quantity"`
	ShipmentId int64            `json:"-"`
	State      string           `json:"-"`
	States     map[string]int64 `json:"states"`
	VariantId  int64            `json:"variant_id"`
	OrderId    int64            `json:"-"`
}

func (this InventoryUnit) TableName() string {
	return "spree_inventory_units"
}
