package repositories

import "github.com/crowdint/gopher-spree-api/domain/json"

type OptionTypeRepository DbRepository

func NewOptionTypeRepo() *OptionTypeRepository {
	return &OptionTypeRepository{
		dbHandler: Spree_db,
	}
}

func (this *OptionTypeRepository) FindByProductIds(productIds []int64) ([]*json.OptionType, error) {
	var optionTypes []*json.OptionType

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
