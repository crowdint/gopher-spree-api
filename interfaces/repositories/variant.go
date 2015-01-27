package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type VariantRepo DbRepo

func NewVariantRepo() *VariantRepo {
	return &VariantRepo{
		dbHandler: spree_db,
	}
}

func (this *VariantRepo) FindByProductId(productId int64) *models.Variant {
	variant := &models.Variant{
		ProductId: productId,
	}

	this.dbHandler.Find(variant)

	return variant
}
