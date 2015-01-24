package api

import (
	"net/http/httputil"
	"net/url"
	"os"

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

func Router() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method != "GET" || missingHandler(c.Request.URL) {
			proxy.ServeHTTP(c.Writer, c.Request)
		}

		c.Next()
	}
}

func missingHandler(url *url.URL) bool {
	// TODO: Determine which GET routes should be proxied to Spree
	// In essence endpoints not yet implemented in Go.
	return false
}
