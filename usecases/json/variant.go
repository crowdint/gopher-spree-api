package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
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

type VariantsMap map[int64][]*domain.Variant

func (this *VariantInteractor) GetVariantsMap(productIds []int64) (VariantsMap, error) {
	variants, err := this.Repository.FindByProductIds(productIds)
	if err != nil {
		return VariantsMap{}, err
	}

	variantsMap, err := this.modelsToVariantsMap(variants)
	if err != nil {
		return variantsMap, err
	}

	return variantsMap, nil
}

func (this *VariantInteractor) modelsToVariantsMap(variantSlice []*domain.Variant) (VariantsMap, error) {
	variantIds := this.getIdSlice(variantSlice)
	jsonAssetsMap, err := this.AssetInteractor.GetJsonAssetsMap(variantIds)
	if err != nil {
		return VariantsMap{}, err
	}

	jsonOptionValuesMap, err := this.OptionValueInteractor.GetJsonOptionValuesMap(variantIds)
	if err != nil {
		return VariantsMap{}, err
	}

	variantsMap := VariantsMap{}

	for _, variant := range variantSlice {
		variant.Images = jsonAssetsMap[variant.Id]
		variant.OptionValues = jsonOptionValuesMap[variant.Id]

		if _, exists := variantsMap[variant.ProductId]; !exists {
			variantsMap[variant.ProductId] = []*domain.Variant{}
		}

		variantsMap[variant.ProductId] = append(variantsMap[variant.ProductId], variant)
	}

	return variantsMap, nil
}

func (this *VariantInteractor) getIdSlice(variantSlice []*domain.Variant) []int64 {
	variantIds := []int64{}

	for _, variant := range variantSlice {
		variantIds = append(variantIds, variant.Id)
	}

	return variantIds
}
