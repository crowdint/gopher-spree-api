package repositories

import "github.com/crowdint/gopher-spree-api/domain/json"

type ShippingCategoryRepository struct {
	DbRepository
}

func NewShippingCategoryRepository() *ShippingCategoryRepository {
	return &ShippingCategoryRepository{
		DbRepository{Spree_db},
	}
}

func (this *ShippingCategoryRepository) AllByShippingMethodAssociation(shippingMethod *json.ShippingMethod) {
	this.dbHandler.Table("spree_shipping_categories").
		Select("spree_shipping_categories.*").
		Where("spree_shipping_method_categories.shipping_method_id = ?", shippingMethod.Id).
		Joins(`INNER JOIN "spree_shipping_method_categories" ON "spree_shipping_categories"."id" = "spree_shipping_method_categories"."shipping_category_id"`).
		Find(&shippingMethod.ShippingCategories)
}
