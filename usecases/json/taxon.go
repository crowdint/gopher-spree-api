package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type TaxonInteractor struct {
	Repo *repositories.TaxonRepo
}

func NewTaxonInteractor() *TaxonInteractor {
	return &TaxonInteractor{
		Repo: repositories.NewTaxonRepo(),
	}
}

type JsonTaxonsMap map[int64][]*json.Taxon

func (this *TaxonInteractor) GetJsonTaxonsMap(productIds []int64) (JsonTaxonsMap, error) {

	taxons, err := this.Repo.FindByProductIds(productIds)
	if err != nil {
		return JsonTaxonsMap{}, err
	}

	taxonsJson := this.modelsToJsonTaxonsMap(taxons)

	return taxonsJson, nil
}

func (this *TaxonInteractor) modelsToJsonTaxonsMap(taxonSlice []*models.Taxon) JsonTaxonsMap {
	jsonTaxonsMap := JsonTaxonsMap{}

	for _, taxon := range taxonSlice {
		taxonJson := this.toJson(taxon)

		if _, exists := jsonTaxonsMap[taxon.Id]; !exists {
			jsonTaxonsMap[taxon.Id] = []*json.Taxon{}
		}

		jsonTaxonsMap[taxon.Id] = append(jsonTaxonsMap[taxon.Id], taxonJson)

	}

	return jsonTaxonsMap
}

func (this *TaxonInteractor) toJson(taxon *models.Taxon) *json.Taxon {
	taxonJson := &json.Taxon{
		ID:         taxon.Id,
		Name:       taxon.Name,
		PrettyName: taxon.PrettyName,
		Permalink:  taxon.Permalink,
		ParentID:   taxon.ParentId,
		TaxonomyID: taxon.TaxonomyId,
		//Taxons: this.GetChildren(taxon.Id),
	}

	return taxonJson
}

func (this *TaxonInteractor) GetChildren(taxonId int64) []*json.Taxon {
	return []*json.Taxon{}
}
