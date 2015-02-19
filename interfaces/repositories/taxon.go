package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type TaxonRepo DbRepo

func NewTaxonRepo() *TaxonRepo {
	return &TaxonRepo{
		dbHandler: Spree_db,
	}
}

var queryConfig = map[string]string{}

func (this *TaxonRepo) FindByProductIds(productIds []int64) ([]*models.Taxon, error) {
	queryConfig["selectString"] = "taxons.*, " +
		"spree_products_taxons.product_id, " +
		"spree_products_taxons.position AS classification_position "

	queryConfig["joinString"] = "INNER JOIN spree_products_taxons " +
		"ON taxons.id = spree_products_taxons.taxon_id "

	queryConfig["queryString"] = "spree_products_taxons.product_id IN (?)"

	return this.findByResourceIds(productIds)
}

func (this *TaxonRepo) FindByTaxonomyIds(taxonomyIds []int64) ([]*models.Taxon, error) {
	queryConfig["selectString"] = "taxons.*, " +
		"spree_taxonomies.id, " +
		"spree_taxonomies.name AS taxonomy_name "

	queryConfig["joinString"] = "INNER JOIN spree_taxonomies " +
		"ON taxons.taxonomy_id = spree_taxonomies.id "

	queryConfig["queryString"] = "spree_taxonomies.id IN (?)"

	return this.findByResourceIds(taxonomyIds)
}

func (this *TaxonRepo) findByResourceIds(resourceIds []int64) ([]*models.Taxon, error) {
	taxons := []*models.Taxon{}

	if len(resourceIds) == 0 {
		return taxons, nil
	}

	query := this.dbHandler.
		Table("spree_taxons taxons").
		Select(queryConfig["selectString"]).
		Joins(queryConfig["joinString"]).
		Where(queryConfig["queryString"], resourceIds).
		Scan(&taxons)

	return taxons, query.Error
}
