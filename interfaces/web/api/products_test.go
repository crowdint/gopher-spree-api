package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestProductsIndex(t *testing.T) {
	t.Skip("Update test after product index implementation")
	r := gin.New()

	method := "GET"
	path := "/api/products/"

	r.GET(path, func(c *gin.Context) {
		ProductsIndex(c)
	})
	w := PerformRequest(r, method, path)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}

	bodyExpected := "[]"
	if NotEqualFromJSONString(&bodyExpected, w.Body.String()) {
		t.Errorf("Body should be %s, but was %s", bodyExpected, w.Body.String())
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
