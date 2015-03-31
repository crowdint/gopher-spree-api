package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
	"github.com/crowdint/gopher-spree-api/utils"
)

func init() {
	taxons := API().Group("/taxons")
	{
		taxons.GET("", TaxonsIndex)
		taxons.GET("/", TaxonsIndex)
	}
}

func TaxonsIndex(c *gin.Context) {
	params := NewRequestParameters(c, json.GRANSAK)

	if taxons, err := json.SpreeResponseFetcher.GetResponse(json.NewTaxonInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		utils.LogrusError(utils.FuncName(), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, taxons)
	}
}
