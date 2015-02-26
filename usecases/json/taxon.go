package json

import (
	"github.com/jinzhu/copier"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
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
	query, gparams, err := params.GetGransakParams()
	if err != nil {
		return TaxonResponse{}, err
	}

	var taxonModelSlice []*domain.Taxon

	err = this.BaseRepository.All(&taxonModelSlice, map[string]interface{}{
		"limit":  perPage,
		"offset": currentPage,
		"order":  "created_at desc",
	}, query, gparams)
	if err != nil {
		return TaxonResponse{}, err
	}

	taxonJsonSlice := this.modelsToJsonTaxonsSlice(taxonModelSlice)

	this.toTaxonTree(taxonJsonSlice)

	return TaxonResponse{
		data: taxonJsonSlice,
	}, nil
}

func (this *TaxonInteractor) modelsToJsonTaxonsSlice(taxonSlice []*domain.Taxon) []*domain.Taxon {
	jsonTaxonsSlice := []*domain.Taxon{}

	for _, taxon := range taxonSlice {
		taxonJson := &domain.Taxon{
			Taxons: []*domain.Taxon{},
		}

		copier.Copy(taxonJson, taxon)
		jsonTaxonsSlice = append(jsonTaxonsSlice, taxonJson)
	}

	return jsonTaxonsSlice
}

func (this *TaxonInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	query, gparams, err := params.GetGransakParams()
	if err != nil {
		return 0, err
	}
	return this.BaseRepository.Count(domain.Taxon{}, query, gparams)
}

func (this *TaxonInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	taxonModelSlice := []*domain.Taxon{}

	//DUMMY UNTIL TAXON SHOW IS IMPLEMENTED

	return taxonModelSlice[0], nil
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
