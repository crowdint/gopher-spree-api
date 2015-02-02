package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

var (
	permissions = map[string]interface{}{
		"products.index": func(user *models.User, args ...interface{}) bool { return true },
		"products.show":  func(user *models.User, args ...interface{}) bool { return true },
	}
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user from Authentication() middleware.
		rawUser, _ := c.Get("CurrentUser")
		// spreeToken, _ := c.Get(SPREE_TOKEN)
		currentUser := rawUser.(*models.User)

		// Get current user's roles.
		repositories.NewDatabaseRepository().UserRoles(currentUser)
		for _, role := range currentUser.Roles {
			if role.Name == "admin" {
				c.Next()
				return
			}
		}

		// Get current action (products.index, products.show, etc).
		//currentAction := action(c.Request.URL)

		//// Check if current user has permissions to perform the action.
		//if !hasPermission(currentUser, currentAction, spreeToken) {
		//unauthorized(c, "You have no permissions to perform this action")
		//return
		//}

		// Temporary, right now we only have read product endopoints
		c.Next()
	}
}

func hasPermission(user *models.User, action string, args ...interface{}) bool {
	if permissionFunc := permissions[action]; permissionFunc != nil {
		return permissionFunc.(func(*models.User, ...interface{}) bool)(user, args...)
	}
	return false
}
