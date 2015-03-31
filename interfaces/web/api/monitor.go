package api

import (
	"os"

	"github.com/brandfolder/gin-gorelic"
	"github.com/crowdint/gopher-spree-api/utils"
	"github.com/gin-gonic/gin"
)

func Monitor() gin.HandlerFunc {
	apiKey, appName := os.Getenv("NEWRELIC_API_KEY"), os.Getenv("NEWRELIC_APP_NAME")

	if apiKey != "" && appName != "" {
		utils.LogrusInfo(utils.FuncName(), "New relic api key: "+apiKey)
		utils.LogrusInfo(utils.FuncName(), "New relic app name: "+appName)

		gorelic.InitNewrelicAgent(apiKey, appName, true)
		utils.LogrusInfo(utils.FuncName(), "Create the new relic agent")

		return gorelic.Handler
	}

	return func(c *gin.Context) {}
}
