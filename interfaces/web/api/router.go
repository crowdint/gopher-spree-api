package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
)

var (
	api    *gin.RouterGroup
	router *gin.Engine
)

func init() {
	Router()
}

func API() *gin.RouterGroup {

	if api == nil {
		r := Router()

		path := "/api"

		namespace := configs.Get(configs.SPREE_NS)

		if namespace != "" {
			path = "/" + namespace + path
		}

		api = r.Group(path)
	}

	return api
}

func Router() *gin.Engine {

	if router == nil {
		router = gin.Default()
		router.Use(Proxy(), Authentication())
	}

	return router
}
