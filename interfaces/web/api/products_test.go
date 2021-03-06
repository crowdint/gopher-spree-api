package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductsIndex(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	r := gin.New()

	method := "GET"
	path := "/api/products/"

	repositories.Spree_db.Create(&domain.User{})

	r.GET(path, func(c *gin.Context) {
		user := &domain.User{}
		repositories.Spree_db.First(user)
		c.Set("CurrentUser", user)
		ProductsIndex(c)
	})
	w := PerformRequest(r, method, path, nil)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}

func TestProductsShow(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	productPrice := 12.0
	repositories.Spree_db.Create(&domain.Product{Id: 1, Name: "Test Product", Price: &productPrice, ShippingCategoryId: 1, Slug: "test-product"})

	r := gin.New()

	method := "GET"
	path := "/api/products/1"

	r.GET("/api/products/:product_id", func(c *gin.Context) {
		ProductsShow(c)
	})
	w := PerformRequest(r, method, path, nil)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}

func TestProductsCreate(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	file, err := os.Open("../../../test/data/products/with_shipping_category.json")
	if err != nil {
		t.Error("An error occured while trying to open JSON file: ", err.Error())
	}

	productParams, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("An error occured: ", err.Error())
	}

	r := gin.New()

	method := "POST"
	path := "/api/products/"

	r.POST(path, func(c *gin.Context) {
		ProductsCreate(c)
	})
	w := PerformRequest(r, method, path, bytes.NewBuffer(productParams))

	if w.Code != http.StatusCreated {
		t.Errorf("Status code should be %d, but was %d -> %s", http.StatusCreated, w.Code, w.Body.String())
	}
}

func TestAuthorizeProducts_AdminUser(t *testing.T) {
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

	method := "POST"
	path := "/api/products"

	r.POST(path, func(c *gin.Context) {
		c.Set("CurrentUser", user)
		authorizeProduct(c)
		ctx = c
	})
	w := PerformRequest(r, method, path, nil)

	if w.Code != 200 {
		t.Errorf("Status code should be 200, but was %d", w.Code)
	}
}

func TestAuthorizeProducts_NormalUser(t *testing.T) {
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

	method := "POST"
	path := "/api/products"

	r.POST(path, func(c *gin.Context) {
		c.Set("CurrentUser", user)
		authorizeProduct(c)
		ctx = c
	})
	w := PerformRequest(r, method, path, nil)

	if w.Code != 401 {
		t.Errorf("Status code should be 401, but was %d", w.Code)
	}
}
