package api

import (
	"os"

	"github.com/brandfolder/gin-gorelic"
	"github.com/gin-gonic/gin"
)

func Monitor() gin.HandlerFunc {
	apiKey, appName := os.Getenv("NEWRELIC_API_KEY"), os.Getenv("NEWRELIC_APP_NAME")

	if apiKey != "" && appName != "" {
		gorelic.InitNewrelicAgent(apiKey, appName, true)
		return gorelic.Handler
	}

	return func(c *gin.Context) {}
}
