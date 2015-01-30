package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

const (
	SPREE_TOKEN_HEADER = "X-Spree-Token"
	SPREE_TOKEN        = "SpreeToken"
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

		c.Set(SPREE_TOKEN, spreeToken)
		c.Set("CurrentUser", user)
		c.Next()
	}
}

func getSpreeToken(c *gin.Context) string {
	spreeToken := c.Request.Header.Get(SPREE_TOKEN_HEADER)

	if len(spreeToken) > 0 {
		return spreeToken
	}

	return c.Request.URL.Query().Get("token")
}
