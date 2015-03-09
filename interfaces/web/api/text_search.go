package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	products := API().Group("/search")
	{
		products.GET("/products", ProductsTextSearch)
	}
}

func ProductsTextSearch(c *gin.Context) {
	params := NewRequestParameters(c, json.ELASTIC_SEARCH)

	productShowResponse(c, params)
}
