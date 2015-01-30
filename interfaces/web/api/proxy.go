package api

import (
	"net/http/httputil"
	"net/url"
	"regexp"

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
	routesPattern = regexRoutesPattern()
}

func Proxy() gin.HandlerFunc {
	return func(c *gin.Context) {

		if shouldRedirectToOrigin(c) {
			c.Set("Proxied", true)
			proxy.ServeHTTP(c.Writer, c.Request)
		}

		c.Next()
	}
}

func shouldRedirectToOrigin(c *gin.Context) bool {
	return c.Request.Method != "GET" || isMissingURL(c.Request.URL)
}

func isMissingURL(url *url.URL) bool {
	for pattern, _ := range routesPattern {
		if match, _ := regexp.MatchString(pattern, url.Path); match {
			return false
		}
	}
	return true
}
