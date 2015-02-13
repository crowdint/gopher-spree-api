package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	taxonomies := API().Group("/taxonomies")
	{
		taxonomies.GET("", taxonomiesIndex)
		taxonomies.GET("/", taxonomiesIndex)
	}
}

func taxonomiesIndex(c *gin.Context) {
	params := NewRequestParameters(c)

	if taxonomies, err := json.SpreeResponseFetcher.GetResponse(json.NewTaxonomyInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, taxonomies)
	}
}

