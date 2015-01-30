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

func (this *VariantRepo) FindByProductIds(productIds []int64) ([]*models.Variant, error) {

	variants := []*models.Variant{}

	if len(productIds) == 0 {
		return variants, nil
	}

	sqlStr := "SELECT spree_variants.*, SUM(count_on_hand) AS real_stock_items_count, spree_stock_items.backorderable, spree_prices.amount price " +
		"FROM spree_variants " +
		"INNER JOIN spree_stock_items ON spree_variants.id = spree_stock_items.variant_id " +
		"INNER JOIN spree_prices ON spree_variants.id = spree_prices.variant_id " +
		"WHERE spree_prices.currency='USD' " +
		"AND spree_variants.product_id IN (?) " +
		"GROUP BY spree_variants.id, spree_variants, backorderable, price"

	query := this.dbHandler.Raw(sqlStr, productIds).Scan(&variants)

	return variants, query.Error
}
