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

func (this *OrderInteractor) GetResponse(currentPage, perPage int) (ContentResponse, error) {
	orders := []models.Order{}
	ordersJson := []json.Order{}

	err := this.Repository.All(&orders, map[string]interface{}{
		"current_page": currentPage,
		"per_page":     perPage,
	})

	if err != nil {
		return &OrderResponse{}, err
	}

	copier.Copy(&ordersJson, &orders)

	return &OrderResponse{data: &ordersJson}, nil
}

func (this *OrderInteractor) GetTotalCount() (int64, error) {
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
