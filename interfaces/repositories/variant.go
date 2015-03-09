package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/jinzhu/gorm"
)

type VariantRepository struct {
	DbRepository
}

func NewVariantRepository() *VariantRepository {
	return &VariantRepository{
		DbRepository{Spree_db},
	}
}

func (this *VariantRepository) FindByProductIds(productIds []int64) ([]*domain.Variant, error) {

	variants := []*domain.Variant{}

	if len(productIds) == 0 {
		return variants, nil
	}

	query := this.dbHandler.
		Table("spree_variants").
		Select("spree_variants.*, SUM(count_on_hand) AS real_stock_items_count, spree_stock_items.backorderable, spree_prices.amount price").
		Joins("INNER JOIN spree_stock_items ON spree_variants.id = spree_stock_items.variant_id INNER JOIN spree_prices ON spree_variants.id = spree_prices.variant_id").
		Where("spree_prices.currency='USD'").
		Where("spree_variants.product_id IN (?)", productIds).
		Group("spree_variants.id, spree_variants, backorderable, price").
		Scan(&variants)

	//	spew.Dump("Variants from the query %-v", variants)
	if query.Error != nil && query.Error != gorm.RecordNotFound {
		return variants, nil
	} else {
		return variants, query.Error
	}
}

func (this *VariantRepository) Create(variant *domain.Variant) error {
	if err := this.DbRepository.Create(variant); err != nil {
		return err
	}

	return this.AfterCreate(variant)
}

func (this *VariantRepository) AfterCreate(variant *domain.Variant) error {
	// TODO: setPosition and setMasterOutOfStock
	return this.createStockItems(variant)
}

func (this *VariantRepository) createStockItems(variant *domain.Variant) error {
	stockLocationRepository := NewStockLocationRepository()
	stockItemRepository := NewStockItemRepository()
	stockLocations, err := stockLocationRepository.AllBy("propagate_all_variants = ?", true)
	if err != nil {
		return err
	}

	for _, stockLocation := range stockLocations {
		stockItem := &domain.StockItem{VariantId: variant.Id, StockLocationId: stockLocation.Id, Backorderable: stockLocation.BackorderableDefault}
		if err = stockItemRepository.Create(stockItem); err != nil {
			return err
		}
		stockItem.StockLocation = stockLocation
		stockItem.Variant = variant
		variant.StockItems = append(variant.StockItems, stockItem)
	}

	return nil
}
