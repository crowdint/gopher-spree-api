package json

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/crowdint/gopher-spree-api/cache"
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type ProductResponse struct {
	data []*domain.Product
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
	ProductRepository         *repositories.ProductRepository
	VariantInteractor         *VariantInteractor
	ProductPropertyInteractor *ProductPropertyInteractor
	ClassificationInteractor  *ClassificationInteractor
	OptionTypeInteractor      *OptionTypeInteractor
}

func NewProductInteractor() *ProductInteractor {
	return &ProductInteractor{
		ProductRepository:         repositories.NewProductRepository(),
		VariantInteractor:         NewVariantInteractor(),
		ProductPropertyInteractor: NewProductPropertyInteractor(),
		ClassificationInteractor:  NewClassificationInteractor(),
		OptionTypeInteractor:      NewOptionTypeInteractor(),
	}
}

func (this *ProductInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	queryData, err := params.GetQuery()
	if err != nil {
		utils.LogrusError("GetResponse", "GET", err)

		return ProductResponse{}, err
	}

	query := queryData.Query
	gparams := queryData.Params

	var productModelSlice []*domain.Product
	err = this.ProductRepository.All(&productModelSlice, map[string]interface{}{
		"limit":  perPage,
		"offset": currentPage,
		"order":  "created_at desc",
	}, query, gparams)

	if err != nil {
		utils.LogrusError("GetResponse", "GET", err)

		return ProductResponse{}, err
	}

	productsCached := this.toCacheData(productModelSlice)
	missingProductsCached, _ := cache.FetchMulti(productsCached)
	if len(missingProductsCached) == 0 {
		return ProductResponse{data: productModelSlice}, nil
	}

	missingProductsIds, missingProducts := this.getMissingProductsFromMissingData(&missingProductsCached)
	if err = this.mergeComplementaryValues(missingProductsIds, missingProducts); err != nil {
		return ProductResponse{}, err
	}

	cache.SetMulti(missingProductsCached)

	return ProductResponse{data: productModelSlice}, nil
}

func (this *ProductInteractor) toCacheData(productSlice []*domain.Product) (productsCached []cache.Cacheable) {
	for _, product := range productSlice {
		productsCached = append(productsCached, product)
	}
	return
}

func (this *ProductInteractor) getMissingProductsFromMissingData(missingData *[]cache.Cacheable) ([]int64, []*domain.Product) {
	missingProductsIds := []int64{}
	missingProducts := []*domain.Product{}
	for _, missingProduct := range *missingData {
		p := missingProduct.(*domain.Product)
		missingProducts = append(missingProducts, p)
		missingProductsIds = append(missingProductsIds, p.Id)
	}

	return missingProductsIds, missingProducts
}

func (this *ProductInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	id, err := params.GetIntParam(ID_PARAM)

	if err != nil {
		utils.LogrusError("GetShowResponse", "GET", err)

		return struct{}{}, errors.New("Invalid parameter type: " + fmt.Sprintf("%v", id))
	}

	product := &domain.Product{}
	err = this.ProductRepository.FindBy(product, nil, map[string]interface{}{"id": id})
	if err != nil {
		utils.LogrusError("GetShowResponse", "GET", err)

		return nil, err
	}

	if err = cache.Find(product); err == nil {
		return product, nil
	}

	productModelSlice := []*domain.Product{}

	productModelSlice = append(productModelSlice, product)

	productJsonSlice, err := this.transformToJsonResponse(productModelSlice)
	if err != nil {
		utils.LogrusError("GetShowResponse", "GET", err)

		return nil, err
	}

	if err = cache.Set(productJsonSlice[0]); err != nil {
		utils.LogrusError("GetShowResponse", "GET", err)

		log.Println("An error occurred while setting the cache: ", err.Error())
	}

	return productJsonSlice[0], nil
}

func (this *ProductInteractor) GetCreateResponse(params ResponseParameters) (interface{}, interface{}, error) {
	productParams := &domain.ProductParams{}
	ok := params.BindPermittedParams("product", productParams)

	if !ok {
		utils.LogrusError("GetCreateResponse", "", errors.New("Error occurred while parsing request parameters."))

		return struct{}{}, nil, errors.New("Error occurred while parsing request parameters.")
	}

	this.setUpShippingCategory(productParams)
	product := domain.NewProductFromPermittedParams(productParams)

	if err := this.ProductRepository.Create(product); err != nil {
		if err == domain.ErrNotValid {
			utils.LogrusError("GetCreateResponse", "", err)

			return struct{}{}, product.Errors(), err
		}
		utils.LogrusError("GetCreateResponse", "", err)

		return struct{}{}, nil, err
	}

	return product, nil, nil
}

func (this *ProductInteractor) setUpShippingCategory(productParams *domain.ProductParams) {
	if category := productParams.PermittedProductParams.ShippingCategory; category != "" {
		shippingCategory := &domain.ShippingCategory{}
		err := this.ProductRepository.FirstOrCreate(shippingCategory, map[string]interface{}{"name": category})
		if err == nil {
			productParams.PermittedProductParams.ShippingCategoryId = shippingCategory.Id
		}
	}
}

func (this *ProductInteractor) transformToJsonResponse(productModelSlice []*domain.Product) ([]*domain.Product, error) {
	productIds := this.getIdSlice(productModelSlice)

	err := this.mergeComplementaryValues(productIds, productModelSlice)
	if err != nil {
		utils.LogrusError("transformToJsonResponse", "", err)

		return []*domain.Product{}, err
	}

	return productModelSlice, nil
}

func (this *ProductInteractor) mergeComplementaryValues(productIds []int64, productJsonSlice []*domain.Product) error {
	variantsMap, err := this.VariantInteractor.GetJsonVariantsMap(productIds)
	if err != nil {
		utils.LogrusError("mergeComplementaryValues", "", errors.New("Error getting variants: "+err.Error()))
		return errors.New("Error getting variants: " + err.Error())
	}

	productPropertiesMap, err := this.ProductPropertyInteractor.GetJsonProductPropertiesMap(productIds)
	if err != nil {
		utils.LogrusError("mergeComplementaryValues", "", errors.New("Error getting product properties: "+err.Error()))

		return errors.New("Error getting product properties: " + err.Error())
	}

	classificationsMap, err := this.ClassificationInteractor.GetJsonClassificationsMap(productIds)
	if err != nil {
		utils.LogrusError("mergeComplementaryValues", "", errors.New("Error getting classifications: "+err.Error()))

		return errors.New("Error getting classifications: " + err.Error())
	}

	optionTypesMap, err := this.OptionTypeInteractor.GetJsonOptionTypesMap(productIds)
	if err != nil {
		utils.LogrusError("mergeComplementaryValues", "", errors.New("Error getting option types: "+err.Error()))

		return errors.New("Error getting option types: " + err.Error())
	}

	this.mergeVariants(productJsonSlice, variantsMap)

	this.mergeProductProperties(productJsonSlice, productPropertiesMap)

	this.mergeClassifications(productJsonSlice, classificationsMap)

	this.mergeOptionTypes(productJsonSlice, optionTypesMap)

	return nil
}

func (this *ProductInteractor) getIdSlice(productSlice []*domain.Product) []int64 {
	productIds := []int64{}

	for _, product := range productSlice {
		productIds = append(productIds, product.Id)
	}

	return productIds
}

func (this *ProductInteractor) mergeVariants(productSlice []*domain.Product, variantsMap JsonVariantsMap) {
	for _, product := range productSlice {
		product.Variants = []domain.Variant{}
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
				product.Price = strconv.FormatFloat(*variant.Price, 'f', 2, 64)
			} else {
				product.Variants = append(product.Variants, *variant)
			}

			if variant.TotalOnHand != nil {
				totalOnHand += *variant.TotalOnHand
			}
		}

		product.TotalOnHand = totalOnHand

		if len(product.Variants) > 0 {
			product.HasVariants = true
		}
	}
}

func (this *ProductInteractor) mergeProductProperties(productSlice []*domain.Product, productPropertiesMap JsonProductPropertiesMap) {
	for _, product := range productSlice {
		product.ProductProperties = []domain.ProductProperty{}

		productPropertiesSlice := productPropertiesMap[product.Id]

		if productPropertiesSlice == nil {
			continue
		}

		for _, productProperty := range productPropertiesSlice {
			product.ProductProperties = append(product.ProductProperties, *productProperty)
		}
	}
}

func (this *ProductInteractor) mergeClassifications(productSlice []*domain.Product, classificationsMap JsonClassificationsMap) {
	for _, product := range productSlice {
		product.TaxonIds = []int{}
		product.Classifications = []domain.Classification{}

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

func (this *ProductInteractor) mergeOptionTypes(productSlice []*domain.Product, optionTypesMap JsonOptionTypesMap) {
	for _, product := range productSlice {
		product.OptionTypes = []domain.OptionType{}

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
	queryData, err := params.GetQuery()
	if err != nil {
		utils.LogrusError("GetTotalCount", "", err)

		return 0, err
	}

	query := queryData.Query
	gparams := queryData.Params

	return this.ProductRepository.Count(&domain.Product{}, query, gparams)
}
