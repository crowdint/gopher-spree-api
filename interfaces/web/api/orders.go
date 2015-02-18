package api

import (
	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	orders := API().Group("/orders")
	{
		orders.GET("", authorizeOrders, OrdersIndex)
		orders.GET("/", authorizeOrders, OrdersIndex)
		orders.GET("/:order_number", findOrder, authorizeOrder, OrdersShow)
	}
}

func findOrder(c *gin.Context) {
	order := currentOrder(c)

	if order == nil {
		order = &models.Order{}
		err := repositories.NewDatabaseRepository().FindBy(order, params{
			"number": c.Params.ByName("order_number"),
		})

		if err != nil {
			fail(c, err)
			return
		}

		c.Set("Order", order)
	}

	c.Next()
}

func currentOrder(c *gin.Context) *models.Order {
	order, err := c.Get("Order")
	if err == nil {
		return order.(*models.Order)
	}
	return nil
}

func authorizeOrders(c *gin.Context) {
	user := currentUser(c)
	if !user.HasRole("admin") {
		unauthorized(c, "You are not authorized to perform that action.")
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

	order := currentOrder(c)
	if order != nil && (*order.UserId == user.Id || order.GuestToken == getOrderToken(c)) {
		c.Next()
	} else {
		unauthorized(c, "You are not authorized to perform that action.")
		return
	}
}

func OrdersIndex(c *gin.Context) {
	params := NewRequestParameters(c)
	orders, err := json.SpreeResponseFetcher.GetResponse(json.NewOrderInteractor(), params)

	if err == nil {
		c.JSON(200, orders)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}

func OrdersShow(c *gin.Context) {
	order, err := json.NewOrderInteractor().Show(currentOrder(c), currentUser(c))

	if err == nil {
		c.JSON(200, order)
	} else {
		internalServerError(c, err.Error())
	}
}
