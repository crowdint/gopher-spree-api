package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type ClassificationInteractor struct {
	TaxonRepo *repositories.TaxonRepo
}

func NewClassificationInteractor() *ClassificationInteractor {
	return &ClassificationInteractor{
		TaxonRepo: repositories.NewTaxonRepo(),
	}
}

type JsonClassificationsMap map[int64][]*json.Classification

func (this *ClassificationInteractor) GetJsonClassificationsMap(productIds []int64) (JsonClassificationsMap, error) {

	taxons, err := this.TaxonRepo.FindByProductIds(productIds)
	if err != nil {
		return JsonClassificationsMap{}, err
	}

	classificationsJson := this.taxonsToClassificationMap(taxons)

	return classificationsJson, nil
}

func (this *ClassificationInteractor) taxonsToClassificationMap(taxonsSlice []*models.Taxon) JsonClassificationsMap {
	jsonClassificationsMap := JsonClassificationsMap{}

	for _, taxon := range taxonsSlice {
		classificationJson := this.taxonToJsonClassification(taxon)

		if _, exists := jsonClassificationsMap[taxon.ProductId]; !exists {
			jsonClassificationsMap[taxon.ProductId] = []*json.Classification{}
		}

		jsonClassificationsMap[taxon.ProductId] = append(jsonClassificationsMap[taxon.ProductId], classificationJson)
	}

	return jsonClassificationsMap
}

func (this *ClassificationInteractor) taxonToJsonClassification(taxon *models.Taxon) *json.Classification {
	classificationJson := &json.Classification{
		TaxonId:  taxon.Id,
		Position: taxon.ClassificationPosition,
		Taxon:    *this.taxonToJson(taxon),
	}

	return classificationJson
}

func (this *ClassificationInteractor) taxonToJson(taxon *models.Taxon) *json.Taxon {
	taxonJson := &json.Taxon{
		Id:         taxon.Id,
		Name:       taxon.Name,
		PrettyName: taxon.PrettyName,
		Permalink:  taxon.Permalink,
		ParentId:   taxon.ParentId,
		TaxonomyId: taxon.TaxonomyId,
	}

	return taxonJson
}
