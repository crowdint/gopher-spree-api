package main

import (
	"os"

	_ "github.com/crowdint/gopher-spree-api/cache"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/interfaces/web/api"
	"github.com/crowdint/gopher-spree-api/utils"
)

func main() {
	err := repositories.InitDB(false)

	if err != nil {
		panic(err)
	}
	utils.LogrusInfo(utils.FuncName(), "Listening and serving HTTP on 0.0.0.0:"+os.Getenv("PORT"))
	api.Router().Run("0.0.0.0:" + os.Getenv("PORT"))

}
