package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SPREE_TOKEN_HEADER = "X-Spree-Token"
	SPREE_TOKEN        = "SpreeToken"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if proxied, _ := c.Get("Proxied"); proxied == true {
			c.Next()
			return
		}

		if spreeToken := getSpreeToken(c); len(spreeToken) > 0 {
			c.Set(SPREE_TOKEN, spreeToken)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You must specify an API key."})
			c.Abort(-1) // If abort index is lower than 0 header is not written
		}
	}
}

func getSpreeToken(c *gin.Context) string {
	spreeToken := c.Request.Header.Get(SPREE_TOKEN_HEADER)

	if len(spreeToken) > 0 {
		return spreeToken
	}

	return c.Request.URL.Query().Get("token")
}
