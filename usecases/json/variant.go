package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type VariantInteractor struct {
	Repo *repositories.VariantRepo
}

func NewVariantInteractor() *VariantInteractor {
	return &VariantInteractor{
		Repo: repositories.NewVariantRepo(),
	}
}

type JsonVariantsMap map[int64][]*json.Variant

func (this *VariantInteractor) GetJsonVariantsMap(productIds []int64) JsonVariantsMap {

	variants := this.Repo.FindByProductIds(productIds)

	variantsJson := this.modelsToJsonVariantsMap(variants)

	return variantsJson
}

func (this *VariantInteractor) modelsToJsonVariantsMap(variantSlice []*models.Variant) JsonVariantsMap {
	jsonVariantsMap := JsonVariantsMap{}

	for _, variant := range variantSlice {
		variantJson := this.toJson(variant)

		if _, exists := jsonVariantsMap[variant.ProductId]; !exists {
			jsonVariantsMap[variant.ProductId] = []*json.Variant{}
		}

		jsonVariantsMap[variant.ProductId] = append(jsonVariantsMap[variant.ProductId], variantJson)

	}

	return jsonVariantsMap
}

func (this *VariantInteractor) toJson(variant *models.Variant) *json.Variant {
	variantJson := &json.Variant{
		ID: variant.Id,
		//Name:
		Sku:      variant.Sku,
		Price:    variant.Price,
		Weight:   variant.Weight,
		Height:   variant.Height,
		Width:    variant.Width,
		Depth:    variant.Depth,
		IsMaster: variant.IsMaster,
		//Slug:
		//Description
		TrackInventory: variant.TrackInventory,
		CostPrice:      variant.CostPrice,
		//DisplayPrice:
		//OptionsText:
		//InStock:
		//IsBackorderable:
		//TotalOnHand:
		//IsDestroyed:
		//OptionValues:
		//Images:
	}
	return variantJson
}
