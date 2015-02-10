package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type ProductResponse struct {
	data []*json.Product
}

func (this ProductResponse) GetCount() int {
	return len(this.data)
}

func (this ProductResponse) GetData() interface{} {
	return this.data
}

func (this ProductResponse) GetTag() string {
	return "products"
}

type ProductInteractor struct {
	ProductRepo               *repositories.ProductRepo
	VariantInteractor         *VariantInteractor
	ProductPropertyInteractor *ProductPropertyInteractor
}

func NewProductInteractor() *ProductInteractor {
	return &ProductInteractor{
		ProductRepo:               repositories.NewProductRepo(),
		VariantInteractor:         NewVariantInteractor(),
		ProductPropertyInteractor: NewProductPropertyInteractor(),
	}
}

func (this *ProductInteractor) GetResponse(currentPage, perPage int, query string) (ContentResponse, error) {
	productModelSlice, err := this.ProductRepo.List(currentPage, perPage, query)
	if err != nil {
		return ProductResponse{}, err
	}

	productJsonSlice := this.modelsToJsonProductsSlice(productModelSlice)

	productIds := this.getIdSlice(productModelSlice)

	variantsMap, err := this.VariantInteractor.GetJsonVariantsMap(productIds)
	if err != nil {
		return ProductResponse{}, err
	}

	productPropertiesMap, err := this.ProductPropertyInteractor.GetJsonProductPropertiesMap(productIds)
	if err != nil {
		return ProductResponse{}, err
	}

	this.mergeVariants(productJsonSlice, variantsMap)

	this.mergeProductProperties(productJsonSlice, productPropertiesMap)

	return ProductResponse{
		data: productJsonSlice,
	}, nil
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

func (this *ProductInteractor) mergeProductProperties(productSlice []*json.Product, productPropertiesMap JsonProductPropertiesMap) {
	for _, product := range productSlice {
		product.ProductProperties = []json.ProductProperty{}

		productPropertiesSlice := productPropertiesMap[product.ID]

		if productPropertiesSlice == nil {
			continue
		}

		for _, productProperty := range productPropertiesSlice {
			product.ProductProperties = append(product.ProductProperties, *productProperty)
		}
	}
}

func (this *ProductInteractor) GetTotalCount() (int64, error) {
	return this.ProductRepo.CountAll()
}
