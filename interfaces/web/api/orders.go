package api

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/usecases/json"
	"github.com/crowdint/gopher-spree-api/utils"
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
		order = &domain.Order{}
		err := repositories.NewDatabaseRepository().FindBy(order, nil, params{
			"number": c.Params.ByName("order_number"),
		})

		if err != nil {
			utils.LogrusError("findOrder", err)
			fail(c, err)
			return
		}

		c.Set("Order", order)
	}

	c.Next()
}

func currentOrder(c *gin.Context) *domain.Order {
	order, err := c.Get("Order")
	if err == nil {
		return order.(*domain.Order)
	}
	utils.LogrusError("currentOrder", err)
	return nil
}

func authorizeOrders(c *gin.Context) {
	user := currentUser(c)
	if !user.HasRole("admin") {
		msg := "You are not authorized to perform that action."
		err := errors.New(msg)
		utils.LogrusWarning("authorizeOrders", err)
		unauthorized(c, msg)
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
		msg := "You are not authorized to perform that action."
		err := errors.New(msg)
		utils.LogrusWarning("authorizeOrder", err)
		unauthorized(c, msg)
		return
	}
}

func OrdersIndex(c *gin.Context) {
	params := NewRequestParameters(c, json.GRANSAK)
	orders, err := json.SpreeResponseFetcher.GetResponse(json.NewOrderInteractor(), params)

	if err == nil || len(orders) == 0 {
		c.JSON(200, orders)
	} else {
		utils.LogrusError("OrdersIndex", err)
		c.JSON(422, gin.H{"error": err.Error()})
	}
}

func OrdersShow(c *gin.Context) {
	order, err := json.NewOrderInteractor().Show(currentOrder(c), currentUser(c))

	if err == nil {
		c.JSON(200, order)
	} else {
		utils.LogrusError("OrdersShow", err)
		internalServerError(c, err.Error())
	}
}
