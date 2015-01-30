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

	sqlStr := "select spree_variants.*, sum(count_on_hand) AS real_stock_items_count, spree_stock_items.backorderable, spree_prices.amount price " +
		"from spree_variants " +
		"inner join spree_stock_items on spree_variants.id = spree_stock_items.variant_id " +
		"inner join spree_prices on spree_variants.id = spree_prices.variant_id " +
		"where spree_prices.currency='USD' " +
		"AND spree_variants.product_id IN (?) " +
		"group by spree_variants.id, spree_variants, backorderable, price"

	query := this.dbHandler.Raw(sqlStr, productIds).Scan(&variants)

	return variants, query.Error
}
