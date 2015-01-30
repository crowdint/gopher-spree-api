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
	req, _ := http.NewRequest("GET", "/store/api/products", nil)
	w := httptest.NewRecorder()

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
	req, _ := http.NewRequest("GET", "/store/api/products", nil)
	w := httptest.NewRecorder()

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
	routesPattern[`/store/api/users(/*)`] = "users.index"
	req, _ := http.NewRequest("GET", "/store/api/users", nil)
	w := httptest.NewRecorder()

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

func TestGetCurrentUserRole(t *testing.T) {
	if err := repositories.InitDB(); err != nil {
		t.Error("An error has ocurred", err)
	}

	repositories.InitDB()
	user := &models.User{}
	repositories.Spree_db.First(user)

	if role := getCurrentUserRole(user); role == nil {
		t.Error("api.getCurrentUserRole should not be nil, but was nil")
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
