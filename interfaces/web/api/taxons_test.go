package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestTaxonsIndex(t *testing.T) {
	r := gin.New()

	method := "GET"
	path := "/api/taxons/"

	r.GET(path, func(c *gin.Context) {
		user := &json.User{}
		repositories.Spree_db.First(user)
		c.Set("CurrentUser", user)
		TaxonsIndex(c)
	})
	w := PerformRequest(r, method, path)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}
