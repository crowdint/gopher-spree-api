package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestFindOrderWhenOrderIsInContext(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Order{Number: "ABC123", GuestToken: "Xrz5qBnbnoBQnYQYzOMQkQ"})

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders/ABC123"

	order := &domain.Order{Number: "ABC123"}
	r.GET(path, func(c *gin.Context) {
		c.Set("Order", order)
		findOrder(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 200 {
		t.Errorf("Status code should be 200, but was %d", w.Code)
	}
}

func TestFindOrderWhenOrderExists(t *testing.T) {
	t.Skip()

	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Order{Number: "R123456789", GuestToken: "Xrz5qBnbnoBQnYQYzOMQkQ"})

	order := &domain.Order{}
	err := repositories.NewDatabaseRepository().FindBy(order, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders/" + order.Number

	r.GET(path, func(c *gin.Context) {
		p := httprouter.Param{Key: "order_number", Value: order.Number}
		c.Params = append(c.Params, p)
		findOrder(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 200 {
		t.Errorf("Status code should be 200, but was %d", w.Code)
	}

	if _, err := ctx.Get("Order"); err != nil {
		t.Errorf("An error occured while setting the order %s", err.Error())
	}
}

func TestFindOrderWhenOrderDoesNotExist(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Order{Number: "R123456789", GuestToken: "Xrz5qBnbnoBQnYQYzOMQkQ"})

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders/testOrderNumber"

	r.GET(path, func(c *gin.Context) {
		p := httprouter.Param{Key: "order_number", Value: "testOrderNumber"}
		c.Params = append(c.Params, p)
		findOrder(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 404 {
		t.Errorf("Status code should be 404, but was %d", w.Code)
	}

	if _, err := ctx.Get("Order"); err == nil {
		t.Error("Order should not be set, it was")
	}
}

func TestGetGinOrderWhenOrderIsInContext(t *testing.T) {
	req, err := http.NewRequest("GET", "/orders/orderNumber", nil)
	if err != nil {
		t.Errorf("An error occured: %s", err.Error())
	}

	ctx := &gin.Context{Request: req}
	ctx.Set("Order", &domain.Order{})
	order := currentOrder(ctx)

	if order == nil {
		t.Error("Order should not be nil, but it was")
	}
}

func TestGinOrderWhenOrderIsNotInContext(t *testing.T) {
	req, err := http.NewRequest("GET", "/orders/orderNumber", nil)
	if err != nil {
		t.Errorf("An error occured: %s", err.Error())
	}

	ctx := &gin.Context{Request: req}
	order := currentOrder(ctx)

	if order != nil {
		t.Errorf("Order should be nil, but it was %v", order)
	}
}

func TestAuthorizeOrdersWhenUserIsSetAndIsAdmin(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Order{Number: "R123456789", GuestToken: "Xrz5qBnbnoBQnYQYzOMQkQ"})

	user := &domain.User{}
	user.Roles = []domain.Role{
		domain.Role{Name: "admin"},
	}

	repositories.Spree_db.Create(user)

	err := repositories.NewDatabaseRepository().FindBy(user, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders"

	r.GET(path, func(c *gin.Context) {
		c.Set("CurrentUser", user)
		authorizeOrders(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 200 {
		t.Errorf("Status code should be 200, but was %d", w.Code)
	}
}

func TestAuthorizeOrdersWhenUserIsSetAndIsNotAdmin(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	user := &domain.User{}
	repositories.Spree_db.Create(user)

	err := repositories.NewDatabaseRepository().FindBy(user, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders"

	r.GET(path, func(c *gin.Context) {
		c.Set("CurrentUser", user)
		authorizeOrders(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 401 {
		t.Errorf("Status code should be 401, but was %d", w.Code)
	}
}

func TestAuthorizeOrderWhenUserIsSetAndIsAdmin(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	user := &domain.User{}
	user.Roles = []domain.Role{
		domain.Role{Name: "admin"},
	}

	repositories.Spree_db.Create(user)

	err := repositories.NewDatabaseRepository().FindBy(user, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders/testOrderNumber"

	r.GET(path, func(c *gin.Context) {
		p := httprouter.Param{Key: "order_number", Value: "testOrderNumber"}
		c.Params = append(c.Params, p)
		c.Set("CurrentUser", user)
		authorizeOrder(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 200 {
		t.Errorf("Status code should be 200, but was %d", w.Code)
	}
}

func TestAuthorizeOrderWhenUserIsNotAdminAndOrderBelongsToHim(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	dbRepo := repositories.NewDatabaseRepository()

	user := &domain.User{}
	repositories.Spree_db.Create(user)

	err := dbRepo.FindBy(user, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	order := &domain.Order{}
	repositories.Spree_db.Create(&domain.Order{Number: "R123456789", GuestToken: "Xrz5qBnbnoBQnYQYzOMQkQ"})

	err = dbRepo.FindBy(order, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	order.UserId = &user.Id

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders/testOrderNumber"

	r.GET(path, func(c *gin.Context) {
		c.Set("CurrentUser", user)
		c.Set("Order", order)
		authorizeOrder(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 200 {
		t.Errorf("Status code should be 200, but was %d", w.Code)
	}
}

func TestAuthorizeOrderWhenUserIsNotAdminAndOrderDoesNotBelongToHim(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	dbRepo := repositories.NewDatabaseRepository()

	user := &domain.User{}
	repositories.Spree_db.Create(user)

	err := dbRepo.FindBy(user, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	order := &domain.Order{}
	repositories.Spree_db.Create(&domain.Order{Number: "R123456789", GuestToken: "Xrz5qBnbnoBQnYQYzOMQkQ"})

	err = dbRepo.FindBy(order, nil, nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	userId := int64(0)
	order.UserId = &userId

	var ctx *gin.Context
	r := gin.New()

	method := "GET"
	path := "/api/orders/testOrderNumber"

	r.GET(path, func(c *gin.Context) {
		c.Set("CurrentUser", user)
		c.Set("Order", order)
		authorizeOrder(c)
		ctx = c
	})
	w := PerformRequest(r, method, path)

	if w.Code != 401 {
		t.Errorf("Status code should be 401, but was %d", w.Code)
	}
}
