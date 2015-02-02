package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type ProductRepo DbRepo

func NewProductRepo() *ProductRepo {
	return &ProductRepo{
		dbHandler: Spree_db,
	}
}

func (this *ProductRepo) FindById(id int64) (*models.Product, error) {
	product := &models.Product{
		Id: id,
	}

	query := this.dbHandler.First(product)

	return product, query.Error
}

func (this *ProductRepo) List(currentPage, perPage int) ([]*models.Product, error) {
	var products []*models.Product

	offset := (currentPage - 1) * perPage

	query := this.dbHandler.Offset(offset).Limit(perPage).Order("created_at desc").Find(&products)

	return products, query.Error
}

func (this *ProductRepo) CountAll() (int64, error) {
	var products []*models.Product
	var count int64

	query := this.dbHandler.Find(&products).Count(&count)

	return count, query.Error
}
