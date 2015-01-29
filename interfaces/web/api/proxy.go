package api

import (
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

var (
	proxy *httputil.ReverseProxy
)

func init() {
	spreeURL, err := url.Parse(os.Getenv("SPREE_URL"))

	if err != nil {
		panic(err)
	}

	proxy = httputil.NewSingleHostReverseProxy(spreeURL)
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
	for _, pattern := range regexRoutesPattern() {
		if match, _ := regexp.MatchString(pattern, url.Path); match {
			return false
		}
	}
	return true
}

func regexRoutesPattern() []string {
	return []string{
		`/api/products(/*)`, // products index
		`/api/products/\d`,  // products show
	}
}
