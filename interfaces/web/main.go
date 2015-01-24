package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/interfaces/web/api"
)

func main() {
	router := gin.Default()

	router.Use(api.Router())

	a := router.Group("/api")

	products := a.Group("/products")
	{
		products.GET("", api.ProductsIndex)
		products.GET("/", api.ProductsIndex)
		products.GET("/:product_id", api.ProductsShow)
	}

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
