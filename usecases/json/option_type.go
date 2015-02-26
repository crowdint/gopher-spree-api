package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type OptionTypeInteractor struct {
	Repository *repositories.OptionTypeRepository
}

func NewOptionTypeInteractor() *OptionTypeInteractor {
	return &OptionTypeInteractor{
		Repository: repositories.NewOptionTypeRepo(),
	}
}

type JsonOptionTypesMap map[int64][]*domain.OptionType

func (this *OptionTypeInteractor) GetJsonOptionTypesMap(productIds []int64) (JsonOptionTypesMap, error) {

	optionTypes, err := this.Repository.FindByProductIds(productIds)
	if err != nil {
		return JsonOptionTypesMap{}, err
	}

	optionTypesJson := this.modelsToJsonOptionTypesMap(optionTypes)

	return optionTypesJson, nil
}

func (this *OptionTypeInteractor) modelsToJsonOptionTypesMap(optionTypeSlice []*domain.OptionType) JsonOptionTypesMap {
	jsonOptionTypesMap := JsonOptionTypesMap{}

	for _, optionType := range optionTypeSlice {
		optionTypeJson := this.toJson(optionType)

		if _, exists := jsonOptionTypesMap[optionType.ProductId]; !exists {
			jsonOptionTypesMap[optionType.ProductId] = []*domain.OptionType{}
		}

		jsonOptionTypesMap[optionType.ProductId] = append(jsonOptionTypesMap[optionType.ProductId], optionTypeJson)

	}

	return jsonOptionTypesMap
}

func (this *OptionTypeInteractor) toJson(optionType *domain.OptionType) *domain.OptionType {
	optionTypeJson := &domain.OptionType{
		Id:           optionType.Id,
		Name:         optionType.Name,
		Presentation: optionType.Presentation,
		Position:     optionType.Position,
	}
	return optionTypeJson
}
