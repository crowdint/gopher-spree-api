package repositories

import "github.com/crowdint/gopher-spree-api/domain"

type ShippingMethodRepository struct {
	DbRepository
}

func NewShippingMethodRepository() *ShippingMethodRepository {
	return &ShippingMethodRepository{
		DbRepository{Spree_db},
	}
}

func (this *ShippingMethodRepository) FindByShippingRateAssociation(shippingRate *domain.ShippingRate) {
	this.Association(shippingRate, &shippingRate.ShippingMethod, "ShippingMethodId")

	shippingZoneRepository := NewShippingZoneRepository()
	shippingZoneRepository.AllByShippingMethodAssociation(&shippingRate.ShippingMethod)

	shippingCategoriesRepository := NewShippingCategoryRepository()
	shippingCategoriesRepository.AllByShippingMethodAssociation(&shippingRate.ShippingMethod)
}
