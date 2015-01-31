package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/domain/models"
)

var (
	routesPattern map[string]string
	permissions   = map[string]interface{}{
		"products.index": func(user *models.User, args ...interface{}) bool { return true },
		"products.show":  func(user *models.User, args ...interface{}) bool { return true },
	}
)

func regexRoutesPattern() map[string]string {
	ns := configs.Get(configs.SPREE_NS)
	if ns != "" {
		ns = "/" + ns
	}

	return map[string]string{
		`^` + ns + `/api/products(/?)$`: "products.index", // pattern -> action
		`^` + ns + `/api/products/\d+$`: "products.show",
	}
}

func hasPermission(user *models.User, action string, args ...interface{}) bool {
	if permissionFunc := permissions[action]; permissionFunc != nil {
		return permissionFunc.(func(*models.User, ...interface{}) bool)(user, args...)
	}
	return false
}

func unauthorized(c *gin.Context, errMsg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	c.Abort(-1)
}
