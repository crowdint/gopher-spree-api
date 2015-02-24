package json

import (
	"encoding/json"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestOrderInteractor_GetResponse(t *testing.T) {
	err := repositories.InitDB(true)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	defer repositories.Spree_db.Close()

	orderInteractor := NewOrderInteractor()

	jsonOrderSlice, err := orderInteractor.GetResponse(1, 10, &FakeResponseParameters{})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	if jsonOrderSlice.(ContentResponse).GetCount() < 1 {
		t.Error("Error: Invalid number of rows")
		return
	}

	jsonBytes, err := json.Marshal(jsonOrderSlice)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	if string(jsonBytes) == "" {
		t.Error("Error: Json string is empty")
		return
	}
}

func TestOrderInteractor_Show(t *testing.T) {
	err := repositories.InitDB(true)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	defer repositories.Spree_db.Close()

	orderInteractor := NewOrderInteractor()
	order := domain.Order{}
	user := domain.User{}

	err = orderInteractor.BaseRepository.FindBy(&order, map[string]interface{}{
		"not": repositories.Not{Key: "user_id", Values: []interface{}{0}},
	}, nil)
	if err != nil {
		t.Error("Error: An error has ocurred while getting an order:", err.Error())
		return
	}

	err = orderInteractor.BaseRepository.FindBy(&user, nil, map[string]interface{}{
		"id": order.UserId,
	})
	if err != nil {
		t.Errorf("Error: An error has ocurred while getting the user with id %d, : %s", order.UserId, err.Error())
		return
	}

	jsonOrder, err := orderInteractor.Show(&order, &user)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	if jsonOrder.Permissions == nil {
		t.Error("Order Permissions should not be nil, but it was")
	}

	if jsonOrder.Quantity < 1 {
		t.Error("Order Quantity should be greater than 0")
	}

	if jsonOrder.LineItems == nil {
		t.Error("Order LineItems should not be nil, but it was")
	}
}
