package api

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(200, []struct{}{})
}

func ProductsShow(c *gin.Context) {
	c.JSON(200, struct{}{})
}
