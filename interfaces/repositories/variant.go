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

func (this *VariantRepo) FindByProductId(productId int64) ([]models.Variant, error) {
	variant := &models.Variant{
		ProductId: productId,
	}

	rows, err := this.dbHandler.Find(variant).Rows()
	if err != nil {
		return nil, err
	}

	result, err := ParseAllRows(&models.Variant{}, rows)
	if err != nil {
		return nil, err
	}

	variantSlice := this.toVariantSlice(result)

	return variantSlice, nil
}

func (this *VariantRepo) toVariantSlice(result []interface{}) []models.Variant {
	variantSlice := []models.Variant{}

	for _, element := range result {
		variant := element.(models.Variant)

		variantSlice = append(variantSlice, variant)
	}

	return variantSlice
}
