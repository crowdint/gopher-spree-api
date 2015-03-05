package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
)

var (
	api    *gin.RouterGroup
	router *gin.Engine

	declaredRoutes = []Route{
		// Products
		Route{"GET", `^` + namespace() + `/api/products(/?)$`, true},   // Index
		Route{"GET", `^` + namespace() + `/api/products/\d+$`, true},   // Show
		Route{"POST", `^` + namespace() + `/api/products(/?)$`, false}, // Create

		// Orders
		Route{"GET", `^` + namespace() + `/api/orders(/?)$`, false}, // Index
		Route{"GET", `^` + namespace() + `/api/orders/\w+$`, false}, // Show

		// Taxonomies
		Route{"GET", `^` + namespace() + `/api/taxonomies(/?)$`, false}, // Index

		// Taxons
		Route{"GET", `^` + namespace() + `/api/taxons(/?)$`, false}, // Index

		// Search
		Route{"GET", `^` + namespace() + `/api/search/products(/?)$`, false}, // Search
	}
)

type Route struct {
	Method       string
	RegexPattern string
	IsPublic     bool
}

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

func routes() []Route {
	return declaredRoutes
}
