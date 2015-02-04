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

func (this *TaxonRepo) FindByProductIds(productIds []int64) ([]*models.Taxon, error) {

	taxons := []*models.Taxon{}

	if len(productIds) == 0 {
		return taxons, nil
	}

	prettyNameSelectString := "(SELECT Array_to_string(Array_agg(name), ' -> ') " +
		"FROM " +
		"(SELECT " +
		"spree_taxons.name " +
		"FROM spree_taxons " +
		"WHERE spree_taxons.lft <= taxons.lft " +
		"AND spree_taxons.rgt >= taxons.rgt " +
		"GROUP BY spree_taxons.name, spree_taxons.lft " +
		"ORDER BY spree_taxons.lft " +
		") AS tree_path " +
		") AS pretty_name "

	selectString := "taxons.*, " +
		"spree_products_taxons.product_id, " +
		"spree_products_taxons.position AS classification_position, " +
		prettyNameSelectString

	productsTaxonsJoinString := "INNER JOIN spree_products_taxons " +
		"ON taxons.id = spree_products_taxons.taxon_id "

	query := this.dbHandler.
		Table("spree_taxons taxons").
		Select(selectString).
		Joins(productsTaxonsJoinString).
		Where("spree_products_taxons.product_id IN (?)", productIds).
		Scan(&taxons)

	return taxons, query.Error
}
