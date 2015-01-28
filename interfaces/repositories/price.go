package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type PriceRepo DbRepo

func NewPriceRepo() *PriceRepo {
	return &PriceRepo{
		dbHandler: Spree_db,
	}
}

func (this *PriceRepo) GetByVariant(variantId int64) models.Price {
	var price models.Price

	this.dbHandler.Where("variant_id = ? AND currency = ?", variantId, "USD").First(&price)

	return price
}
