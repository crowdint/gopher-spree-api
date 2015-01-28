package usecases

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type ProductInteractor struct {
	ProductRepo       *repositories.ProductRepo
	VariantInteractor *VariantInteractor
}

func NewProductInteractor() *ProductInteractor {
	return &ProductInteractor{
		ProductRepo:       repositories.NewProductRepo(),
		VariantInteractor: NewVariantInteractor(),
	}
}

func (this *ProductInteractor) GetMergedResponse() []*json.Product {
	productModelSlice := this.ProductRepo.List()

	productJsonSlice := this.modelsToJsonProductsSlice(productModelSlice)

	productIds := this.getIdSlice(productModelSlice)

	variantsMap := this.VariantInteractor.GetJsonVariantsMap(productIds)

	this.mergeVariants(productJsonSlice, variantsMap)

	return productJsonSlice
}

func (this *ProductInteractor) getIdSlice(productSlice []*models.Product) []int64 {
	productIds := []int64{}

	for _, product := range productSlice {
		productIds = append(productIds, product.Id)
	}

	return productIds
}

func (this *ProductInteractor) modelsToJsonProductsSlice(productSlice []*models.Product) []*json.Product {
	jsonProductsSlice := []*json.Product{}

	for _, product := range productSlice {
		productJson := this.toJson(product)

		jsonProductsSlice = append(jsonProductsSlice, productJson)
	}

	return jsonProductsSlice
}

func (this *ProductInteractor) toJson(product *models.Product) *json.Product {
	productJson := &json.Product{
		ID:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		//Price:
		//DisplayPrice:
		AvailableOn:     product.AvailableOn,
		Slug:            product.Slug,
		MetaDescription: product.MetaDescription,
		MetaKeyWords:    product.MetaDescription,
		//ShippingCategoryId
		//TaxonIds
		//TotalOnHand
		//HasVariants
		//Master
		//Variants
		//OptionTypes
		//ProductProperties
		//Classifications
	}

	return productJson
}

func (this *ProductInteractor) mergeVariants(productSlice []*json.Product, variantsMap JsonVariantsMap) {
	for _, product := range productSlice {
		product.Variants = []json.Variant{}

		variantSlice := variantsMap[product.ID]

		if variantSlice == nil {
			continue
		}

		for _, variant := range variantSlice {
			if variant.IsMaster {
				product.Master = *variant
			}
			product.Variants = append(product.Variants, *variant)
		}
	}
}
