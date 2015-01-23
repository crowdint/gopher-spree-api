package endpoints

import (
	"github.com/gin-gonic/gin"
)

func mountProducts(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", productsIndex)
		products.GET("/:id", productsShow)
	}
}

func productsIndex(c *gin.Context) {
	c.JSON(200, []struct{}{})
}

func productsShow(c *gin.Context) {
	c.JSON(200, struct{}{})
}
