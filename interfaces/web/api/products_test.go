package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductsIndex(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer func() {
		repositories.Spree_db.Rollback()
		repositories.Spree_db.Close()
	}()

	r := gin.New()

	method := "GET"
	path := "/api/products/"

	repositories.Spree_db.Create(&domain.User{})

	r.GET(path, func(c *gin.Context) {
		user := &domain.User{}
		repositories.Spree_db.First(user)
		c.Set("CurrentUser", user)
		ProductsIndex(c)
	})
	w := PerformRequest(r, method, path)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}

func TestProductsShow(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer func() {
		repositories.Spree_db.Rollback()
		repositories.Spree_db.Close()
	}()

	repositories.Spree_db.Create(&domain.Product{Id: 1})

	r := gin.New()

	method := "GET"
	path := "/api/products/1"

	r.GET("/api/products/:product_id", func(c *gin.Context) {
		ProductsShow(c)
	})
	w := PerformRequest(r, method, path)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}
