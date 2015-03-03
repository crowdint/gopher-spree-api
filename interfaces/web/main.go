package main

import (
	"os"

	_ "github.com/crowdint/gopher-spree-api/cache"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/interfaces/web/api"
)

func main() {
	err := repositories.InitDB(false)

	if err != nil {
		panic(err)
	}

	api.Router().Run("0.0.0.0:" + os.Getenv("PORT"))
}
