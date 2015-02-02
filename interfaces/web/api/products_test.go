package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductsIndex(t *testing.T) {
	r := gin.New()

	method := "GET"
	path := "/api/products/"

	r.GET(path, func(c *gin.Context) {
		user := &models.User{}
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
	r := gin.New()

	method := "GET"
	path := "/api/products/1"

	r.GET("/api/products/:id", func(c *gin.Context) {
		ProductsShow(c)
	})
	w := PerformRequest(r, method, path)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}

	bodyExpected := "{}"
	if NotEqualFromJSONString(&bodyExpected, w.Body.String()) {
		t.Errorf("Body should be %s, but was %s", bodyExpected, w.Body.String())
	}
}
