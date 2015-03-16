package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type ProductPropertyInteractor struct {
	Repository *repositories.ProductPropertyRepository
}

func NewProductPropertyInteractor() *ProductPropertyInteractor {
	return &ProductPropertyInteractor{
		Repository: repositories.NewProductPropertyRepo(),
	}
}

type JsonProductPropertiesMap map[int64][]*domain.ProductProperty

func (this *ProductPropertyInteractor) GetJsonProductPropertiesMap(productIds []int64) (JsonProductPropertiesMap, error) {

	productProperties, err := this.Repository.FindByProductIds(productIds)
	if err != nil {
		utils.LogrusError("GetJsonProductPropertiesMap", "", err)

		return JsonProductPropertiesMap{}, err
	}

	productPropertiesJson := this.modelsToJsonProductPropertiesMap(productProperties)

	return productPropertiesJson, nil
}

func (this *ProductPropertyInteractor) modelsToJsonProductPropertiesMap(productPropertySlice []*domain.ProductProperty) JsonProductPropertiesMap {
	jsonProductPropertiesMap := JsonProductPropertiesMap{}

	for _, productProperty := range productPropertySlice {
		if _, exists := jsonProductPropertiesMap[productProperty.ProductId]; !exists {
			jsonProductPropertiesMap[productProperty.ProductId] = []*domain.ProductProperty{}
		}

		jsonProductPropertiesMap[productProperty.ProductId] = append(jsonProductPropertiesMap[productProperty.ProductId], productProperty)

	}

	return jsonProductPropertiesMap
}
