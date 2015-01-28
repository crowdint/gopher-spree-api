package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type TaxonRepo DbRepo

func NewTaxonRepo() *TaxonRepo {
	return &TaxonRepo{
		dbHandler: spree_db,
	}
}

func (this *TaxonRepo) FindById(id int64) *models.Taxon {
	taxon := &models.Taxon{
		ID: id,
	}

	this.dbHandler.First(taxon)

	return taxon
}

func (this *TaxonRepo) List() ([]models.Taxon, error) {
	taxon := &models.Taxon{}

	rows, err := this.dbHandler.Find(taxon).Rows()
	if err != nil {
		return nil, err
	}

	result, err := ParseAllRows(&models.Taxon{}, rows)
	if err != nil {
		return nil, err
	}

	taxonSlice := this.toTaxonSlice(result)

	return taxonSlice, nil
}

func (this *TaxonRepo) toTaxonSlice(result []interface{}) []models.Taxon {
	taxonSlice := []models.Taxon{}

	for _, element := range result {
		taxon := element.(models.Taxon)

		taxonSlice = append(taxonSlice, taxon)
	}

	return taxonSlice
}
