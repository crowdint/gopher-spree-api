package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type ProductRepo DbRepo

func NewProductRepo() *ProductRepo {
	return &ProductRepo{
		dbHandler: spree_db,
	}
}

func (this *ProductRepo) FindById(id int64) *models.Product {
	product := &models.Product{
		ID: id,
	}

	this.dbHandler.First(product)

	return product
}

func (this *ProductRepo) List() []*models.Product {
	var products []*models.Product

	this.dbHandler.Find(&products)

	return products
}
