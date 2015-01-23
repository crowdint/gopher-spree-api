package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/interfaces/api/endpoints"
)

func main() {
	router := gin.Default()

	endpoints.Mount(router)

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
