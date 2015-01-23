package endpoints

import (
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.Engine) {
	api := router.Group("/api")
	mountProducts(api)
}
