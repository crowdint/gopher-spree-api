package repositories

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/json"
)

func TestOrderRepository(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	orderRepository := NewOrderRepository()

	order := &json.Order{}
	orderRepository.dbHandler.First(order)

	quantities, err := orderRepository.SumLineItemsQuantityByOrderIds([]int64{order.Id})
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if quantities[order.Id] < 1 {
		t.Error("Quantity should be greater than 0, but it wasn't")
	}
}
