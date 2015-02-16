package json

import (
	"github.com/jinzhu/copier"

	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type OrderInteractor struct {
	Repository *repositories.DbRepo
}

func (this *OrderInteractor) Show(o *models.Order, u *models.User) (*json.Order, error) {
	order := json.Order{}

	// Copy from order model to order json
	copier.Copy(&order, o)

	order.Permissions = this.getPermissions(&order, u)
	order.Quantity = this.getQuantity(&order)
	order.BillAddress = this.getAddress(&order, "BillAddressId")
	order.ShipAddress = this.getAddress(&order, "ShipAddressId")
	order.LineItems = this.getLineItems(&order)

	return &order, nil
}

func (this *OrderInteractor) getQuantity(order *json.Order) int64 {
	quantities, _ := this.Repository.SumLineItemsQuantityByOrderIds([]int64{order.Id})
	return quantities[order.Id]
}

func (this *OrderInteractor) getPermissions(order *json.Order, user *models.User) *json.Permissions {
	updatePermission := user.HasRole("admin") || (*order.UserId == user.Id)
	permissions := &json.Permissions{CanUpdate: &updatePermission}
	return permissions
}

func (this *OrderInteractor) getLineItems(order *json.Order) *[]json.LineItem {
	lineItems := &[]json.LineItem{}
	this.Repository.Association(order, lineItems, "OrderId")
	return lineItems
}

func (this *OrderInteractor) getAddress(order *json.Order, id string) *json.Address {
	address := &json.Address{}

	this.Repository.Association(order, address, id)

	address.Country = &json.Country{}
	this.Repository.Association(address, address.Country, "CountryId")

	address.State = &json.State{}
	this.Repository.Association(address, address.State, "StateId")
	address.StateName = address.State.Name
	address.StateText = address.State.Abbr

	return address
}

func (this *OrderInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	return nil, nil
}

func (this *OrderInteractor) GetResponse(currentPage, perPage int, query string) (ContentResponse, error) {
	orders := []models.Order{}
	ordersJson := []json.Order{}

	err := this.Repository.All(&orders, repositories.Q{
		"current_page": currentPage,
		"per_page":     perPage,
		"q":            query,
	})

	if err != nil {
		return &OrderResponse{}, err
	}

	var orderIds []int64
	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
	}

	quantities, err := this.Repository.SumLineItemsQuantityByOrderIds(orderIds)
	for index, order := range orders {
		orders[index].Quantity = quantities[order.Id]
	}

	copier.Copy(&ordersJson, &orders)

	return &OrderResponse{data: &ordersJson}, nil
}

func (this *OrderInteractor) GetTotalCount(query string) (int64, error) {
	return this.Repository.Count(&models.Order{})
}

func NewOrderInteractor() *OrderInteractor {
	return &OrderInteractor{
		Repository: repositories.NewDatabaseRepository(),
	}
}

type OrderResponse struct {
	data *[]json.Order
}

func (this *OrderResponse) GetCount() int {
	return len(*this.data)
}

func (this *OrderResponse) GetData() interface{} {
	return this.data
}

func (this OrderResponse) GetTag() string {
	return "orders"
}
