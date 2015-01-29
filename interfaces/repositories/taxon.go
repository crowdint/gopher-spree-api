package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type TaxonRepo DbRepo

func NewTaxonRepo() *TaxonRepo {
	return &TaxonRepo{
		dbHandler: Spree_db,
	}
}

func (this *TaxonRepo) FindById(id int64) *models.Taxon {
	taxon := &models.Taxon{
		Id: id,
	}

	this.dbHandler.First(taxon)

	return taxon
}

func (this *TaxonRepo) List() []*models.Taxon {
	var taxons []*models.Taxon

	this.dbHandler.Find(&taxons)

	return taxons
}
