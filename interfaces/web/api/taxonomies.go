package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	taxonomies := API().Group("/taxonomies")
	{
		taxonomies.GET("", TaxonomiesIndex)
		taxonomies.GET("/", TaxonomiesIndex)
	}
}

func TaxonomiesIndex(c *gin.Context) {
	params := NewRequestParameters(c)

	if taxonomies, err := domain.SpreeResponseFetcher.GetResponse(domain.NewTaxonomyInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, taxonomies)
	}
}
