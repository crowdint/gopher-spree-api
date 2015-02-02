package api

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
)

type params map[string]interface{}

func action(url *url.URL) string {
	for pattern, action := range routesPattern {
		if match, _ := regexp.MatchString(pattern, url.Path); match {
			return action
		}
	}

	return ""
}

func unauthorized(c *gin.Context, errMsg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	c.Abort(-1)
}
