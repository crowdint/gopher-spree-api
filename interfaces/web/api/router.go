package api

import (
	"os"

	"github.com/brandfolder/gin-gorelic"
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
		gorelic.InitNewrelicAgent(os.Getenv("NEWRELIC_API_KEY"), os.Getenv("NEWRELIC_APP_NAME"), true)
		router.Use(Proxy(), Authentication(), gorelic.Handler)
	}

	return router
}

func regexRoutesPattern() map[string]string {
	return map[string]string{
		`^` + namespace() + `/api/products(/?)$`: "products.index", // pattern -> action
		`^` + namespace() + `/api/products/\d+$`: "products.show",
		`^` + namespace() + `/api/orders(/?)$`:   "orders.index",
		`^` + namespace() + `/api/orders/\w+$`:   "orders.show",
	}
}
