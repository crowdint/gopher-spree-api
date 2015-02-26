package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	taxons := API().Group("/taxons")
	{
		taxons.GET("", TaxonsIndex)
		taxons.GET("/", TaxonsIndex)
	}
}

func TaxonsIndex(c *gin.Context) {
	params := NewRequestParameters(c)

	if taxons, err := domain.SpreeResponseFetcher.GetResponse(domain.NewTaxonInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, taxons)
	}
}
