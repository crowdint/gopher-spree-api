package api

import (
	"github.com/gin-gonic/gin"
)

func ProductsIndex(c *gin.Context) {
	c.JSON(200, []struct{}{})
}

func ProductsShow(c *gin.Context) {
	c.JSON(200, struct{}{})
}
