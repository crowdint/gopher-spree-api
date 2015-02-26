package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	products := API().Group("/products")
	{
		products.GET("", ProductsIndex)
		products.GET("/", ProductsIndex)
		products.GET("/:product_id", ProductsShow)
	}
}

func ProductsIndex(c *gin.Context) {
	params := NewRequestParameters(c)

	if products, err := json.SpreeResponseFetcher.GetResponse(json.NewProductInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, products)
	}
}

func ProductsShow(c *gin.Context) {
	params := NewRequestParameters(c)

	product, err := json.SpreeResponseFetcher.GetShowResponse(json.NewProductInteractor(), params)

	if err == nil {
		c.JSON(200, product)
		return
	}

	if err.Error() == "Record Not Found" {
		notFound(c)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
