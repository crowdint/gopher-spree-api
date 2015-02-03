package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func init() {
	orders := API().Group("/orders")
	{
		orders.GET("", authorizeOrders, OrdersIndex)
		orders.GET("/", OrdersIndex)
		orders.GET("/:order_number", findOrder, authorizeOrder, OrdersShow)
	}
}

func findOrder(c *gin.Context) {
	var order models.Order

	err := repositories.NewDatabaseRepository().FindBy(&order, params{
		"number": c.Params.ByName("order_number"),
	})

	if err != nil {
		notFound(c)
		return
	}

	c.Set("Order", &order)
	c.Next()
}

func getGinOrder(c *gin.Context) *models.Order {
	rawOrder, err := c.Get("Order")
	if err == nil {
		return rawOrder.(*models.Order)
	}
	return nil
}

func authorizeOrders(c *gin.Context) {
	user := currentUser(c)
	if !user.HasRole("admin") {
		unauthorized(c, "You are not authorized to perfomr that action.")
		return
	}

	c.Next()
}

func authorizeOrder(c *gin.Context) {
	user := currentUser(c)
	if user.HasRole("admin") {
		c.Next()
		return
	}

	order := getGinOrder(c)
	if order != nil && (order.UserId == user.Id || order.GuestToken == getOrderToken(c)) {
		c.Next()
	} else {
		unauthorized(c, "You are not authorized to perform that action.")
		return
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
	order, _ := c.Get("Order")
	c.JSON(200, order.(*models.Order))
}

func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"error": "Record Not Found"})
	c.Abort(-1)
}
