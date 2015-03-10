package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
)

type StockItemRepository struct {
	DbRepository
}

func NewStockItemRepository() *StockItemRepository {
	return &StockItemRepository{
		DbRepository{Spree_db},
	}
}

func (this *StockItemRepository) Create(stockItem *domain.StockItem) error {
	return this.DbRepository.Create(stockItem)
}
