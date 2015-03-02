package repositories

import (
	"strconv"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestOrderRepository(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	orderRepository := NewOrderRepository()

	order := &domain.Order{}

	Spree_db.Create(order)
	Spree_db.Exec("INSERT INTO spree_line_items(order_id, quantity, price) values(" + strconv.Itoa(int(order.Id)) + ", 1, 10)")

	orderRepository.dbHandler.First(order)

	quantities, err := orderRepository.SumLineItemsQuantityByOrderIds([]int64{order.Id})
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if quantities[order.Id] < 1 {
		t.Error("Quantity should be greater than 0, but it wasn't")
	}
}
