package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func init() {
	orders := API().Group("/orders")
	{
		orders.GET("", OrdersIndex)
		orders.GET("/", OrdersIndex)
		orders.GET("/:order_number", OrdersShow)
	}
}

func OrdersIndex(c *gin.Context) {
	var orders []models.Order

	err := repositories.NewDatabaseRepository().All(&orders, nil)

	if err == nil {
		c.JSON(200, orders)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}

func OrdersShow(c *gin.Context) {
	var order models.Order

	err := repositories.NewDatabaseRepository().FindBy(&order, params{
		"number": c.Params.ByName("order_number"),
	})

	if err == nil {
		c.JSON(200, order)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}
