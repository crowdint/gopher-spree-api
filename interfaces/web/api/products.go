package api

import (
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
	products, err := json.NewResponseInteractor().GetResponse()

	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, products)
	}
}

func ProductsShow(c *gin.Context) {
	c.JSON(200, struct{}{})
}
