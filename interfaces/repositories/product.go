package repositories

import "github.com/crowdint/gopher-spree-api/domain"

type ProductRepository struct {
	DbRepository
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		DbRepository{Spree_db},
	}
}

func (this *ProductRepository) Create(product *domain.Product) (err error) {
	err = this.CreateWithSlug(product)
	if err != nil {
		return
	}

	variantRepository := NewVariantRepository()
	// TODO: AfterCreate for each variant in product.Variants
	err = variantRepository.AfterCreate(product.Master)

	return
}
