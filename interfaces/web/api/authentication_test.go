package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthenticationWhenProxied(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set(SPREE_TOKEN_HEADER, "spree123")
	w := httptest.NewRecorder()
	var context *gin.Context

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("Proxied", true)
		context = c
	}, Authentication())

	r.GET("/products")
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}

	if spreeToken, err := context.Get(SPREE_TOKEN); err == nil {
		t.Errorf("api.Authentication spree token should be nil, but was %s", spreeToken)
	}
}

func TestAuthenticationWithToken(t *testing.T) {
	t.Skip("TestAuthenticationWithToken: Implement database connection")
	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set(SPREE_TOKEN_HEADER, "spree123")
	w := httptest.NewRecorder()
	var spreeToken interface{}

	r := gin.New()
	r.Use(Authentication(), func(c *gin.Context) {
		spreeToken, _ = c.Get(SPREE_TOKEN)
	})
	r.GET("/products")
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}

	if spreeToken != "spree123" {
		t.Errorf("api.Authentication spree token should be %s, but was %v", "spree123", spreeToken)
	}
}

func TestAuthenticationWithoutToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	var spreeToken interface{}

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

func TestGetSpreeTokenWhenNotPresent(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)

	context := &gin.Context{Request: req}

	if token := getSpreeToken(context); token != "" {
		t.Error("api.getSpreeToken should be \"\", but was %s", token)
	}
}

func TestGetSpreeTokenFromHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set(SPREE_TOKEN_HEADER, "spree123")

	context := &gin.Context{Request: req}

	if token := getSpreeToken(context); token != "spree123" {
		t.Errorf("api.getSpreeToken should be %s, but was %s", "spree123", token)
	}
}

func TestGetSpreeTokenFromURL(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products?token=spree123", nil)

	context := &gin.Context{Request: req}

	if token := getSpreeToken(context); token != "spree123" {
		t.Errorf("api.getSpreeToken should be %s, but was %s", "spree123", token)
	}
}