package api

import (
	"errors"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/utils"
)

var (
	proxy *httputil.ReverseProxy
)

func init() {
	spreeURL, err := url.Parse(configs.Get(configs.SPREE_URL))

	if err != nil {
		utils.LogrusError(utils.FuncName(), err)
		panic(err)
	}
	utils.LogrusInfo(utils.FuncName(), "Request to spreeURL")

	proxy = httputil.NewSingleHostReverseProxy(spreeURL)

}

func Proxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		setOriginPolicyHeaders(c)
		if shouldRedirectToOrigin(c) {
			c.Abort()
			proxy.ServeHTTP(c.Writer, c.Request)
			return
		}

		c.Next()
	}
}

func shouldRedirectToOrigin(c *gin.Context) bool {
	url := c.Request.URL
	method := c.Request.Method
	return isNotRequestToApi(url) || isMissingRoute(method, url)
}

func isNotRequestToApi(url *url.URL) bool {
	err := errors.New(url.String())
	utils.LogrusError(utils.FuncName(), err)
	return !strings.Contains(url.Path, configs.Get(configs.SPREE_NS)+"/api/")
}

func isMissingRoute(method string, url *url.URL) bool {
	for _, route := range routes() {
		if match, _ := regexp.MatchString(route.RegexPattern, url.Path); match && route.Method == method {
			return false
		}
	}
	err := errors.New("Method: " + method + "url: " + url.String())

	utils.LogrusError(utils.FuncName(), err)
	return true
}

func setOriginPolicyHeaders(c *gin.Context) {
	originConfig := configs.Get(configs.X_ORIGIN)
	if originConfig != "" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", configs.Get(configs.X_ORIGIN))
		c.Writer.Header().Set("Access-Control-Allow-Methods", configs.Get(configs.X_METHODS))
	}

}
