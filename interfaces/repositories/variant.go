package repositories

import (
	"github.com/gopher-spree-api/domain/models"
)

type VariantRepo DbRepo

func NewVariantRepo() *VariantRepo {
	return &VariantRepo{
		dbHandler: spree_db,
	}
}

func (this *VariantRepo) FindByProductId(productId int64) ([]*models.Variant, error) {
	variant := &models.Variant{
		ProductId: productId,
	}

	variantSlice := []*models.Variant{}

	rows, err := this.dbHandler.Find(variant).Rows()
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice

	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		rows.Scan(dest...)

		newVariant := &models.Variant{}

		ParseRow(rawResult, newVariant)

		variantSlice = append(variantSlice, newVariant)
	}

	return variantSlice, nil
}
