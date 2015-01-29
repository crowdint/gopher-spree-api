package main

import (
	"os"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/interfaces/web/api"
)

func main() {
	repositories.InitDB()
	api.Router().Use(api.Proxy())
	api.Router().Run("0.0.0.0:" + os.Getenv("PORT"))
}
