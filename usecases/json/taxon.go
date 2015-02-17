package json

import (
	"github.com/jinzhu/copier"

	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type TaxonResponse struct {
	data []*json.Taxon
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
	TaxonRepo *repositories.TaxonRepo
}

func NewTaxonInteractor() *TaxonInteractor {
	return &TaxonInteractor{
		TaxonRepo: repositories.NewTaxonRepo(),
	}
}

func (this *TaxonInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	query, gparams, err := params.GetGransakParams()
	if err != nil {
		return TaxonResponse{}, err
	}

	taxonModelSlice, err := this.TaxonRepo.List(currentPage, perPage, query, gparams)
	if err != nil {
		return TaxonResponse{}, err
	}

	taxonJsonSlice := this.modelsToJsonTaxonsSlice(taxonModelSlice)

	return TaxonResponse{
		data: taxonJsonSlice,
	}, nil
}

func (this *TaxonInteractor) modelsToJsonTaxonsSlice(taxonSlice []*models.Taxon) []*json.Taxon {
	jsonTaxonsSlice := []*json.Taxon{}

	for _, taxon := range taxonSlice {
		taxonJson := &json.Taxon{}
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
	return this.TaxonRepo.CountAll(query, gparams)
}

func (this *TaxonInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	taxonModelSlice := []*models.Taxon{}

	//DUMMY UNTIL TAXON SHOW IS IMPLEMENTED

	return taxonModelSlice[0], nil
}
