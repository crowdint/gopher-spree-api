package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type OptionTypeInteractor struct {
	Repo *repositories.OptionTypeRepo
}

func NewOptionTypeInteractor() *OptionTypeInteractor {
	return &OptionTypeInteractor{
		Repo: repositories.NewOptionTypeRepo(),
	}
}

type JsonOptionTypesMap map[int64][]*json.OptionType

func (this *OptionTypeInteractor) GetJsonOptionTypesMap(productIds []int64) (JsonOptionTypesMap, error) {

	optionTypes, err := this.Repo.FindByProductIds(productIds)
	if err != nil {
		return JsonOptionTypesMap{}, err
	}

	optionTypesJson := this.modelsToJsonOptionTypesMap(optionTypes)

	return optionTypesJson, nil
}

func (this *OptionTypeInteractor) modelsToJsonOptionTypesMap(optionTypeSlice []*models.OptionType) JsonOptionTypesMap {
	jsonOptionTypesMap := JsonOptionTypesMap{}

	for _, optionType := range optionTypeSlice {
		optionTypeJson := this.toJson(optionType)

		if _, exists := jsonOptionTypesMap[optionType.ProductId]; !exists {
			jsonOptionTypesMap[optionType.ProductId] = []*json.OptionType{}
		}

		jsonOptionTypesMap[optionType.ProductId] = append(jsonOptionTypesMap[optionType.ProductId], optionTypeJson)

	}

	return jsonOptionTypesMap
}

func (this *OptionTypeInteractor) toJson(optionType *models.OptionType) *json.OptionType {
	optionTypeJson := &json.OptionType{
		ID:           optionType.Id,
		Name:         optionType.Name,
		Presentation: optionType.Presentation,
		Position:     optionType.Position,
	}
	return optionTypeJson
}
