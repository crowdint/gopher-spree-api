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

		router.Use(Monitor(), Proxy(), Authentication())
	}

	return router
}

func routes() map[string]bool {
	return map[string]bool{
		`^` + namespace() + `/api/products(/?)$`:   true,
		`^` + namespace() + `/api/products/\d+$`:   true,
		`^` + namespace() + `/api/orders(/?)$`:     false,
		`^` + namespace() + `/api/orders/\w+$`:     false,
		`^` + namespace() + `/api/taxonomies(/?)$`: false,
		`^` + namespace() + `/api/taxons(/?)$`:     false,
	}
}
