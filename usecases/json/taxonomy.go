package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/jinzhu/copier"
)

type TaxonomyResponse struct {
	data []*json.Taxonomy
}

func (this TaxonomyResponse) GetCount() int {
	return len(this.data)
}

func (this TaxonomyResponse) GetData() interface{} {
	return this.data
}

func (this TaxonomyResponse) GetTag() string {
	return "taxonomies"
}

type TaxonomyInteractor struct {
	BaseRepository *repositories.DbRepo
}

func NewTaxonomyInteractor() *TaxonomyInteractor {
	return &TaxonomyInteractor{
		BaseRepository: repositories.NewDatabaseRepository(),
	}
}

func (this *TaxonomyInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	query, gparams, err := params.GetGransakParams()
	if err != nil {
		return TaxonomyResponse{}, err
	}

	var taxonomyModelSlice []*models.Taxonomy
	this.BaseRepository.All(&taxonomyModelSlice, map[string]interface{}{
		"limit":  perPage,
		"offset": currentPage,
		"order":  "created_at desc",
	}, query, gparams)
	if err != nil {
		return TaxonomyResponse{}, err
	}

	taxonomyJsonSlice, err := this.transformToJsonResponse(taxonomyModelSlice)
	if err != nil {
		return TaxonomyResponse{}, err
	}

	return TaxonomyResponse{
		data: taxonomyJsonSlice,
	}, nil
}

func (this *TaxonomyInteractor) GetShowResponse(param ResponseParameters) (interface{}, error) {
	return true, nil
}

func (this *TaxonomyInteractor) transformToJsonResponse(taxonomyModelSlice []*models.Taxonomy) ([]*json.Taxonomy, error) {
	taxonomyJsonSlice := this.modelsToJsonTaxonomiesSlice(taxonomyModelSlice)

	//WIP MERGE TAXONS

	return taxonomyJsonSlice, nil
}

func (this *TaxonomyInteractor) getIdSlice(taxonomySlice []*models.Taxonomy) []int64 {
	taxonomyIds := []int64{}

	for _, taxonomy := range taxonomySlice {
		taxonomyIds = append(taxonomyIds, taxonomy.Id)
	}

	return taxonomyIds
}

func (this *TaxonomyInteractor) modelsToJsonTaxonomiesSlice(taxonomySlice []*models.Taxonomy) []*json.Taxonomy {
	jsonTaxonomySlice := []*json.Taxonomy{}

	for _, taxonomy := range taxonomySlice {
		taxonomyJson := &json.Taxonomy{}
		copier.Copy(taxonomyJson, taxonomy)

		jsonTaxonomySlice = append(jsonTaxonomySlice, taxonomyJson)
	}

	return jsonTaxonomySlice
}

func (this *TaxonomyInteractor) GetTotalCount(param ResponseParameters) (int64, error) {
	query, gparams, err := param.GetGransakParams()
	if err != nil {
		return 0, err
	}
	return this.BaseRepository.Count(models.Taxonomy{}, query, gparams)
}
