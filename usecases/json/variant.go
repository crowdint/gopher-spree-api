package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type VariantInteractor struct {
	Repository            *repositories.VariantRepository
	AssetInteractor       *AssetInteractor
	OptionValueInteractor *OptionValueInteractor
}

func NewVariantInteractor() *VariantInteractor {
	return &VariantInteractor{
		Repository:            repositories.NewVariantRepository(),
		AssetInteractor:       NewAssetInteractor(),
		OptionValueInteractor: NewOptionValueInteractor(),
	}
}

type JsonVariantsMap map[int64][]*domain.Variant

func (this *VariantInteractor) GetJsonVariantsMap(productIds []int64) (JsonVariantsMap, error) {
	variants, err := this.Repository.FindByProductIds(productIds)
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return JsonVariantsMap{}, err
	}

	variantsJson, err := this.modelsToJsonVariantsMap(variants)
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return variantsJson, err
	}

	return variantsJson, nil
}

func (this *VariantInteractor) modelsToJsonVariantsMap(variantSlice []*domain.Variant) (JsonVariantsMap, error) {
	variantIds := this.getIdSlice(variantSlice)
	jsonAssetsMap, err := this.AssetInteractor.GetJsonAssetsMap(variantIds)
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return JsonVariantsMap{}, err
	}

	jsonOptionValuesMap, err := this.OptionValueInteractor.GetJsonOptionValuesMap(variantIds)
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return JsonVariantsMap{}, err
	}

	jsonVariantsMap := JsonVariantsMap{}

	for _, variant := range variantSlice {
		variantJson := this.toJson(variant)

		variantJson.Images = jsonAssetsMap[variant.Id]
		variantJson.OptionValues = jsonOptionValuesMap[variant.Id]

		if _, exists := jsonVariantsMap[variant.ProductId]; !exists {
			jsonVariantsMap[variant.ProductId] = []*domain.Variant{}
		}

		jsonVariantsMap[variant.ProductId] = append(jsonVariantsMap[variant.ProductId], variantJson)

	}

	return jsonVariantsMap, nil
}

func (this *VariantInteractor) toJson(variant *domain.Variant) *domain.Variant {
	variantJson := &domain.Variant{
		Id: variant.Id,
		//Name: from product
		Sku:      variant.Sku,
		Price:    variant.Price,
		Weight:   variant.Weight,
		Height:   variant.Height,
		Width:    variant.Width,
		Depth:    variant.Depth,
		IsMaster: variant.IsMaster,
		//Slug: from product
		//Description: from product
		TrackInventory: variant.TrackInventory,
		CostPrice:      variant.CostPrice,
		//DisplayPrice:
		//OptionsText:
		InStock:         variant.RealStockItemsCount > 0,
		IsBackorderable: variant.Backorderable,
		TotalOnHand:     &variant.RealStockItemsCount,
		IsDestroyed:     !variant.DeletedAt.IsZero(),
		//OptionValues:
		//Images:
	}
	return variantJson
}

func (this *VariantInteractor) getIdSlice(variantSlice []*domain.Variant) []int64 {
	variantIds := []int64{}

	for _, variant := range variantSlice {
		variantIds = append(variantIds, variant.Id)
	}

	return variantIds
}
