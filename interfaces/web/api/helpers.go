package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
)

const (
	SPREE_TOKEN_HEADER       = "X-Spree-Token"
	SPREE_TOKEN              = "SpreeToken"
	SPREE_ORDER_TOKEN_HEADER = "X-Spree-Order-Token"
)

type params map[string]interface{}

func currentUser(c *gin.Context) *models.User {
	currentUser, _ := c.Get("CurrentUser")
	user := currentUser.(*models.User)
	return user
}

func getOrderToken(c *gin.Context) string {
	orderToken := c.Request.Header.Get(SPREE_ORDER_TOKEN_HEADER)

	if orderToken != "" {
		return orderToken
	}

	return c.Request.URL.Query().Get("order_token")
}

func getSpreeToken(c *gin.Context) string {
	spreeToken := c.Request.Header.Get(SPREE_TOKEN_HEADER)

	if len(spreeToken) > 0 {
		return spreeToken
	}

	return c.Request.URL.Query().Get("token")
}

func unauthorized(c *gin.Context, errMsg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	c.Abort(-1)
}
