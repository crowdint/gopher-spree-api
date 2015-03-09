package repositories

import (
	"strconv"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestProductRepository_Create(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	productRepository := NewProductRepository()
	product := &domain.Product{
		Name:               "Test Product",
		Description:        "Test Description",
		Price:              "12.50",
		Slug:               "test-product",
		ShippingCategoryId: 1,
	}

	price, err := strconv.ParseFloat(product.Price, 64)
	position := int64(1)
	product.Master = domain.Variant{
		IsMaster:     true,
		Price:        &price,
		Product:      product,
		ProductId:    product.Id,
		DefaultPrice: domain.Price{Amount: price},
		Position:     &position,
	}

	if err = productRepository.Create(product); err != nil {
		t.Error("An error occured while creating the product:", err.Error())
	}

	if product.Master.Id == 0 {
		t.Error("Master variant was not created")
	}

	if len(product.Master.StockItems) == 0 {
		t.Error("Master should have stock items")
	}
}
