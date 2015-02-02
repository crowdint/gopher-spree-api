package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestAuthenticationWithValidToken(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	dbSpreeToken := "testUser"
	repositories.Spree_db.FirstOrCreate(&models.User{}, models.User{SpreeApiKey: dbSpreeToken})

	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set(SPREE_TOKEN_HEADER, dbSpreeToken)
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

	if spreeToken != dbSpreeToken {
		t.Errorf("api.Authentication spree token should be %s, but was %v", dbSpreeToken, spreeToken)
	}
}

func TestAuthenticationWithInvalidToken(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set(SPREE_TOKEN_HEADER, "fooTest")
	w := httptest.NewRecorder()
	var spreeToken interface{}

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
