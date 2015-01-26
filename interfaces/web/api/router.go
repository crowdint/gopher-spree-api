package api

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	Router()
}

func Router() (router *gin.Engine) {
	if router == nil {
		router = gin.Default()
	}
	return
}
