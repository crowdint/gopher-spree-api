package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type PriceRepository DbRepository

func NewPriceRepo() *PriceRepository {
	return &PriceRepository{
		dbHandler: Spree_db,
	}
}

func (this *PriceRepository) GetByVariant(variantId int64) models.Price {
	var price models.Price

	this.dbHandler.Where("variant_id = ? AND currency = ?", variantId, "USD").First(&price)

	return price
}
