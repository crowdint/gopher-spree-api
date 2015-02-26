package repositories

import "github.com/crowdint/gopher-spree-api/domain"

type ProductPropertyRepository DbRepository

func NewProductPropertyRepo() *ProductPropertyRepository {
	return &ProductPropertyRepository{
		dbHandler: Spree_db,
	}
}

func (this *ProductPropertyRepository) FindByProductIds(productIds []int64) ([]*domain.ProductProperty, error) {
	var productProperties []*domain.ProductProperty

	if len(productIds) == 0 {
		return productProperties, nil
	}

	query := this.dbHandler.
		Table("spree_product_properties").
		Select("spree_product_properties.*, spree_properties.name AS property_name").
		Joins("INNER JOIN spree_properties on spree_product_properties.property_id = spree_properties.id").
		Where("spree_product_properties.product_id in (?)", productIds).
		Scan(&productProperties)

	return productProperties, query.Error
}
