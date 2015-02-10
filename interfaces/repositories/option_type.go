package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type OptionTypeRepo DbRepo

func NewOptionTypeRepo() *OptionTypeRepo {
	return &OptionTypeRepo{
		dbHandler: Spree_db,
	}
}

func (this *OptionTypeRepo) FindByProductIds(productIds []int64) ([]*models.OptionType, error) {
	var optionTypes []*models.OptionType
	if len(productIds) == 0 {
		return optionTypes, nil
	}
	query := this.dbHandler.
		Table("spree_option_types").
		Select("spree_option_types.id, spree_option_types.name, spree_option_types.presentation, spree_option_types.position, spree_product_option_types.product_id").
		Joins("INNER JOIN spree_product_option_types ON spree_option_types.id = spree_product_option_types.option_type_id").
		Where("spree_product_option_types.product_id IN (?)", productIds).
		Scan(&optionTypes)

	return optionTypes, query.Error
}
