package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type TaxonResponse struct {
	data []*domain.Taxon
}

func (this TaxonResponse) GetCount() int {
	return len(this.data)
}

func (this TaxonResponse) GetData() interface{} {
	return this.data
}

func (this TaxonResponse) GetTag() string {
	return "taxons"
}

type TaxonInteractor struct {
	BaseRepository  *repositories.DbRepository
	TaxonRepository *repositories.TaxonRepository
}

func NewTaxonInteractor() *TaxonInteractor {
	return &TaxonInteractor{
		BaseRepository:  repositories.NewDatabaseRepository(),
		TaxonRepository: repositories.NewTaxonRepo(),
	}
}

func (this *TaxonInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	queryData, err := params.GetQuery()
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return TaxonResponse{}, err
	}

	query := queryData.Query
	gparams := queryData.Params

	var taxonModelSlice []*domain.Taxon

	err = this.BaseRepository.All(&taxonModelSlice, map[string]interface{}{
		"limit":  perPage,
		"offset": currentPage,
		"order":  "created_at desc",
	}, query, gparams)
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return TaxonResponse{}, err
	}

	this.toTaxonTree(taxonModelSlice)

	return TaxonResponse{
		data: taxonModelSlice,
	}, nil
}

func (this *TaxonInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	queryData, err := params.GetQuery()
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)

		return 0, err
	}

	query := queryData.Query
	gparams := queryData.Params

	return this.BaseRepository.Count(domain.Taxon{}, query, gparams)
}

func (this *TaxonInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	taxonModelSlice := []*domain.Taxon{}

	//DUMMY UNTIL TAXON SHOW IS IMPLEMENTED

	return taxonModelSlice[0], nil
}

func (this *TaxonInteractor) GetCreateResponse(params ResponseParameters) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (this *TaxonInteractor) toTaxonTree(nodes []*domain.Taxon) {
	for _, node := range nodes {
		for _, childNode := range nodes {
			if node.Lft < childNode.Rgt && node.Rgt > childNode.Rgt && (node.Depth+1) == childNode.Depth {
				childNode.PrettyName = node.PrettyName + " -> " + childNode.PrettyName
				node.Taxons = append(node.Taxons, childNode)
			}
		}
	}
}
