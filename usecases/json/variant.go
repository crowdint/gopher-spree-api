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

func (this *VariantInteractor) GetJsonVariantsMap(productIds []int64) (JsonVariantsMap, error) {

	variants, err := this.Repo.FindByProductIds(productIds)
	if err != nil {
		return JsonVariantsMap{}, err
	}

	variantsJson := this.modelsToJsonVariantsMap(variants)

	return variantsJson, nil
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
		//Name: from product
		Sku:      variant.Sku,
		Price:    variant.Price,
		Weight:   json.ToS(variant.Weight),
		Height:   json.ToS(variant.Height),
		Width:    json.ToS(variant.Width),
		Depth:    json.ToS(variant.Depth),
		IsMaster: variant.IsMaster,
		//Slug: from product
		//Description: from product
		TrackInventory: variant.TrackInventory,
		CostPrice:      variant.CostPrice,
		//DisplayPrice:
		//OptionsText:
		InStock:         variant.RealStockItemsCount > 0,
		IsBackorderable: variant.Backorderable,
		TotalOnHand:     variant.RealStockItemsCount,
		IsDestroyed:     !variant.DeletedAt.IsZero(),
		//OptionValues:
		//Images:
	}
	return variantJson
}
