package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	djson "github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/usecases/json"
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
	order := getGinOrder(c)

	if order == nil {
		order = &models.Order{}
		err := repositories.NewDatabaseRepository().FindBy(order, params{
			"number": c.Params.ByName("order_number"),
		})

		if err != nil {
			notFound(c)
			return
		}

		c.Set("Order", order)
	}

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

	order := getGinOrder(c)
	if order != nil && (*order.UserId == user.Id || order.GuestToken == getOrderToken(c)) {
		c.Next()
	} else {
		unauthorized(c, "You are not authorized to perform that action.")
		return
	}
}

func OrdersIndex(c *gin.Context) {
	orders, err := json.SpreeResponseFetcher.GetResponse(json.NewOrderInteractor(), 1, 0)

	if err == nil {
		c.JSON(200, orders)
	} else {
		c.JSON(422, gin.H{"error": err.Error()})
	}
}

func OrdersShow(c *gin.Context) {
	order := getGinOrder(c)
	orderJson := djson.Order{}

	isAdmin := currentUser(c).HasRole("admin")
	r := repositories.NewDatabaseRepository()

	quantities, _ := r.SumLineItemsQuantityByOrderIds([]int64{order.Id})
	order.Quantity = quantities[order.Id]
	copier.Copy(&orderJson, order)

	orderJson.Permissions = &djson.Permissions{CanUpdate: &isAdmin}

	orderJson.BillAddress = &djson.Address{}
	r.Association(&orderJson, orderJson.BillAddress, "BillAddressId")

	orderJson.BillAddress.Country = &djson.Country{}
	r.Association(orderJson.BillAddress, orderJson.BillAddress.Country, "CountryId")

	orderJson.BillAddress.State = &djson.State{}
	r.Association(orderJson.BillAddress, orderJson.BillAddress.State, "StateId")
	orderJson.BillAddress.StateName = orderJson.BillAddress.State.Name
	orderJson.BillAddress.StateText = orderJson.BillAddress.State.Abbr

	orderJson.ShipAddress = &djson.Address{}
	r.Association(&orderJson, orderJson.ShipAddress, "ShipAddressId")

	orderJson.ShipAddress.Country = &djson.Country{}
	r.Association(orderJson.ShipAddress, orderJson.ShipAddress.Country, "CountryId")

	orderJson.ShipAddress.State = &djson.State{}
	r.Association(orderJson.ShipAddress, orderJson.ShipAddress.State, "StateId")
	orderJson.ShipAddress.StateName = orderJson.ShipAddress.State.Name
	orderJson.ShipAddress.StateText = orderJson.ShipAddress.State.Abbr

	orderJson.LineItems = &[]djson.LineItem{}
	r.Association(&orderJson, orderJson.LineItems, "OrderId")

	c.JSON(200, orderJson)
}

func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"error": "Record Not Found"})
	c.Abort()
}
