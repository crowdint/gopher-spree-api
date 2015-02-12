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

func (this *OrderInteractor) GetResponse(currentPage, perPage int, query string) (ContentResponse, error) {
	orders := []models.Order{}
	ordersJson := []json.Order{}

	err := this.Repository.All(&orders, repositories.Query{
		"current_page": currentPage,
		"per_page":     perPage,
		"condition":    query,
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
