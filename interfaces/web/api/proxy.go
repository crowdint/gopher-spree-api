package api

import (
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs"
)

var (
	proxy *httputil.ReverseProxy
)

func init() {
	spreeURL, err := url.Parse(configs.Get(configs.SPREE_URL))

	if err != nil {
		panic(err)
	}

	proxy = httputil.NewSingleHostReverseProxy(spreeURL)
}

func Proxy() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	return !strings.Contains(url.Path, configs.Get(configs.SPREE_NS)+"/api/")
}

func isMissingRoute(method string, url *url.URL) bool {
	for _, route := range routes() {
		if match, _ := regexp.MatchString(route.RegexPattern, url.Path); match && route.Method == method {
			return false
		}
	}

	return true
}
