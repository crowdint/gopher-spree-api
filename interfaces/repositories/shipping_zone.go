package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
)

type ShippingZoneRepository struct {
	DbRepository
}

func NewShippingZoneRepository() *ShippingZoneRepository {
	return &ShippingZoneRepository{
		DbRepository{Spree_db},
	}
}

func (this *ShippingZoneRepository) AllByShippingMethodAssociation(shippingMethod *domain.ShippingMethod) {
	this.dbHandler.Table("spree_zones").
		Select("spree_zones.*").
		Where("spree_shipping_methods_zones.shipping_method_id = ?", shippingMethod.Id).
		Joins(`INNER JOIN "spree_shipping_methods_zones" ON "spree_zones"."id" = "spree_shipping_methods_zones"."zone_id"`).
		Find(&shippingMethod.Zones)
}
