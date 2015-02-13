package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/jinzhu/gorm"
)

type TaxonomyRepo DbRepo

func NewTaxonomyRepo() *TaxonomyRepo {
	return &TaxonomyRepo{
		dbHandler: Spree_db,
	}
}

func (this *TaxonomyRepo) List(currentPage, perPage int, gransakQuery string) ([]*models.Taxonomy, error) {
	var taxonomies []*models.Taxonomy

	offset := (currentPage - 1) * perPage

	var query *gorm.DB

	if gransakQuery == "" {
		query = this.dbHandler.Offset(offset).Limit(perPage).Order("created_at desc").Find(&taxonomies)
	} else {
		query = this.dbHandler.Where(gransakQuery).Offset(offset).Limit(perPage).Order("created_at desc").Find(&taxonomies)
	}

	return taxonomies, query.Error
}

func (this *TaxonomyRepo) CountAll(queryFilter string) (int64, error) {
	var count int64

	var query *gorm.DB
	if queryFilter == "" {
		query = this.dbHandler.Model(models.Taxonomy{}).Count(&count)
	} else {
		query = this.dbHandler.Model(models.Taxonomy{}).Where(queryFilter).Count(&count)
	}

	return count, query.Error
}
