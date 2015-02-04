package api

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		isGuestUser := false
		spreeToken := getSpreeToken(c)
		dbRepo := repositories.NewDatabaseRepository()

		// If spreeToken is empty, check if orderToken is set and action is orders show
		if spreeToken == "" {
			if isOrdersShowAction(c.Request.URL.Path) {
				// Get order token
				orderToken := getOrderToken(c)

				// Return if order token is not provided
				if orderToken == "" {
					unauthorized(c, "You must specify an API key.")
					return
				}

				// Find the order by guest token (order token)
				order := &models.Order{}
				err := dbRepo.FindBy(order, map[string]interface{}{"guest_token": orderToken})
				if err != nil {
					unauthorized(c, "You are not authorized to perform that action.")
					return
				}

				// Get order number and verify if is equal to the order's number from guest token
				orderNumber := getOrderNumber(c)
				if order.Number != orderNumber {
					unauthorized(c, "You are not authorized to perform that action.")
					return
				}

				isGuestUser = true
				c.Set("Order", order)
			} else {
				unauthorized(c, "You must specify an API key.")
				return
			}
		}

		user := &models.User{}
		if !isGuestUser {
			err := dbRepo.FindBy(user, map[string]interface{}{"spree_api_key": spreeToken})

			if err != nil {
				unauthorized(c, "Invalid API key ("+spreeToken+") specified.")
				return
			}

			dbRepo.UserRoles(user)
		}

		c.Set(SPREE_TOKEN, spreeToken)
		c.Set("CurrentUser", user)
		c.Next()
	}
}

func isOrdersShowAction(path string) bool {
	match, _ := regexp.MatchString(`^`+namespace()+`/api/orders/\w+$`, path)
	return match
}

func getOrderNumber(c *gin.Context) string {
	pathArray := strings.Split(c.Request.URL.Path, "/")
	return pathArray[len(pathArray)-1]
}
