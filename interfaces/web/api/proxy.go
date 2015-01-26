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
	spreeUrl, err := url.Parse(os.Getenv("SPREE_URL"))

	if err != nil {
		panic(err)
	}

	proxy = httputil.NewSingleHostReverseProxy(spreeUrl)
}

func Proxy() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method != "GET" || isMissingURL(c.Request.URL) {
			proxy.ServeHTTP(c.Writer, c.Request)
		}

		c.Next()
	}
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
