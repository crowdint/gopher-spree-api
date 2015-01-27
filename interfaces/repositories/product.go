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

func (this *ProductRepo) List() ([]models.Product, error) {
	product := &models.Product{}

	rows, err := this.dbHandler.Find(product).Rows()
	if err != nil {
		return nil, err
	}

	result, err := ParseAllRows(&models.Product{}, rows)
	if err != nil {
		return nil, err
	}

	productSlice := this.toProductSlice(result)

	return productSlice, nil
}

func (this *ProductRepo) toProductSlice(result []interface{}) []models.Product {
	productSlice := []models.Product{}

	for _, element := range result {
		product := element.(models.Product)

		productSlice = append(productSlice, product)
	}

	return productSlice
}
