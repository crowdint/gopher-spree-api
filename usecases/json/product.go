package json

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"

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
	ClassificationInteractor  *ClassificationInteractor
	OptionTypeInteractor      *OptionTypeInteractor
}

func NewProductInteractor() *ProductInteractor {
	return &ProductInteractor{
		ProductRepo:               repositories.NewProductRepo(),
		VariantInteractor:         NewVariantInteractor(),
		ProductPropertyInteractor: NewProductPropertyInteractor(),
		ClassificationInteractor:  NewClassificationInteractor(),
		OptionTypeInteractor:      NewOptionTypeInteractor(),
	}
}

func (this *ProductInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	query, gparams, err := params.GetGransakParams()
	if err != nil {
		return ProductResponse{}, err
	}

	productModelSlice, err := this.ProductRepo.List(currentPage, perPage, query, gparams)
	if err != nil {
		return ProductResponse{}, err
	}

	productJsonSlice, err := this.transformToJsonResponse(productModelSlice)
	if err != nil {
		return ProductResponse{}, err
	}

	return ProductResponse{
		data: productJsonSlice,
	}, nil
}

func (this *ProductInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	id, err := params.GetIntParam(ID_PARAM)

	if err != nil {
		return struct{}{}, errors.New("Invalid parameter type: " + fmt.Sprintf("%v", id))
	}

	product, err := this.ProductRepo.FindById(int64(id))

	if err != nil {
		return nil, err
	}

	productModelSlice := []*models.Product{}

	productModelSlice = append(productModelSlice, product)

	productJsonSlice, err := this.transformToJsonResponse(productModelSlice)
	if err != nil {
		return nil, err
	}

	return productJsonSlice[0], nil
}

func (this *ProductInteractor) transformToJsonResponse(productModelSlice []*models.Product) ([]*json.Product, error) {
	productJsonSlice := this.modelsToJsonProductsSlice(productModelSlice)

	productIds := this.getIdSlice(productModelSlice)

	err := this.mergeComplementaryValues(productIds, productJsonSlice)
	if err != nil {
		return []*json.Product{}, err
	}

	return productJsonSlice, nil
}

func (this *ProductInteractor) mergeComplementaryValues(productIds []int64, productJsonSlice []*json.Product) error {
	variantsMap, err := this.VariantInteractor.GetJsonVariantsMap(productIds)
	if err != nil {
		return errors.New("Error getting variants: " + err.Error())
	}

	productPropertiesMap, err := this.ProductPropertyInteractor.GetJsonProductPropertiesMap(productIds)
	if err != nil {
		return errors.New("Error getting product properties: " + err.Error())
	}

	classificationsMap, err := this.ClassificationInteractor.GetJsonClassificationsMap(productIds)
	if err != nil {
		return errors.New("Error getting classifications: " + err.Error())
	}

	optionTypesMap, err := this.OptionTypeInteractor.GetJsonOptionTypesMap(productIds)
	if err != nil {
		return errors.New("Error getting option types: " + err.Error())
	}

	this.mergeVariants(productJsonSlice, variantsMap)

	this.mergeProductProperties(productJsonSlice, productPropertiesMap)

	this.mergeClassifications(productJsonSlice, classificationsMap)

	this.mergeOptionTypes(productJsonSlice, optionTypesMap)

	return nil
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
		productJson := &json.Product{}
		copier.Copy(productJson, product)

		jsonProductsSlice = append(jsonProductsSlice, productJson)
	}

	return jsonProductsSlice
}

func (this *ProductInteractor) toJson(product *models.Product) *json.Product {
	productJson := &json.Product{
		Id:          product.Id,
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

		variantSlice := variantsMap[product.Id]

		if variantSlice == nil {
			continue
		}

		for _, variant := range variantSlice {
			variant.Description = product.Description
			variant.Slug = product.Slug
			variant.Name = product.Name

			if variant.IsMaster {
				product.Master = *variant
				product.Price = variant.Price
			} else {
				product.Variants = append(product.Variants, *variant)
			}

			totalOnHand += variant.TotalOnHand

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

		productPropertiesSlice := productPropertiesMap[product.Id]

		if productPropertiesSlice == nil {
			continue
		}

		for _, productProperty := range productPropertiesSlice {
			product.ProductProperties = append(product.ProductProperties, *productProperty)
		}
	}
}

func (this *ProductInteractor) mergeClassifications(productSlice []*json.Product, classificationsMap JsonClassificationsMap) {
	for _, product := range productSlice {
		product.TaxonIds = []int{}
		product.Classifications = []json.Classification{}

		classificationsSlice := classificationsMap[product.Id]

		if classificationsSlice == nil {
			continue
		}

		for _, classification := range classificationsSlice {
			product.Classifications = append(product.Classifications, *classification)
			product.TaxonIds = append(product.TaxonIds, int(classification.TaxonId))
		}
	}
}

func (this *ProductInteractor) mergeOptionTypes(productSlice []*json.Product, optionTypesMap JsonOptionTypesMap) {
	for _, product := range productSlice {
		product.OptionTypes = []json.OptionType{}

		optionTypesSlice := optionTypesMap[product.Id]

		if optionTypesSlice == nil {
			continue
		}

		for _, optionType := range optionTypesSlice {
			product.OptionTypes = append(product.OptionTypes, *optionType)
		}
	}
}

func (this *ProductInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	query, gparams, err := params.GetGransakParams()
	if err != nil {
		return 0, err
	}
	return this.ProductRepo.CountAll(query, gparams)
}
