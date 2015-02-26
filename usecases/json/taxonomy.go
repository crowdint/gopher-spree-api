package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type TaxonomyResponse struct {
	data []*domain.Taxonomy
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
	BaseRepository *repositories.DbRepository
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

	var taxonomyModelSlice []*domain.Taxonomy
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

func (this *TaxonomyInteractor) transformToJsonResponse(taxonomyModelSlice []*domain.Taxonomy) ([]*domain.Taxonomy, error) {
	//WIP MERGE TAXONS

	return taxonomyModelSlice, nil
}

func (this *TaxonomyInteractor) getIdSlice(taxonomySlice []*domain.Taxonomy) []int64 {
	taxonomyIds := []int64{}

	for _, taxonomy := range taxonomySlice {
		taxonomyIds = append(taxonomyIds, taxonomy.Id)
	}

	return taxonomyIds
}

func (this *TaxonomyInteractor) GetTotalCount(param ResponseParameters) (int64, error) {
	query, gparams, err := param.GetGransakParams()
	if err != nil {
		return 0, err
	}
	return this.BaseRepository.Count(domain.Taxonomy{}, query, gparams)
}
