package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type OptionValueRepository struct {
	DbRepository
}

func NewOptionValueRepo() *OptionValueRepository {
	return &OptionValueRepository{
		DbRepository{dbHandler: Spree_db},
	}
}

func (this *OptionValueRepository) AllByVariantAssociation(variant *json.Variant) []models.OptionValue {
	optionValues := []models.OptionValue{}
	this.Association(variant, &optionValues, "OptionValues")
	return optionValues
}

func (this *OptionValueRepository) FindByVariantIds(variantIds []int64) ([]*models.OptionValue, error) {
	var optionValues []*models.OptionValue

	if len(variantIds) == 0 {
		return optionValues, nil
	}

	selectString := "spree_option_values.id, " +
		"spree_option_values.name, " +
		"spree_option_values.presentation, " +
		"spree_option_types.name AS option_type_name, " +
		"spree_option_types.presentation AS option_type_presentation, " +
		"spree_option_types.id AS option_type_id, " +
		"spree_option_values_variants.variant_id "

	optionTypesJoin := "INNER JOIN spree_option_types " +
		"ON spree_option_types.id = spree_option_values.option_type_id "

	valuesVariantsJoin := "INNER JOIN spree_option_values_variants " +
		"ON spree_option_values.id = spree_option_values_variants.option_value_id"

	joinString := optionTypesJoin + valuesVariantsJoin

	query := this.dbHandler.
		Table("spree_option_values").
		Select(selectString).
		Joins(joinString).
		Where("spree_option_values_variants.variant_id in (?)", variantIds).
		Scan(&optionValues)

	return optionValues, query.Error
}
