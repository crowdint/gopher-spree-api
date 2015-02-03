package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		spreeToken := getSpreeToken(c)

		// Return if spree token is not provided
		if spreeToken == "" {
			unauthorized(c, "You must specify an API key.")
			return
		}

		user := &models.User{}
		err := repositories.NewDatabaseRepository().FindBy(user, map[string]interface{}{"spree_api_key": spreeToken})

		if err != nil {
			unauthorized(c, "Invalid API key ("+spreeToken+") specified.")
			return
		}

		repositories.NewDatabaseRepository().UserRoles(user)

		c.Set(SPREE_TOKEN, spreeToken)
		c.Set("CurrentUser", user)
		c.Next()
	}
}
