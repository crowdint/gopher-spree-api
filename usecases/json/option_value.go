package json

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type OptionValueInteractor struct {
	Repository *repositories.OptionValueRepository
}

func NewOptionValueInteractor() *OptionValueInteractor {
	return &OptionValueInteractor{
		Repository: repositories.NewOptionValueRepo(),
	}
}

type OptionValuesMap map[int64][]models.OptionValue

func (this *OptionValueInteractor) GetJsonOptionValuesMap(variantIds []int64) (OptionValuesMap, error) {

	optionValues, err := this.Repository.FindByVariantIds(variantIds)
	if err != nil {
		return OptionValuesMap{}, err
	}

	optionValuesJson := this.modelsToJsonOptionValuesMap(optionValues)

	return optionValuesJson, nil
}

func (this *OptionValueInteractor) modelsToJsonOptionValuesMap(optionValueSlice []*models.OptionValue) OptionValuesMap {
	jsonOptionValuesMap := OptionValuesMap{}

	for _, optionValue := range optionValueSlice {
		if _, exists := jsonOptionValuesMap[optionValue.VariantId]; !exists {
			jsonOptionValuesMap[optionValue.VariantId] = []models.OptionValue{}
		}

		jsonOptionValuesMap[optionValue.VariantId] =
			append(jsonOptionValuesMap[optionValue.VariantId], *optionValue)

	}

	return jsonOptionValuesMap
}
