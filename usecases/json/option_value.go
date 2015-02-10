package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/jinzhu/copier"
)

type OptionValueInteractor struct {
	Repo *repositories.OptionValueRepo
}

func NewOptionValueInteractor() *OptionValueInteractor {
	return &OptionValueInteractor{
		Repo: repositories.NewOptionValueRepo(),
	}
}

type JsonOptionValuesMap map[int64][]*json.OptionValue

func (this *OptionValueInteractor) GetJsonOptionValuesMap(variantIds []int64) (JsonOptionValuesMap, error) {

	optionValues, err := this.Repo.FindByVariantIds(variantIds)
	if err != nil {
		return JsonOptionValuesMap{}, err
	}

	optionValuesJson := this.modelsToJsonOptionValuesMap(optionValues)

	return optionValuesJson, nil
}

func (this *OptionValueInteractor) modelsToJsonOptionValuesMap(optionValueSlice []*models.OptionValue) JsonOptionValuesMap {
	jsonOptionValuesMap := JsonOptionValuesMap{}

	for _, optionValue := range optionValueSlice {
		optionValueJson := &json.OptionValue{}
		copier.Copy(optionValueJson, optionValue)

		if _, exists := jsonOptionValuesMap[optionValue.VariantId]; !exists {
			jsonOptionValuesMap[optionValue.VariantId] = []*json.OptionValue{}
		}

		jsonOptionValuesMap[optionValue.VariantId] =
			append(jsonOptionValuesMap[optionValue.VariantId], optionValueJson)

	}

	return jsonOptionValuesMap
}
