package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
)

var (
	api           *gin.RouterGroup
	router        *gin.Engine
	routesPattern map[string]string
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

func regexRoutesPattern() map[string]string {
	ns := configs.Get(configs.SPREE_NS)
	if ns != "" {
		// If namespace has '/', then remove them
		ns = "/" + strings.Replace(ns, "/", "", -1)
	}

	return map[string]string{
		`^` + ns + `/api/products(/?)$`: "products.index", // pattern -> action
		`^` + ns + `/api/products/\d+$`: "products.show",
		`^` + ns + `/api/orders(/?)$`:   "orders.index",
		`^` + ns + `/api/orders/\w+$`:   "orders.show",
	}
}
