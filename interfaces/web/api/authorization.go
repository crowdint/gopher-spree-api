package api

import (
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user from Authentication() middleware.
		rawUser, _ := c.Get("CurrentUser")
		spreeToken, _ := c.Get(SPREE_TOKEN)
		currentUser := rawUser.(*models.User)

		// Get current user's roles.
		role := getCurrentUserRole(currentUser)
		if role.Name == "admin" {
			c.Next()
			return
		}

		// Get current action (products.index, products.show, etc).
		//currentAction := getCurrentAction(c.Request.URL)

		//// Check if current user has permissions to perform the action.
		//if !hasPermission(currentUser, currentAction, spreeToken) {
		//unauthorized(c, "You have no permissions to perform this action")
		//return
		//}

		// Temporary, right now we only have read product endopoints
		c.Next()
	}
}

func getCurrentUserRole(currentUser *models.User) *models.Role {
	repositories.NewDatabaseRepository().UserRoles(currentUser)
	return &currentUser.Roles[0]
}

func getCurrentAction(url *url.URL) string {
	for pattern, action := range routesPattern {
		if match, _ := regexp.MatchString(pattern, url.Path); match {
			return action
		}
	}

	return ""
}
