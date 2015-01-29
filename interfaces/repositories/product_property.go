package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type ProductPropertyRepo DbRepo

func NewProductPropertyRepo() *ProductPropertyRepo {
  return &ProductPropertyRepo {
    dbHandler: Spree_db,
  }
}

func (this *ProductPropertyRepo) FindByProductIds(productIds []int64) ([]*models.ProductProperty, error) {
  var productProperties []*models.ProductProperty

  query := this.dbHandler.Where("product_id in (?)", productIds).Find(&productProperties)
  return productProperties, query.Error
}
