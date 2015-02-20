package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type ProductPropertyInteractor struct {
	Repo *repositories.ProductPropertyRepository
}

func NewProductPropertyInteractor() *ProductPropertyInteractor {
	return &ProductPropertyInteractor{
		Repo: repositories.NewProductPropertyRepo(),
	}
}

type JsonProductPropertiesMap map[int64][]*json.ProductProperty

func (this *ProductPropertyInteractor) GetJsonProductPropertiesMap(productIds []int64) (JsonProductPropertiesMap, error) {

	productProperties, err := this.Repo.FindByProductIds(productIds)
	if err != nil {
		return JsonProductPropertiesMap{}, err
	}

	productPropertiesJson := this.modelsToJsonProductPropertiesMap(productProperties)

	return productPropertiesJson, nil
}

func (this *ProductPropertyInteractor) modelsToJsonProductPropertiesMap(productPropertySlice []*models.ProductProperty) JsonProductPropertiesMap {
	jsonProductPropertiesMap := JsonProductPropertiesMap{}

	for _, productProperty := range productPropertySlice {
		productPropertyJson := this.toJson(productProperty)

		if _, exists := jsonProductPropertiesMap[productProperty.ProductId]; !exists {
			jsonProductPropertiesMap[productProperty.ProductId] = []*json.ProductProperty{}
		}

		jsonProductPropertiesMap[productProperty.ProductId] = append(jsonProductPropertiesMap[productProperty.ProductId], productPropertyJson)

	}

	return jsonProductPropertiesMap
}

func (this *ProductPropertyInteractor) toJson(productProperty *models.ProductProperty) *json.ProductProperty {
	productPropertyJson := &json.ProductProperty{
		Id:           productProperty.Id,
		ProductId:    productProperty.ProductId,
		PropertyId:   productProperty.PropertyId,
		Value:        productProperty.Value,
		PropertyName: productProperty.PropertyName,
	}
	return productPropertyJson
}
