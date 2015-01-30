package json

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

func (this *ProductInteractor) GetResponse(currentPage, perPage int) ([]*json.Product, error) {
	productModelSlice, err := this.ProductRepo.List(currentPage, perPage)
	if err != nil {
		return []*json.Product{}, err
	}

	productJsonSlice := this.modelsToJsonProductsSlice(productModelSlice)

	productIds := this.getIdSlice(productModelSlice)

	variantsMap, err := this.VariantInteractor.GetJsonVariantsMap(productIds)
	if err != nil {
		return []*json.Product{}, err
	}

	this.mergeVariants(productJsonSlice, variantsMap)

	return productJsonSlice, nil
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
		//Price: from master variant
		//DisplayPrice:
		AvailableOn:     product.AvailableOn,
		Slug:            product.Slug,
		MetaDescription: product.MetaDescription,
		MetaKeyWords:    product.MetaDescription,
		//ShippingCategoryId
		//TaxonIds
		//TotalOnHand: from variants
		//HasVariants: form variants
		//Master: master variant
		//Variants: from JsonVariantsMap
		//OptionTypes
		//ProductProperties
		//Classifications
	}

	return productJson
}

func (this *ProductInteractor) mergeVariants(productSlice []*json.Product, variantsMap JsonVariantsMap) {
	for _, product := range productSlice {
		product.Variants = []json.Variant{}
		var totalOnHand int64

		variantSlice := variantsMap[product.ID]

		if variantSlice == nil {
			continue
		}

		for _, variant := range variantSlice {
			if variant.IsMaster {
				product.Master = *variant
				product.Price = variant.Price
			} else {
				product.Variants = append(product.Variants, *variant)
			}

			totalOnHand += variant.TotalOnHand

			variant.Description = product.Description
			variant.Slug = product.Slug
			variant.Name = product.Name
		}

		product.TotalOnHand = totalOnHand

		if len(product.Variants) > 0 {
			product.HasVariants = true
		}
	}
}

func (this *ProductInteractor) GetTotalCount() (int64, error) {
	return this.ProductRepo.CountAll()
}
