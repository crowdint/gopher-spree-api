package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
	"github.com/crowdint/gopher-spree-api/utils"
)

func init() {
	taxonomies := API().Group("/taxonomies")
	{
		taxonomies.GET("", TaxonomiesIndex)
		taxonomies.GET("/", TaxonomiesIndex)
	}
}

func TaxonomiesIndex(c *gin.Context) {
	params := NewRequestParameters(c, json.GRANSAK)

	if taxonomies, err := json.SpreeResponseFetcher.GetResponse(json.NewTaxonomyInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		utils.LogrusError("TaxonomiesIndex", "GET", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, taxonomies)
	}
}
