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

func (this *ProductRepo) List() ([]*models.Product, error) {
	var products []*models.Product

	query := this.dbHandler.Find(&products)

	return products, query.Error
}
