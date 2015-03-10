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
	queryData, err := params.GetQuery()
	if err != nil {
		return TaxonomyResponse{}, err
	}

	query := queryData.Query
	gparams := queryData.Params

	var taxonomies []*domain.Taxonomy
	this.BaseRepository.All(&taxonomies, map[string]interface{}{
		"limit":  perPage,
		"offset": currentPage,
		"order":  "created_at desc",
	}, query, gparams)
	if err != nil {
		return TaxonomyResponse{}, err
	}

	this.mergeTaxons(taxonomies)

	return TaxonomyResponse{
		data: taxonomies,
	}, nil
}

func (this *TaxonomyInteractor) GetShowResponse(param ResponseParameters) (interface{}, error) {
	return true, nil
}

func (this *TaxonomyInteractor) GetTotalCount(param ResponseParameters) (int64, error) {
	queryData, err := param.GetQuery()
	if err != nil {
		return 0, err
	}

	query := queryData.Query
	gparams := queryData.Params

	return this.BaseRepository.Count(domain.Taxonomy{}, query, gparams)
}

func (this *TaxonomyInteractor) mergeTaxons(taxonomies []*domain.Taxonomy) {
	taxonomyIds := []int64{}
	for _, taxonomy := range taxonomies {
		taxonomyIds = append(taxonomyIds, taxonomy.Id)
	}

	var taxons []*domain.Taxon
	this.BaseRepository.All(&taxons, nil, "taxonomy_id IN (?)", taxonomyIds)
	this.toTaxonTree(taxons, 2)

	for _, taxonomy := range taxonomies {
		for _, taxon := range taxons {
			if taxon.TaxonomyId == taxonomy.Id && taxon.ParentId == 0 {
				taxonomy.Root = taxon
			}
		}
	}
}

func (this *TaxonomyInteractor) toTaxonTree(nodes []*domain.Taxon, maxDepth int64) {
	for _, node := range nodes {
		for _, childNode := range nodes {
			if node.Lft < childNode.Rgt && node.Rgt > childNode.Rgt && (node.Depth+1) == childNode.Depth && childNode.Depth < maxDepth {
				childNode.PrettyName = node.PrettyName + " -> " + childNode.PrettyName
				node.Taxons = append(node.Taxons, childNode)
			}
		}
	}
}
