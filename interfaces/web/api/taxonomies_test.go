package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestTaxonomiesIndex(t *testing.T) {
	err := repositories.InitDB(true)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	defer repositories.Spree_db.Close()

	r := gin.New()

	method := "GET"
	path := "/api/taxonomies/"

	r.GET(path, func(c *gin.Context) {
		user := &models.User{}
		repositories.Spree_db.First(user)
		c.Set("CurrentUser", user)
		TaxonomiesIndex(c)
	})
	w := PerformRequest(r, method, path)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}
