package repositories

import (
	"github.com/jinzhu/gorm"

	"github.com/crowdint/gopher-spree-api/domain/models"
)

type ProductRepo DbRepo

func NewProductRepo() *ProductRepo {
	return &ProductRepo{
		dbHandler: Spree_db,
	}
}

func (this *ProductRepo) FindById(id int64) (*models.Product, error) {
	product := &models.Product{Id: id}

	if id == 0 {
		return product, gorm.RecordNotFound
	}

	query := this.dbHandler.Find(product)

	return product, query.Error
}

func (this *ProductRepo) List(currentPage, perPage int, gransakQuery string, params []interface{}) ([]*models.Product, error) {
	var products []*models.Product

	offset := (currentPage - 1) * perPage

	var query *gorm.DB

	if gransakQuery == "" {
		query = this.dbHandler.Offset(offset).Limit(perPage).Order("created_at desc").Find(&products)
	} else {
		query = this.dbHandler.Where(gransakQuery, params).Offset(offset).Limit(perPage).Order("created_at desc").Find(&products)
	}

	return products, query.Error
}

func (this *ProductRepo) CountAll(queryFilter string, params []interface{}) (int64, error) {
	var count int64

	var query *gorm.DB
	if queryFilter == "" {
		query = this.dbHandler.Model(models.Product{}).Count(&count)
	} else {
		query = this.dbHandler.Model(models.Product{}).Where(queryFilter, params).Count(&count)
	}

	return count, query.Error
}
