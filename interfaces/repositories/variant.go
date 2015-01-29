package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type VariantRepo DbRepo

func NewVariantRepo() *VariantRepo {
	return &VariantRepo{
		dbHandler: Spree_db,
	}
}

func (this *VariantRepo) FindByProductIds(productIds []int64) []*models.Variant {

	var variants []*models.Variant

	this.dbHandler.Where("product_id in (?)", productIds).Find(&variants)

	return variants
}
