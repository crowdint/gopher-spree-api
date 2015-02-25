package repositories

import "github.com/crowdint/gopher-spree-api/domain/json"

type TaxonRepository DbRepository

func NewTaxonRepo() *TaxonRepository {
	return &TaxonRepository{
		dbHandler: Spree_db,
	}
}

var queryConfig = map[string]string{}

func (this *TaxonRepository) FindByProductIds(productIds []int64) ([]*json.Taxon, error) {
	queryConfig["selectString"] = "taxons.*, " +
		"spree_products_taxons.product_id, " +
		"spree_products_taxons.position AS classification_position "

	queryConfig["joinString"] = "INNER JOIN spree_products_taxons " +
		"ON taxons.id = spree_products_taxons.taxon_id "

	queryConfig["queryString"] = "spree_products_taxons.product_id IN (?)"

	return this.findByResourceIds(productIds)
}

func (this *TaxonRepository) FindByTaxonomyIds(taxonomyIds []int64) ([]*json.Taxon, error) {
	queryConfig["selectString"] = "taxons.*, " +
		"spree_taxonomies.id, " +
		"spree_taxonomies.name AS taxonomy_name "

	queryConfig["joinString"] = "INNER JOIN spree_taxonomies " +
		"ON taxons.taxonomy_id = spree_taxonomies.id "

	queryConfig["queryString"] = "spree_taxonomies.id IN (?)"

	return this.findByResourceIds(taxonomyIds)
}

func (this *TaxonRepository) findByResourceIds(resourceIds []int64) ([]*json.Taxon, error) {
	taxons := []*json.Taxon{}

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
