package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type OptionValueInteractor struct {
	Repository *repositories.OptionValueRepository
}

func NewOptionValueInteractor() *OptionValueInteractor {
	return &OptionValueInteractor{
		Repository: repositories.NewOptionValueRepo(),
	}
}

type OptionValuesMap map[int64][]domain.OptionValue

func (this *OptionValueInteractor) GetJsonOptionValuesMap(variantIds []int64) (OptionValuesMap, error) {

	optionValues, err := this.Repository.FindByVariantIds(variantIds)
	if err != nil {
		utils.LogrusError("GetJsonOptionValuesMap", "GET", err)

		return OptionValuesMap{}, err
	}

	optionValuesJson := this.modelsToJsonOptionValuesMap(optionValues)

	return optionValuesJson, nil
}

func (this *OptionValueInteractor) modelsToJsonOptionValuesMap(optionValueSlice []*domain.OptionValue) OptionValuesMap {
	jsonOptionValuesMap := OptionValuesMap{}

	for _, optionValue := range optionValueSlice {
		if _, exists := jsonOptionValuesMap[optionValue.VariantId]; !exists {
			jsonOptionValuesMap[optionValue.VariantId] = []domain.OptionValue{}
		}

		jsonOptionValuesMap[optionValue.VariantId] =
			append(jsonOptionValuesMap[optionValue.VariantId], *optionValue)

	}

	return jsonOptionValuesMap
}
