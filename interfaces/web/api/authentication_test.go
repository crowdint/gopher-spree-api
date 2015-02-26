package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestAuthenticationWithValidToken(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	dbSpreeToken := "testUser"
	repositories.Spree_db.FirstOrCreate(&models.User{}, models.User{SpreeApiKey: dbSpreeToken})

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_TOKEN_HEADER, dbSpreeToken)
	w := httptest.NewRecorder()
	var spreeToken interface{}

	spree.Set(spree.API_AUTHENTICATION, "true")
	r := gin.New()
	r.Use(Authentication(), func(c *gin.Context) {
		spreeToken, _ = c.Get(SPREE_TOKEN)
	})
	r.GET("/products")
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}

	if spreeToken != dbSpreeToken {
		t.Errorf("api.Authentication spree token should be %s, but was %v", dbSpreeToken, spreeToken)
	}
}

func TestAuthenticationWithInvalidToken(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_TOKEN_HEADER, "fooTest")
	w := httptest.NewRecorder()
	var spreeToken interface{}

	spree.Set(spree.API_AUTHENTICATION, "true")
	r := gin.New()
	r.Use(Authentication(), func(c *gin.Context) {
		spreeToken, _ = c.Get(SPREE_TOKEN)
	})
	r.GET("/products")
	r.ServeHTTP(w, req)

	if w.Code != 401 {
		t.Errorf("api.Authentication response code should be 401, but was: %d", w.Code)
	}

	if spreeToken != nil {
		t.Errorf("api.Authentication spree token should be nil, but was %v", spreeToken)
	}
}

func TestAuthenticationWithoutToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	w := httptest.NewRecorder()
	var spreeToken interface{}

	spree.Set(spree.API_AUTHENTICATION, "true")
	r := gin.New()
	r.Use(Authentication(), func(c *gin.Context) {
		spreeToken, _ = c.Get(SPREE_TOKEN)
	})
	r.GET("/products")
	r.ServeHTTP(w, req)

	if spreeToken != nil {
		t.Error("api.Authentication spree token was %v, but should be nil", spreeToken)
	}

	if w.Code != 401 {
		t.Errorf("api.Authentication response code should be 401, but was: %d", w.Code)
	}
}

func TestAuthenticationWithValidOrderToken(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	order := &json.Order{}
	err := repositories.NewDatabaseRepository().FindBy(order, nil, nil)
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	path := "/api/orders/" + order.Number
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_ORDER_TOKEN_HEADER, order.GuestToken)
	w := httptest.NewRecorder()

	r := gin.New()
	var user *models.User
	r.Use(Authentication(), func(c *gin.Context) {
		user = currentUser(c)
	})
	r.GET(path)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}

	if user.Id != -1 {
		t.Errorf("api.Authentication user id should be 0, but was %d", user.Id)
	}
}

func TestAuthenticationWithInvalidOrderToken(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	path := "/api/orders/testOrderNumber"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_ORDER_TOKEN_HEADER, "testOrderToken")
	w := httptest.NewRecorder()

	spree.Set(spree.API_AUTHENTICATION, "true")
	r := gin.New()
	r.Use(Authentication())
	r.GET(path)
	r.ServeHTTP(w, req)

	if w.Code != 401 {
		t.Errorf("api.Authentication response code should be 401, but was: %d", w.Code)
	}
}

func TestAuthenticationWithValidOrderTokenAndActionIsNotOrderShow(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	order := &json.Order{}
	err := repositories.NewDatabaseRepository().FindBy(order, nil, nil)
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	path := "/api/orders"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_ORDER_TOKEN_HEADER, order.GuestToken)
	w := httptest.NewRecorder()

	spree.Set(spree.API_AUTHENTICATION, "true")
	r := gin.New()
	var user *models.User
	r.Use(Authentication(), func(c *gin.Context) {
		user = currentUser(c)
	})
	r.GET(path)
	r.ServeHTTP(w, req)

	if w.Code != 401 {
		t.Errorf("api.Authentication response code should be 401, but was: %d", w.Code)
	}
}

func TestAuthenticationWithoutTokenAndAuthenticationRequiredIsFalse(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	path := "/api/products"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	w := httptest.NewRecorder()

	spree.Set(spree.API_AUTHENTICATION, "false")
	r := gin.New()
	var user *models.User
	r.Use(Authentication(), func(c *gin.Context) {
		user = currentUser(c)
	})
	r.GET(path)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}
}

func TestAuthenticationWithoutTokenAndAuthenticationRequiredIsFalseAndActionIsNotRead(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	path := "/api/products"
	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	w := httptest.NewRecorder()

	spree.Set(spree.API_AUTHENTICATION, "false")
	r := gin.New()
	var user *models.User
	r.Use(Authentication(), func(c *gin.Context) {
		user = currentUser(c)
	})
	r.POST(path)
	r.ServeHTTP(w, req)

	if w.Code != 401 {
		t.Errorf("api.Authentication response code should be 401, but was: %d", w.Code)
	}
}

func TestAuthenticationWithTokenAndAuthenticationRequiredIsFalseAndActionsIsNotRead(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	user := &models.User{}
	var currentUsr *models.User
	dbSpreeToken := "testUser"
	repositories.Spree_db.FirstOrCreate(user, models.User{SpreeApiKey: dbSpreeToken})

	path := "/api/products"
	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_TOKEN_HEADER, dbSpreeToken)
	w := httptest.NewRecorder()

	spree.Set(spree.API_AUTHENTICATION, "false")
	r := gin.New()
	r.Use(Authentication(), func(c *gin.Context) {
		currentUsr = currentUser(c)
	})
	r.POST(path)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}

	if user.Id != currentUsr.Id {
		t.Errorf("api.Authentication user id should be %d, but was: %d", user.Id, currentUsr.Id)
	}
}

func TestGetSpreeTokenWhenNotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	context := &gin.Context{Request: req}

	if token := getSpreeToken(context); token != "" {
		t.Error("api.getSpreeToken should be \"\", but was %s", token)
	}
}

func TestGetSpreeTokenFromHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	req.Header.Set(SPREE_TOKEN_HEADER, "spree123")

	context := &gin.Context{Request: req}

	if token := getSpreeToken(context); token != "spree123" {
		t.Errorf("api.getSpreeToken should be %s, but was %s", "spree123", token)
	}
}

func TestGetSpreeTokenFromURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/products?token=spree123", nil)
	if err != nil {
		t.Error("An error occurred:", err.Error())
	}

	context := &gin.Context{Request: req}

	if token := getSpreeToken(context); token != "spree123" {
		t.Errorf("api.getSpreeToken should be %s, but was %s", "spree123", token)
	}
}
