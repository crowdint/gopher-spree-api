package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/usecases/json"
)

func init() {
	products := API().Group("/products")
	{
		products.GET("", ProductsIndex)
		products.GET("/", ProductsIndex)
		products.GET("/:product_id", ProductsShow)
		products.POST("", authorizeProduct, ProductsCreate)
		products.POST("/", authorizeProduct, ProductsCreate)
	}
}

func ProductsIndex(c *gin.Context) {
	params := NewRequestParameters(c, json.GRANSAK)

	productResponse(c, params)
}

func ProductsShow(c *gin.Context) {
	params := NewRequestParameters(c, json.GRANSAK)

	product, err := json.SpreeResponseFetcher.GetShowResponse(json.NewProductInteractor(), params)

	if err == nil {
		c.JSON(200, product)
		return
	}

	if err.Error() == "Record Not Found" {
		notFound(c)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func productResponse(c *gin.Context, params *RequestParameters) {
	if products, err := json.SpreeResponseFetcher.GetResponse(json.NewProductInteractor(), params); err != nil && err.Error() != "Record Not Found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, products)
	}
}

func ProductsCreate(c *gin.Context) {
	params := NewRequestParameters(c, json.GRANSAK)
	product, productError, err := json.SpreeResponseFetcher.GetCreateResponse(json.NewProductInteractor(), params)

	if err != nil && productError == nil {
		c.JSON(422, gin.H{"error": err.Error()})
	} else if productError != nil {
		c.JSON(422, gin.H{"error": err.Error(), "errors": productError})
	} else {
		c.JSON(201, product)
	}
}

func authorizeProduct(c *gin.Context) {
	user := currentUser(c)
	if user.HasRole("admin") {
		c.Next()
		return
	}

	unauthorized(c, "You are not authorized to perform that action.")
}
