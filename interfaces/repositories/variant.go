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

	sqlStr := "SELECT " +
		"* " +
		"FROM " +
		"spree_variants " +
		"INNER JOIN" +
		"(SELECT " +
		"variant_id, " +
		"SUM(count_on_hand) AS real_stock_items_count, " +
		"backorderable " +
		"FROM " +
		"spree_stock_items " +
		"GROUP BY " +
		"variant_id, backorderable) AS si " +
		"ON " +
		"si.variant_id = spree_variants.id " +
		"INNER JOIN " +
		"(SELECT " +
		"id as price_id, " +
		"amount as price " +
		"FROM " +
		"spree_prices " +
		"WHERE " +
		"currency='USD') as sp " +
		"ON " +
		"sp.price_id = spree_variants.id " +
		"AND " +
		"spree_variants.product_id IN (?)"

	query := this.dbHandler.Raw(sqlStr, productIds).Scan(&variants)

	return variants, query.Error
}
