package json

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/crowdint/gopher-spree-api/cache"
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestOrderInteractor_ToCacheData(t *testing.T) {
	orderInteractor := NewOrderInteractor()
	orderSlice := []*domain.Order{
		&domain.Order{
			Id:    100,
			Email: "test@email.com",
		},
	}

	cacheSlice := orderInteractor.toCacheData(orderSlice)
	if len(cacheSlice) != len(orderSlice) {
		t.Fatalf("The len of cache Slice should be %d, but was %d", len(orderSlice), len(cacheSlice))
	}
}

func TestOrderInteractor_GetResponse(t *testing.T) {
	if err := cache.SetupMemcached(); err != nil {
		t.Error("Couldn't find memcached")
	}

	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()
	defer cache.KillMemcached()

	order := &domain.Order{}

	repositories.Spree_db.Create(order)
	repositories.Spree_db.Exec("INSERT INTO spree_line_items(order_id, quantity, price) values(" + strconv.Itoa(int(order.Id)) + ", 1, 10)")

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

	ordersCached := []cache.Cacheable{
		order,
	}

	_, err = cache.FetchMultiWithPrefix("index", ordersCached)
	if err != nil {
		t.Error("An error ocurred while finding cached orders: ", err.Error())
	}
}

func TestOrderInteractor_Show(t *testing.T) {
	if err := cache.SetupMemcached(); err != nil {
		t.Error("Couldn't find memcached")
	}

	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()
	defer cache.KillMemcached()

	oid := int64(1)

	order := domain.Order{
		Id:     379,
		UserId: &oid,
	}

	repositories.Spree_db.Create(&order)
	repositories.Spree_db.Exec("INSERT INTO spree_users(id, deleted_at) VALUES(1, null)")

	repositories.Spree_db.Create(&domain.LineItem{Id: 1, OrderId: 379, VariantId: 1, Quantity: 1})
	repositories.Spree_db.Create(&domain.LineItem{Id: 2, OrderId: 379, VariantId: 2, Quantity: 1})

	repositories.Spree_db.Create(&domain.Variant{Id: 1, ProductId: 1, CostPrice: "10"})
	repositories.Spree_db.Create(&domain.Variant{Id: 2, ProductId: 2, CostPrice: "10"})

	repositories.Spree_db.Exec("INSERT INTO spree_stock_items(variant_id) values(1)")
	repositories.Spree_db.Exec("INSERT INTO spree_prices(variant_id, currency) values(1, 'USD')")

	repositories.Spree_db.Exec("INSERT INTO spree_stock_items(variant_id) values(2)")
	repositories.Spree_db.Exec("INSERT INTO spree_prices(variant_id, currency) values(2, 'USD')")

	tmpl := `INSERT INTO spree_products(id, name, description, available_on, deleted_at, slug, meta_description, meta_keywords, tax_category_id, shipping_category_id, created_at, updated_at, promotionable, meta_title) VALUES(%s)`

	sql1 := fmt.Sprintf(tmpl, `1,'Spree Ringer T-Shirt','Labore ut sint neque exercitationem aliquid consequuntur ea dolores.Quo asperiores eligendi ipsam officia.Autem aliquid temporibus est blanditiis','2015-02-24 17:57:13.788353',null,'spree-ringer-t-shirt',null,null,1,1,'2015-02-24 17:57:15.214292','2015-02-24 17:57:39.946429','t',null`)
	sql2 := fmt.Sprintf(tmpl, `2, 'Ruby on Rails Mug','Labore ut sint neque exercitationem aliquid consequuntur ea dolores.Quo asperiores eligendi ipsam officia.Autem aliquid temporibus est blanditiis.','2015-02-24 17:57:13.788353',null,'ruby-on-rails-mug',null,null,null,1,'2015-02-24 17:57:15.518985','2015-02-24 17:57:33.982174','t',null`)

	repositories.Spree_db.Exec(sql1)
	repositories.Spree_db.Exec(sql2)

	repositories.Spree_db.Create(&domain.StockItem{Id: 1, VariantId: 1, StockLocationId: 1})
	repositories.Spree_db.Create(&domain.StockItem{Id: 2, VariantId: 2, StockLocationId: 1})

	orderInteractor := NewOrderInteractor()
	user := domain.User{}

	err := orderInteractor.OrderRepository.FindBy(&order, map[string]interface{}{
		"not": repositories.Not{Key: "user_id", Values: []interface{}{0}},
	}, nil)
	if err != nil {
		t.Error("Error: An error has ocurred while getting an order:", err.Error())
		return
	}

	err = orderInteractor.OrderRepository.FindBy(&user, nil, map[string]interface{}{
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

	if err = cache.Find(&order); err != nil {
		t.Error("Order should be cached, but it wasn't:", err.Error())
	}
}
