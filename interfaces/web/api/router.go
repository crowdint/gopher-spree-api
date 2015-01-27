package api

import (
	"os"

	"github.com/gin-gonic/gin"
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

		namespace := os.Getenv("SPREE_NS")

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
	}

	return router
}
