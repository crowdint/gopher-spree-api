package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/utils"
)

const (
	SPREE_TOKEN_HEADER       = "X-Spree-Token"
	SPREE_TOKEN              = "SpreeToken"
	SPREE_ORDER_TOKEN_HEADER = "X-Spree-Order-Token"
)

var (
	ns *string
)

type params map[string]interface{}

func currentUser(c *gin.Context) *domain.User {
	currentUser, _ := c.Get("CurrentUser")
	user := currentUser.(*domain.User)
	return user
}

func fail(c *gin.Context, err error) {
	if err.Error() == "Record Not Found" {
		notFound(c)
	} else {
		internalServerError(c, err.Error())
	}
}

func getOrderToken(c *gin.Context) string {
	orderToken := c.Request.Header.Get(SPREE_ORDER_TOKEN_HEADER)

	if orderToken != "" {
		utils.LogrusInfo(utils.FuncName(), "orderToken= "+orderToken)
		return orderToken
	}

	utils.LogrusInfo(utils.FuncName(), c.Request.URL.Query().Get("order_token"))
	return c.Request.URL.Query().Get("order_token")
}

func getSpreeToken(c *gin.Context) string {
	spreeToken := c.Request.Header.Get(SPREE_TOKEN_HEADER)

	if len(spreeToken) > 0 {
		utils.LogrusInfo(utils.FuncName(), spreeToken)
		return spreeToken
	}
	utils.LogrusInfo(utils.FuncName(), c.Request.URL.Query().Get("token"))
	return c.Request.URL.Query().Get("token")
}

func internalServerError(c *gin.Context, errMsg string) {
	c.JSON(500, gin.H{"error": errMsg})
	c.Abort()
}

func namespace() string {
	if ns == nil {
		temp := configs.Get(configs.SPREE_NS)
		if temp != "" {
			temp = "/" + strings.Replace(temp, "/", "", -1)
		}
		ns = &temp
	}

	return *ns
}

func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"error": "Record Not Found"})
	c.Abort()
}

func unauthorized(c *gin.Context, errMsg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	c.Abort()
}
