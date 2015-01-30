package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type ProductPropertyRepo DbRepo

func NewProductPropertyRepo() *ProductPropertyRepo {
	return &ProductPropertyRepo{
		dbHandler: Spree_db,
	}
}

func (this *ProductPropertyRepo) FindByProductIds(productIds []int64) ([]*models.ProductProperty, error) {
	var productProperties []*models.ProductProperty

	query := this.dbHandler.
		Table("spree_product_properties").
		Select("spree_product_properties.*, spree_properties.name AS property_name").
		Joins("INNER JOIN spree_properties on spree_product_properties.property_id = spree_properties.id").
		Where("spree_product_properties.product_id in (?)", productIds).
		Scan(&productProperties)

	return productProperties, query.Error
}
