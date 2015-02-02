package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestAuthorizationWhenCurrentUserIsAdmin(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/store/api/products", nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	repositories.InitDB()
	user := &models.User{}
	repositories.Spree_db.First(user)

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("CurrentUser", user)
	}, Authorization())

	r.GET("/store/api/products")
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}
}
func TestAuthorizationWhenCurrentUserIsNotAdminAndHasPermissions(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/store/api/products", nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	repositories.InitDB()
	user := &models.User{}
	repositories.Spree_db.FirstOrCreate(user, models.User{SpreeApiKey: "testUser"})

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("CurrentUser", user)
	}, Authorization())

	r.GET("/store/api/products")
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("api.Authentication response code should be 200, but was: %d", w.Code)
	}
}

func TestAuthorizationWhenCurrentUserIsNotAdminAndDoesNotHavePermissions(t *testing.T) {
	t.Skip("After authorization permissions are uncommented, unskiped this")
	routesPattern[`/store/api/users(/*)`] = "users.index"
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/store/api/users", nil)
	if err != nil {
		t.Error("An error occurred: " + err.Error())
	}

	repositories.InitDB()
	user := &models.User{}
	repositories.Spree_db.FirstOrCreate(user, models.User{SpreeApiKey: "testUser"})

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("CurrentUser", user)
	}, Authorization())

	r.GET("/store/api/users")
	r.ServeHTTP(w, req)

	if w.Code != 401 {
		t.Errorf("api.Authentication response code should be 401, but was: %d", w.Code)
	}
}

func TestGetCurrentAction(t *testing.T) {
	testUrl, err := url.Parse("http://localhost:5000/store/api/products")

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if action := getCurrentAction(testUrl); action == "" {
		t.Errorf("api.getCurrentAction should not be empty, but was empty")
	}
}
